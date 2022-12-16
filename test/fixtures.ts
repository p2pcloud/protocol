import { ethers, upgrades } from "hardhat";

import type { Contract, BigNumberish, BytesLike } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { BrokerV1, Token } from "../typechain-types";
import { PromiseOrValue } from "../typechain-types/common";
import bs58 from 'bs58'

export type Fixture = {
    broker: BrokerV1,
    token: Token
    miner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
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
    return await _deployBrokerFixture()
}

async function _deployBrokerFixture(): Promise<Fixture> {
    const [admin, miner, user, anotherUser] = await ethers.getSigners();

    const Broker = await ethers.getContractFactory("BrokerV1");
    const broker = await upgrades.deployProxy(Broker) as BrokerV1;

    const Token = await ethers.getContractFactory("Token");
    const token = await Token.connect(admin).deploy('10000000000000000000000');

    await broker.SetCommunityContract(admin.address)
    await broker.SetCoinAddress(token.address)

    return { broker, token, miner, user, admin, anotherUser };
}

export async function brokerWithOfferAndUserBalance(): Promise<Fixture> {
    const fixture = await _brokerWithFiveOffers()
    const { user, admin, token, broker, anotherUser } = fixture

    const amt = '10000000'
    await token.connect(admin).transfer(user.address, amt)
    await token.connect(user).approve(broker.address, amt)
    await broker.connect(user).DepositCoin(amt)

    await token.connect(admin).transfer(anotherUser.address, amt)
    await token.connect(anotherUser).approve(broker.address, amt)
    await broker.connect(anotherUser).DepositCoin(amt)

    return fixture
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


export function bookingFromRaw(bookingRaw: any[]) {
    const [index, deprecated__vmTypeId, miner, user, pricePerSecond, bookedAt, lastPayment, offerIndex] = bookingRaw
    return { index, deprecated__vmTypeId, miner, user, pricePerSecond, bookedAt, lastPayment, offerIndex }
}

const specCid = "QmYnq93f9NJ1aCBLCoboncFE6GSZJDqn5RCDVV3ywziXd9"
export const exampleSpecBytes = "0x" + Buffer.from(bs58.decode(specCid).slice(2)).toString('hex')

export async function brokerWithFiveOffers(): Promise<Fixture> {
    return await _brokerWithFiveOffers()
}

async function _brokerWithFiveOffers(): Promise<Fixture> {
    const fixture = await _deployBrokerFixture()

    const { broker, miner } = fixture

    const offers: OffersItem[] = [
        [1, 1, exampleSpecBytes],
        [2, 2, exampleSpecBytes],
        [3, 3, exampleSpecBytes],
        [4, 4, exampleSpecBytes],
        [5, 5, exampleSpecBytes],
    ]

    await Promise.all(offers.map(offer => broker.connect(miner).AddOffer(...offer)))

    return fixture;
}