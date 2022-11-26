import { ethers, upgrades } from "hardhat";

import type { Contract, BigNumberish } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { BrokerV1, Token } from "../typechain-types";
import { PromiseOrValue } from "../typechain-types/common";

export type Fixture = {
    broker: BrokerV1,
    token: Token
    miner: SignerWithAddress,
    user: SignerWithAddress,
}

export type Offer = {
    Index: number,
    Miner: string,
    PPS: number,
    Availablility: number,
    VmTypeId: number
}

export type OffersItem = [
    pricePerSecond: PromiseOrValue<BigNumberish>,
    vmTypeId: PromiseOrValue<BigNumberish>,
    machinesAvailable: PromiseOrValue<BigNumberish>
]

export async function deployBrokerFixture(): Promise<Fixture> {
    const [miner, user] = await ethers.getSigners();

    const Broker = await ethers.getContractFactory("BrokerV1");
    const broker = await upgrades.deployProxy(Broker) as BrokerV1;

    const Token = await ethers.getContractFactory("Token");
    const token = await Token.deploy('10000000000000000000000');

    return { broker, token, miner, user };
}

export async function deployOffersFixture(): Promise<Fixture> {
    const [miner, user] = await ethers.getSigners();

    const Broker = await ethers.getContractFactory("BrokerV1");
    const broker = await upgrades.deployProxy(Broker) as BrokerV1;

    const txs = await Promise.all(offers.map(async (offer) => {
        const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
        return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
    }))

    const Token = await ethers.getContractFactory("Token");
    const token = await Token.deploy('10000000000000000000000');

    return { broker, token, miner, user };
}

export function offerFromRaw(offerRaw: any[]) {
    const [Index, Miner, PPS, Availablility, VmTypeId] = offerRaw

    return {
        VmTypeId: VmTypeId.toNumber(),
        PPS: PPS.toNumber(),
        Availablility: Availablility.toNumber(),
        Miner: Miner,
        Index: Index.toNumber(),
    }
}

export const offers: OffersItem[] = [
    [1, 1, 1],
    [2, 2, 2],
    [3, 3, 3],
    [4, 4, 4],
    [5, 5, 5],
]