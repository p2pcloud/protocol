import { ethers, upgrades } from "hardhat";
import prompt from "prompt";


if (String(process.env.HARDAT_PRIVATE_KEY).length < 20) {
    console.log("env HARDAT_PRIVATE_KEY is empty");
    process.exit(1);
}

const CONTRACT_TO_DEPLOY = "Marketplace";
const { COIN_ADDRESS, PROXY_ADDRESS } = process.env;

async function main() {
    if (!PROXY_ADDRESS) {
        console.log(`env PROXY_ADDRESS is empty. Type 'deploy' to deploy a new ${CONTRACT_TO_DEPLOY} contract`);
        const { response } = await prompt.get(['response']);
        if (response !== "deploy") {
            console.log("Aborting");
            return;
        }

        const Broker = await ethers.getContractFactory(CONTRACT_TO_DEPLOY);
        const broker = await upgrades.deployProxy(Broker, [COIN_ADDRESS]);
        await broker.deployed();
        console.log(CONTRACT_TO_DEPLOY + " deployed to:", broker.address);
    } else {
        console.log(`env PROXY_ADDRESS is ${PROXY_ADDRESS} \n Type 'upgrade' to deploy a new ${CONTRACT_TO_DEPLOY} contract`);
        const { response } = await prompt.get(['response']);
        if (response !== "upgrade") {
            console.log("Aborting");
            return;
        }

        const Contract = await ethers.getContractFactory(CONTRACT_TO_DEPLOY);
        await upgrades.upgradeProxy(PROXY_ADDRESS, Contract);
        console.log("Contract upgraded");
    }
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
