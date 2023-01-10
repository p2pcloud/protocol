import { ethers, upgrades } from "hardhat";

import type { Contract, BigNumberish, BytesLike } from "ethers";
import { BigNumber } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { Broker, Token } from "../typechain-types";
import { PromiseOrValue } from "../typechain-types/common";
import bs58 from 'bs58'

export type Fixture = {
    broker: Broker,
    token: Token
    miner: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
}

export async function deployBrokerFixture(): Promise<Fixture> {
    return await _deployBrokerFixture()
}


async function _deployBrokerFixture(): Promise<Fixture> {
    const [admin, miner, user, anotherUser] = await ethers.getSigners();

    const BrokerContract = await ethers.getContractFactory("Broker");
    const broker = await upgrades.deployProxy(BrokerContract) as Broker;

    //TODO: test upgradability

    const Token = await ethers.getContractFactory("Token");
    const token = await Token.connect(admin).deploy('10000000000000000000000');

    await broker.SetCommunityContract(admin.address)
    await broker.SetCoinAddress(token.address)

    const fee = await broker.MINER_REGISTRATION_FEE()
    await token.connect(admin).transfer(miner.address, fee)
    await token.connect(miner).increaseAllowance(broker.address, fee)
    await broker.connect(miner).DepositCoin(fee)
    await broker.connect(miner).RegisterMiner()

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

const specCid = "QmYnq93f9NJ1aCBLCoboncFE6GSZJDqn5RCDVV3ywziXd9"
export const exampleSpecBytes = "0x" + Buffer.from(bs58.decode(specCid).slice(2)).toString('hex')

export async function brokerWithFiveOffers(): Promise<Fixture> {
    return await _brokerWithFiveOffers()
}

async function _brokerWithFiveOffers(): Promise<Fixture> {
    const fixture = await _deployBrokerFixture()

    const { broker, miner } = fixture

    type OfferType = [number, number, BytesLike]

    const offers: OfferType[] = [
        [1, 1, exampleSpecBytes],
        [2, 2, exampleSpecBytes],
        [3, 3, exampleSpecBytes],
        [4, 4, exampleSpecBytes],
        [5, 5, exampleSpecBytes],
    ]

    await Promise.all(offers.map(offer => broker.connect(miner).AddOffer(...offer)))

    return fixture;
}