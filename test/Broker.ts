import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployMarketplaceFixture } from './fixtures'
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "ethers";

const HARDHAT_NETWORK_ID = 31337;

type UnsignedOffer = {
    specs: string;
    pricePerMinute: number;
    client: string;
    expiresAt: number;
    nonce: number;
}

async function signOffer(
    provider: SignerWithAddress,
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
            { name: 'specs', type: 'bytes32' },
            { name: 'pricePerMinute', type: 'uint256' },
            { name: 'client', type: 'address' },
            { name: 'expiresAt', type: 'uint256' },
            { name: 'nonce', type: 'uint32' },
        ]
    }

    return provider._signTypedData(domain, types, offer)
}


describe("Broker", function () {
    describe("bsookResource", function () {
        it("should create a new Booking", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);

            const tx = await marketplace.connect(user).bookResource(offer, signature);
            const rc = await tx.wait();

            const event = rc.events?.find(event => event.event === 'BookingCreated');
            const newId = event?.args?.bookingId;

            const bookingFromChain = await marketplace.getBooking(newId);

            expect(bookingFromChain.specs).to.equal(offer.specs);
            expect(bookingFromChain.pricePerMinute).to.equal(offer.pricePerMinute);
            expect(bookingFromChain.client).to.equal(user.address);
            expect(bookingFromChain.provider).to.equal(provider.address);
        })

        it("should increase locked balance", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };

            const [, locked1] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(locked1).to.equal(0);

            const signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            const [, locked2] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(locked2).to.equal(100 * 60 * 24 * 7);
        })
        it.skip("should execute claim payment", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };

            const signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            expect(false).to.equal(true, "not implemented");
        })
        it("should fail if user tries to reuse the same offer", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };

            const [, locked1] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(locked1).to.equal(0);

            const signature = await signOffer(provider, offer, marketplace.address);
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.not.be.reverted;
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.be.revertedWith("Invalid offer nonce");
        })
        it("should fail if not enough balance to cover a new VM booking")
    })
    describe("cancelBooking", function () {
        it("should remove a booking")
        it("should decrease locked balance")
        it("should fail if booking is already cancelled")
        it("should create an event with reason = 0 or 1 depending on client's selection")
        it("should create an event with reason = 2 if provider cancelled with no reason")
        it("should create an event with reason = 3 if provider cancelled and user's balance is zero")
    })
    describe("claimPayment", function () {
        it("should increase provider's balance")
        it("should transfer not more than user's balance")
    });

});
