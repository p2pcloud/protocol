import { ethers, upgrades } from "hardhat";

let PROXY_ADDRESS = "0x0a1B15B5ce08532114c09Ff9df2a9e7b15C9Ab1e";

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
