import { ethers, upgrades } from "hardhat";

import type { Contract, BigNumberish } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { BrokerV1, Token } from "../typechain-types";
import { PromiseOrValue } from "../typechain-types/common";
import bs58 from 'bs58'

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
    machinesAvailable: PromiseOrValue<BigNumberish>,
    specsIpfsHash: PromiseOrValue<BytesLike>,
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

    await Promise.all(offers.map(async (offer) => {
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

const specCid = "QmYnq93f9NJ1aCBLCoboncFE6GSZJDqn5RCDVV3ywziXd9"
export const exampleSpecBytes = "0x" + Buffer.from(bs58.decode(specCid).slice(2)).toString('hex')

// export async function brokerWithFiveOffers(): Promise<Fixture> {
//     const { broker, token, miner, user } = await deployOffersFixture()


//     const offers: OffersItem[] = [
//         [1, 1, exampleSpecBytes],
//         [2, 2, exampleSpecBytes],
//         [3, 3, exampleSpecBytes],
//         [4, 4, exampleSpecBytes],
//         [5, 5, exampleSpecBytes],
//     ]

//     await Promise.all(offers.map(offer => broker.AddOffer(...offer)))

//     return { broker, token, miner, user };
// }