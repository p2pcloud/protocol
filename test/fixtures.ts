import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { type TestableMarketplace, type MockERC20, FiatMarketplace__factory, FiatMarketplace } from "../typechain-types";
import { US_HEX, signKYC, signVoucher } from "./lib";

type Fixture = {
    provider: SignerWithAddress,
    providersSigner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
    kycSigner: SignerWithAddress,
}

export type MarketplaceFixture = Fixture & {
    marketplace: TestableMarketplace,
    token: MockERC20
}

export type FiatMarketplaceFixture = Fixture & {
    marketplace: FiatMarketplace,
    voucherSigner: SignerWithAddress,
}

export const DEFAULT_USER_BALANCE = 10000000

export async function deployFiatMarketplaceFixture(): Promise<FiatMarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner, voucherSigner, kycSigner] = await ethers.getSigners();

    const Marketplace = await ethers.getContractFactory("FiatMarketplace");
    const marketplace = await upgrades.deployProxy(Marketplace, ["0x0000000000000000000000000000000000000000"]) as FiatMarketplace;

    await marketplace.connect(admin).setVoucherSigner(voucherSigner.address)

    await marketplace.connect(admin).registerFiatProvider(provider.address)
    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    const voucher = { amount: DEFAULT_USER_BALANCE, paymentId: ethers.utils.formatBytes32String("fixtue2") }
    const signature = await signVoucher(voucherSigner, voucher, marketplace.address)
    await marketplace.connect(anotherUser).claimVoucher(voucher, signature, user.address)

    return { marketplace, provider, user, admin, anotherUser, providersSigner, voucherSigner, kycSigner };
}

export async function deployMarketplaceFixture(): Promise<MarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner, kycSigner] = await ethers.getSigners();

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const token = await MockERC20.connect(admin).deploy('10000000000000000000000');

    const Marketplace = await ethers.getContractFactory("TestableMarketplace");
    const marketplace = await upgrades.deployProxy(Marketplace, [token.address]) as TestableMarketplace;

    const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
    await token.connect(admin).transfer(provider.address, fee)
    await token.connect(provider).increaseAllowance(marketplace.address, fee)
    await marketplace.connect(provider).depositCoin(fee)

    await marketplace.connect(admin).allowProviderCountry(US_HEX)
    await marketplace.connect(admin).allowUserCountry(US_HEX)
    await marketplace.connect(admin).setKYCSigner(kycSigner.address)

    await marketplace.connect(provider).submitKYC(
        provider.address,
        US_HEX,
        await signKYC(provider.address, US_HEX, kycSigner)
    )

    //TODO: generate signature and set user KYC
    await marketplace.connect(provider).registerProvider()

    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    await token.connect(admin).transfer(user.address, DEFAULT_USER_BALANCE)
    await token.connect(user).approve(marketplace.address, DEFAULT_USER_BALANCE)
    await marketplace.connect(user).depositCoin(DEFAULT_USER_BALANCE)

    return { marketplace, token, provider, user, admin, anotherUser, providersSigner, kycSigner };
}