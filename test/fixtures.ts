import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { type TestableMarketplace, type MockERC20, FiatMarketplace__factory, FiatMarketplace, P2PCloudCredit } from "../typechain-types";
import { signVoucher } from "./lib";

type Fixture = {
    provider: SignerWithAddress,
    providersSigner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
    marketplace: TestableMarketplace,
}

export type MarketplaceFixture = Fixture & {
    token: MockERC20
}

export type FiatMarketplaceFixture = Fixture & {
    token: P2PCloudCredit,
}


export const DEFAULT_USER_BALANCE = 10000000

export async function deployFiatMarketplaceFixture(): Promise<FiatMarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner] = await ethers.getSigners();

    const P2PCloudCredit = await ethers.getContractFactory("P2PCloudCredit");
    const token = await upgrades.deployProxy(P2PCloudCredit, [admin.address]) as P2PCloudCredit;

    await token.connect(admin).mint(admin.address, '10000000000000000000000')

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