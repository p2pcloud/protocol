import { ethers, upgrades } from "hardhat";

async function main() {
    const Broker = await ethers.getContractFactory("BrokerV1");
    const broker = await upgrades.deployProxy(Broker);
    await broker.deployed();
    console.log("BrokerV1 deployed to:", broker.address);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
