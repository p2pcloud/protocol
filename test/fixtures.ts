import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { type TestableMarketplace, type MockERC20, FiatMarketplace__factory, FiatMarketplace } from "../typechain-types";
import { signVoucher } from "./lib";

type Fixture = {
    token: MockERC20
    provider: SignerWithAddress,
    providersSigner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
}

export type MarketplaceFixture = Fixture & {
    marketplace: TestableMarketplace,
}

export type FiatMarketplaceFixture = Fixture & {
    marketplace: FiatMarketplace,
    voucherSigner: SignerWithAddress,
}


export const DEFAULT_USER_BALANCE = 10000000

export async function deployFiatMarketplaceFixture(): Promise<FiatMarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner, voucherSigner] = await ethers.getSigners();

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const token = await MockERC20.connect(admin).deploy('10000000000000000000000');

    const Marketplace = await ethers.getContractFactory("FiatMarketplace");
    const marketplace = await upgrades.deployProxy(Marketplace, [token.address]) as FiatMarketplace;

    await marketplace.connect(admin).setVoucherSigner(voucherSigner.address)

    const fee = (await marketplace.PROVIDER_REGISTRATION_FEE()).toNumber()
    let voucher = { amount: fee, paymentId: ethers.utils.formatBytes32String("fixtue1") }
    let signature = await signVoucher(voucherSigner, voucher, marketplace.address)
    await marketplace.connect(provider).claimVoucher(voucher, signature)


    await marketplace.connect(admin).registerFiatProvider(provider.address)
    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    voucher = { amount: DEFAULT_USER_BALANCE, paymentId: ethers.utils.formatBytes32String("fixtue2") }
    signature = await signVoucher(voucherSigner, voucher, marketplace.address)
    await marketplace.connect(user).claimVoucher(voucher, signature)

    return { marketplace, token, provider, user, admin, anotherUser, providersSigner, voucherSigner };
}

export async function deployMarketplaceFixture(): Promise<MarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner] = await ethers.getSigners();

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const token = await MockERC20.connect(admin).deploy('10000000000000000000000');

    const Marketplace = await ethers.getContractFactory("TestableMarketplace");
    const marketplace = await upgrades.deployProxy(Marketplace, [token.address]) as TestableMarketplace;

    const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
    await token.connect(admin).transfer(provider.address, fee)
    await token.connect(provider).increaseAllowance(marketplace.address, fee)
    await marketplace.connect(provider).depositCoin(fee)
    await marketplace.connect(provider).registerProvider()
    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    await token.connect(admin).transfer(user.address, DEFAULT_USER_BALANCE)
    await token.connect(user).approve(marketplace.address, DEFAULT_USER_BALANCE)
    await marketplace.connect(user).depositCoin(DEFAULT_USER_BALANCE)


    return { marketplace, token, provider, user, admin, anotherUser, providersSigner };
}