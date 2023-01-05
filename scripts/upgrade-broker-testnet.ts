import { ethers, upgrades } from "hardhat";

let PROXY_ADDRESS = "0x935df18a119F5CfD13d12c1cAaF74c004869592F";

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
