import { ethers, upgrades } from "hardhat";

import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { Broker, TestableBroker, MockERC20 } from "../typechain-types";

export type Fixture = {
    broker: TestableBroker,
    token: MockERC20
    provider: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
}

export async function deployBrokerFixture(): Promise<Fixture> {
    return await _deployBrokerFixture()
}

async function _deployBrokerFixture(): Promise<Fixture> {
    const [admin, provider, user, anotherUser] = await ethers.getSigners();

    const Broker = await ethers.getContractFactory("TestableBroker");
    const broker = await upgrades.deployProxy(Broker) as TestableBroker;

    const MockERC20 = await ethers.getContractFactory("MockERC20");
    const token = await MockERC20.connect(admin).deploy('10000000000000000000000');

    // await broker.connect(admin).initialize()
    await broker.connect(admin).setCoin(token.address)

    const fee = await broker.PROVIDER_REGISTRATION_FEE()
    await token.connect(admin).transfer(provider.address, fee)
    await token.connect(provider).increaseAllowance(broker.address, fee)
    await broker.connect(provider).depositCoin(fee)
    await broker.connect(provider).registerProvider()

    return { broker, token, provider, user, admin, anotherUser };
}