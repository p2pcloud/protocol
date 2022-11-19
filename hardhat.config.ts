import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import '@openzeppelin/hardhat-upgrades';
import dotenv from "dotenv";
dotenv.config();

const fuji = {
    url: process.env.HARDAT_FUJI_ENDPOINT,
    accounts: [process.env.HARDAT_PRIVATE_KEY || ""]
}

const config: HardhatUserConfig = {
    solidity: "0.8.17",
    networks: {
        fuji_staging: fuji,
        fuji_testnet: fuji,
    }
};

export default config;
