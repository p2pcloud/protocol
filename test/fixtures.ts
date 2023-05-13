import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { type TestableMarketplace, type MockERC20, FiatMarketplace__factory, FiatMarketplace } from "../typechain-types";

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


export async function deployFiatMarketplaceFixture(): Promise<FiatMarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner, voucherSigner] = await ethers.getSigners();

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const token = await MockERC20.connect(admin).deploy('10000000000000000000000');

    const Marketplace = await ethers.getContractFactory("FiatMarketplace");
    const marketplace = await upgrades.deployProxy(Marketplace, [token.address]) as FiatMarketplace;

    const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
    await token.connect(admin).transfer(provider.address, fee)
    await token.connect(provider).increaseAllowance(marketplace.address, fee)
    await marketplace.connect(provider).depositCoin(fee)
    await marketplace.connect(admin).registerFiatProvider(provider.address)
    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    await token.connect(admin).transfer(user.address, '10000000')
    await token.connect(user).approve(marketplace.address, '10000000')
    await marketplace.connect(user).depositCoin('10000000')


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
    await token.connect(admin).transfer(user.address, '10000000')
    await token.connect(user).approve(marketplace.address, '10000000')
    await marketplace.connect(user).depositCoin('10000000')


    return { marketplace, token, provider, user, admin, anotherUser, providersSigner };
}