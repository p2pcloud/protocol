import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { Marketplace, TestableMarketplace, MockERC20 } from "../typechain-types";

export type Fixture = {
    marketplace: TestableMarketplace,
    token: MockERC20
    provider: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
}

export async function deployMarketplaceFixture(): Promise<Fixture> {
    return await _deployMarketplaceFixture()
}

async function _deployMarketplaceFixture(): Promise<Fixture> {
    const [admin, provider, user, anotherUser] = await ethers.getSigners();

    const Marketplace = await ethers.getContractFactory("TestableMarketplace");
    const marketplace = await upgrades.deployProxy(Marketplace) as TestableMarketplace;

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const token = await MockERC20.connect(admin).deploy('10000000000000000000000');

    // await marketplace.connect(admin).initialize()
    await marketplace.connect(admin).setCoin(token.address)

    const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
    await token.connect(admin).transfer(provider.address, fee)
    await token.connect(provider).increaseAllowance(marketplace.address, fee)
    await marketplace.connect(provider).depositCoin(fee)
    await marketplace.connect(provider).registerProvider()

    return { marketplace, token, provider, user, admin, anotherUser };
}