import { ethers, upgrades } from "hardhat";

import type { Contract, BigNumberish, BytesLike } from "ethers";
import { BigNumber } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import type { Broker, OldBroker, Token } from "../typechain-types";
import { PromiseOrValue } from "../typechain-types/common";
import bs58 from 'bs58'

export type Fixture = {
    broker: Broker,
    token: Token
    provider: SignerWithAddress,
    user: SignerWithAddress,
    admin: SignerWithAddress,
    anotherUser: SignerWithAddress,
}

export async function deployBrokerFixture(): Promise<Fixture> {
    return await _deployBrokerFixture()
}


async function _deployBrokerFixture(): Promise<Fixture> {
    const [admin, provider, user, anotherUser] = await ethers.getSigners();

    const OldBrokerContract = await ethers.getContractFactory("OldBroker");
    const oldBroker = await upgrades.deployProxy(OldBrokerContract) as OldBroker;

    const Token = await ethers.getContractFactory("Token");
    const token = await Token.connect(admin).deploy('10000000000000000000000');

    await oldBroker.SetCommunityContract(admin.address)
    await oldBroker.SetCoinAddress(token.address)

    const fee = await oldBroker.PROVIDER_REGISTRATION_FEE()
    await token.connect(admin).transfer(provider.address, fee)
    await token.connect(provider).increaseAllowance(oldBroker.address, fee)
    await oldBroker.connect(provider).DepositCoin(fee)
    await oldBroker.connect(provider).RegisterProvider()

    //upgrade broker
    const Broker = await ethers.getContractFactory("Broker");
    const broker = await upgrades.upgradeProxy(oldBroker.address, Broker) as Broker;

    return { broker, token, provider, user, admin, anotherUser };
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

    const { broker, provider } = fixture

    type OfferType = [number, number, BytesLike]

    const offers: OfferType[] = [
        [1, 1, exampleSpecBytes],
        [2, 2, exampleSpecBytes],
        [3, 3, exampleSpecBytes],
        [4, 4, exampleSpecBytes],
        [5, 5, exampleSpecBytes],
    ]

    await Promise.all(offers.map(offer => broker.connect(provider).AddOffer(...offer)))

    return fixture;
}