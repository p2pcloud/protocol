import { ethers, upgrades } from "hardhat";

import type { Contract } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { Token } from "../typechain-types";

export async function deployBrokerFixture(): Promise<Fixture> {
    const [miner, user] = await ethers.getSigners();

    const Broker = await ethers.getContractFactory("BrokerV1");
    const broker = await upgrades.deployProxy(Broker);

    const Token = await ethers.getContractFactory("Token");
    const token = await Token.deploy('10000000000000000000000');

    return [broker, miner, user, token];
}

export type Fixture = [
    broker: Contract,
    miner: SignerWithAddress,
    user: SignerWithAddress,
    token: Token,
]

export type Offer = {
    Index: number,
    Miner: string,
    PPS: number,
    Availablility: number,
    VmTypeId: number
}