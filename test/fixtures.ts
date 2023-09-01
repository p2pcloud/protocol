import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { US_HEX, randomBytes12, signKYC } from "./lib";
import { MockERC20V3, TestableMarketplaceV3 } from "../typechain-types";

type BaseFixture = {
    provider: SignerWithAddress,
    providersSigner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
    kycSigner: SignerWithAddress,
    marketplace: TestableMarketplaceV3,
}

export type MarketplaceFixture = BaseFixture & {
    stablecoin: MockERC20V3,
}

export type CreditMarketplaceFixture = BaseFixture & {
}

export const DEFAULT_USER_BALANCE = 10000000


export async function deployMarketplaceV3Fixture(): Promise<MarketplaceFixture> {
    const fixture = await fixtureFactory("stablecoin") as MarketplaceFixture
    return fixture
}

export async function deployCreditMarketplaceV3Fixture(): Promise<CreditMarketplaceFixture> {
    const fixture = await fixtureFactory("credits") as CreditMarketplaceFixture
    return fixture
}

async function fixtureFactory(fixtureVariant: "credits" | "stablecoin"): Promise<MarketplaceFixture | CreditMarketplaceFixture> {
    const [admin, provider, user, anotherUser, providersSigner, kycSigner] = await ethers.getSigners();

    let stablecoin: MockERC20V3 | null = null
    let coinAddress: string = ""

    if (fixtureVariant === "stablecoin") {
        const MockERC20V3 = await ethers.getContractFactory("MockERC20V3");
        stablecoin = await MockERC20V3.connect(admin).deploy('10000000000000000000000')
        coinAddress = stablecoin.address
    } else if (fixtureVariant === "credits") {
        coinAddress = ethers.constants.AddressZero
    } else {
        throw "Unknown fixture variant"
    }

    const Marketplace = await ethers.getContractFactory("TestableMarketplaceV3");
    const marketplace = await upgrades.deployProxy(Marketplace, [coinAddress]) as TestableMarketplaceV3;

    //KYC quirks
    await marketplace.connect(admin).allowProviderCountry(US_HEX)
    await marketplace.connect(admin).allowUserCountry(US_HEX)
    await marketplace.connect(admin).setKYCSigner(kycSigner.address)

    //KYC user, admin and provider
    await marketplace.connect(admin).submitKYC(
        admin.address,
        US_HEX,
        await signKYC(admin.address, US_HEX, kycSigner)
    )

    await marketplace.connect(provider).submitKYC(
        provider.address,
        US_HEX,
        await signKYC(provider.address, US_HEX, kycSigner)
    )

    await marketplace.connect(user).submitKYC(
        user.address,
        US_HEX,
        await signKYC(user.address, US_HEX, kycSigner)
    )

    //register provider
    await marketplace.connect(provider).submitKYC(
        provider.address,
        US_HEX,
        await signKYC(provider.address, US_HEX, kycSigner)
    )

    //TODO: generate signature and set user KYC
    await marketplace.connect(admin).registerProviderByCommunity(provider.address)

    await marketplace.connect(provider).setSigner(providersSigner.address)

    //transfer some tokens to user
    if (stablecoin !== null) {
        await stablecoin.connect(admin).transfer(user.address, DEFAULT_USER_BALANCE)
        await stablecoin.connect(user).approve(marketplace.address, DEFAULT_USER_BALANCE)
        await marketplace.connect(user).depositCoin(DEFAULT_USER_BALANCE)
    } else {
        await marketplace.connect(kycSigner).mint(user.address, DEFAULT_USER_BALANCE, randomBytes12())
    }

    const baseFixture: BaseFixture = { marketplace, provider, user, admin, anotherUser, providersSigner, kycSigner };

    if (stablecoin !== null) {
        return { ...baseFixture, stablecoin } as MarketplaceFixture
    } else {
        return { ...baseFixture } as CreditMarketplaceFixture
    }
}