import { ethers, upgrades } from "hardhat";

export async function deployBrokerFixture() {
    const [owner, otherAccount] = await ethers.getSigners();

    const Broker = await ethers.getContractFactory("BrokerV1");
    const broker = await upgrades.deployProxy(Broker);

    return { broker, owner, otherAccount };
}