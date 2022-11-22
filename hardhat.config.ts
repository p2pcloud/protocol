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
                runs: 1000,
            },
        },
    },
    networks: {}
};

if (process.env.HARDAT_PRIVATE_KEY) {
    const fuji = {
        url: process.env.HARDAT_FUJI_ENDPOINT,
        accounts: [process.env.HARDAT_PRIVATE_KEY || ""]
    }
    if (config.networks) {
        config.networks.fuji_testnet = fuji;
        config.networks.fuji_staging = fuji;
    }
}

export default config;
