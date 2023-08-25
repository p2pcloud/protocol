import { ethers, upgrades } from "hardhat";
import { FiatMarketplaceV2, MarketplaceV2, MarketplaceV3 } from "../typechain-types";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { BigNumber, BigNumberish } from "ethers";
import { expect } from "chai";
import { time } from "@nomicfoundation/hardhat-network-helpers";
import { MarketplaceFixture } from "./fixtures";

const HARDHAT_NETWORK_ID = 31337;

const DEFAULT_USER_BALANCE = 100000
type UnsignedOffer = {
    specs: string;
    pricePerMinute: number;
    client: string;
    expiresAt: number;
    nonce: number;
}

async function signOffer(
    signer: SignerWithAddress,
    offer: UnsignedOffer,
    brokerAddress: string,
): Promise<string> {
    const domain = {
        chainId: HARDHAT_NETWORK_ID,
        name: 'p2pcloud.io',
        verifyingContract: brokerAddress,
        version: '2',
    }

    const types = {
        UnsignedOffer: [
            { name: 'client', type: 'address' },
            { name: 'pricePerMinute', type: 'uint64' },
            { name: 'nonce', type: 'uint32' },
            { name: 'specs', type: 'bytes32' },
            { name: 'expiresAt', type: 'uint256' },
        ]
    }

    return signer._signTypedData(domain, types, offer)
}

type UnsignedVoucher = {
    amount: number;
    paymentId: string;
}

async function signVoucher(
    signer: SignerWithAddress,
    voucher: UnsignedVoucher,
    brokerAddress: string,
): Promise<string> {
    const domain = {
        chainId: HARDHAT_NETWORK_ID,
        name: 'p2pcloud.io',
        verifyingContract: brokerAddress,
        version: '2',
    }

    const types = {
        UnsignedVoucher: [
            { name: 'amount', type: 'uint256' },
            { name: 'paymentId', type: 'bytes32' },
        ]
    }

    return signer._signTypedData(domain, types, voucher)
}

type V2V3MigrationFixture = {
    provider: SignerWithAddress,
    providersSigner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
    marketplace: FiatMarketplaceV2,
    voucherSigner: SignerWithAddress,
    marketplaceV3: MarketplaceV3
}

export async function deployMarketplaceV2Fixture(): Promise<V2V3MigrationFixture> {
    const [admin, provider, user, anotherUser, providersSigner, voucherSigner] = await ethers.getSigners();

    const Marketplace = await ethers.getContractFactory("FiatMarketplaceV2");
    const marketplace = await upgrades.deployProxy(Marketplace, ["0x0000000000000000000000000000000000000000"]) as FiatMarketplaceV2;

    await marketplace.connect(admin).setVoucherSigner(voucherSigner.address)

    await marketplace.connect(admin).registerFiatProvider(provider.address)
    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    const voucher = { amount: DEFAULT_USER_BALANCE, paymentId: ethers.utils.formatBytes32String("fixtue2") }
    const signature = await signVoucher(voucherSigner, voucher, marketplace.address)
    await marketplace.connect(anotherUser).claimVoucher(voucher, signature, user.address)

    //book vm1
    const offer1: UnsignedOffer = {
        specs: ethers.utils.formatBytes32String("hello world"),
        pricePerMinute: 1,
        client: user.address,
        expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
        nonce: await marketplace.getNonce(user.address),
    };
    const signature1 = await signOffer(providersSigner, offer1, marketplace.address);
    await marketplace.connect(user).bookResource(offer1, signature1);

    //book vm 2
    const offer2: UnsignedOffer = {
        specs: ethers.utils.formatBytes32String("hello world"),
        pricePerMinute: 3,
        client: user.address,
        expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
        nonce: await marketplace.getNonce(user.address),
    };
    const signature2 = await signOffer(providersSigner, offer2, marketplace.address);
    await marketplace.connect(user).bookResource(offer2, signature2);

    //book vm 3
    const offer3: UnsignedOffer = {
        specs: ethers.utils.formatBytes32String("hello world"),
        pricePerMinute: 2,
        client: user.address,
        expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
        nonce: await marketplace.getNonce(user.address),
    };
    const signature3 = await signOffer(providersSigner, offer3, marketplace.address);
    await marketplace.connect(user).bookResource(offer3, signature3);

    //wait for 1 hour
    await time.latest();
    await time.increase(3600);
    await marketplace.connect(provider).claimPayment(user.address)

    await marketplace.connect(user).cancelBooking(1, true);

    const booking1 = await marketplace.getBooking(1);
    expect(booking1.specs).to.equal(ethers.utils.formatBytes32String(""), "booking1 specs in fixture");
    expect(booking1.pricePerMinute).to.equal(0, "booking1 pricePerMinute in fixture");

    // There is a bug in V2 that leaves the provider and client address in the booking
    // expect(booking2.provider).to.equal(ethers.constants.AddressZero, "booking2 provider in fixture");
    // expect(booking2.client).to.equal(ethers.constants.AddressZero, "booking2 client in fixture");

    //remember balances
    const userbalance1 = await marketplace.getBalance(user.address);
    expect(userbalance1.free).to.equal('69400');
    expect(userbalance1.locked).to.equal('30240');

    const providerBalance1 = await marketplace.getBalance(provider.address);
    expect(providerBalance1.free).to.equal('288');
    expect(providerBalance1.locked).to.equal('0');

    const MarketplaceV3 = await ethers.getContractFactory("MarketplaceV3");
    const marketplaceV3 = await upgrades.deployProxy(MarketplaceV3, ["0x0000000000000000000000000000000000000000"]) as MarketplaceV3;

    await marketplaceV3.connect(admin).performMigration(marketplace.address);

    return { marketplace, provider, user, admin, anotherUser, providersSigner, voucherSigner, marketplaceV3 };
}


describe("V2_V3_Migration", () => {
    it("should keep balances", async () => {
        const { marketplace, provider, user, admin, marketplaceV3 } = await deployMarketplaceV2Fixture();

        //check balances staid the same
        const userbalance1 = await marketplace.getBalance(user.address);
        expect(userbalance1.free).to.equal('69400');
        expect(userbalance1.locked).to.equal('30240');

        const providerBalance1 = await marketplace.getBalance(provider.address);
        expect(providerBalance1.free).to.equal('288');
        expect(providerBalance1.locked).to.equal('0');
    })
    it("should keep non deleted bookings", async () => {
        const { marketplace, provider, user, admin, marketplaceV3 } = await deployMarketplaceV2Fixture();

        const booking0 = await marketplaceV3.getBooking(0);
        expect(booking0.provider).to.equal(provider.address);
        expect(booking0.client).to.equal(user.address);
        expect(booking0.specs).to.equal(ethers.utils.formatBytes32String("hello world"));
        expect(booking0.pricePerMinute).to.equal(1);

        const booking1 = await marketplaceV3.getBooking(1);
        expect(booking1.provider).to.equal(ethers.constants.AddressZero);
        expect(booking1.client).to.equal(ethers.constants.AddressZero);
        expect(booking1.specs).to.equal(ethers.utils.formatBytes32String(""));
        expect(booking1.pricePerMinute).to.equal(0);

        const booking2 = await marketplaceV3.getBooking(2);
        expect(booking2.specs).to.equal(ethers.utils.formatBytes32String("hello world"));
        expect(booking2.pricePerMinute).to.equal(2);
        expect(booking2.provider).to.equal(provider.address);
        expect(booking2.client).to.equal(user.address);
    })
})