import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import '@openzeppelin/hardhat-upgrades';
import "hardhat-contract-sizer"
import dotenv from "dotenv";
dotenv.config();

const config: HardhatUserConfig = {
    solidity: {
        version: "0.8.17",
        settings: {
            optimizer: {
                enabled: true,
                runs: 3000,
            },
        },
    },
    networks: {},
    contractSizer: {
        alphaSort: true,
        disambiguatePaths: false,
        runOnCompile: true,
        strict: true,
        only: [':Marketplace$'],
    }
};

if (config.networks) {
    config.networks.ganache = {
        url: "http://localhost:8545",
        accounts: ["0x6be878a6ef0718de9c50c3119c18aa71d64028523093c8415edf35e071e0fc34"] // 1st account from ganache
    }
}

if (process.env.HARDHAT_PRIVATE_KEY) {
    if (!config.networks) throw new Error("typescript is ugly");

    const fuji = {
        url: "https://api.avax-test.network/ext/bc/C/rpc",
        accounts: [process.env.HARDHAT_PRIVATE_KEY || ""]
    }
    const private_subnet = {
        url: "https://subnet.p2pcloud.io",
        accounts: [process.env.HARDHAT_PRIVATE_KEY || ""]
    }

    config.networks.fiat_dev = private_subnet;
    config.networks.fiat_staging = private_subnet;
    config.networks.fiat_preview = private_subnet;

} else {
    console.log("HARDHAT_PRIVATE_KEY is empty")
}

export default config;
