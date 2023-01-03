import { ethers, upgrades } from "hardhat";

const PROXY_ADDRESS = "0x751572bE5Ce0730829e4Ae99D70C0dAAD6D2E291";

async function main() {
    if (PROXY_ADDRESS === "") {
        const Broker = await ethers.getContractFactory("BrokerV1");
        const broker = await upgrades.deployProxy(Broker);
        await broker.deployed();
        console.log("BrokerV1 deployed to:", broker.address);
    } else {
        const Contract = await ethers.getContractFactory("BrokerV1");
        await upgrades.upgradeProxy(PROXY_ADDRESS, Contract);
        console.log("Contract upgraded");
    }
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
