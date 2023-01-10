import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import '@openzeppelin/hardhat-upgrades';
import dotenv from "dotenv";
dotenv.config();



const config: HardhatUserConfig = {
    solidity: {
        version: "0.8.17",
        settings: {
            optimizer: {
                enabled: true,
                runs: 500000,
            },
        },
    },
    networks: {},
};

if (process.env.HARDAT_PRIVATE_KEY) {
    const fuji = {
        url: "https://api.avax-test.network/ext/bc/C/rpc",
        accounts: [process.env.HARDAT_PRIVATE_KEY || ""]
    }
    if (config.networks) {
        config.networks.fuji_testnet = fuji;
        config.networks.fuji_staging = fuji;

        config.networks.ava_mainnet = {
            url: "https://api.avax.network/ext/bc/C/rpc",
            accounts: [process.env.HARDAT_PRIVATE_KEY || ""]
        }
    }

}

export default config;
