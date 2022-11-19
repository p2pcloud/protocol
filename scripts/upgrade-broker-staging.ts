import { ethers, upgrades } from "hardhat";

const PROXY_ADDRESS = "0xfe1146269737d9e4BA7BE0277c233C356af1a234";

async function main() {
    const Contract = await ethers.getContractFactory("BrokerV1");
    const instance = await upgrades.upgradeProxy(PROXY_ADDRESS, Contract);
    console.log("Contract upgraded");
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
