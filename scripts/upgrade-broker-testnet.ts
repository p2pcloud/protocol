import { ethers, upgrades } from "hardhat";

const PROXY_ADDRESS = "0x951e84e6e252355724bf59F741f5CAEf89CA2Bc2";

async function main() {
    const Contract = await ethers.getContractFactory("BrokerV1");
    const instance = await upgrades.upgradeProxy(PROXY_ADDRESS, Contract);
    console.log("Contract upgraded");
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
