import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'

describe("Broker_getMinerUrl", function () {
    it("should set miner url", async function () {
        const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

        const fee = await broker.MINER_REGISTRATION_FEE()
        await token.connect(admin).transfer(anotherUser.address, fee)
        await token.connect(anotherUser).increaseAllowance(broker.address, fee)
        await broker.connect(anotherUser).DepositCoin(fee)
        await broker.connect(anotherUser).RegisterMiner()

        const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
        await broker.connect(anotherUser).SetMinerUrl(urlBytes)
        const url = await broker.connect(anotherUser).GetMinerUrl(anotherUser.address);
        expect(url).to.equal(urlBytes);
    });

    it("should not set miner url for non-registered miner", async function () {
        const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);


        const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
        expect(broker.connect(anotherUser).SetMinerUrl(urlBytes)).to.be.revertedWith("Miner not registered")

        const fee = await broker.MINER_REGISTRATION_FEE()
        await token.connect(admin).transfer(anotherUser.address, fee)
        await token.connect(anotherUser).increaseAllowance(broker.address, fee)
        await broker.connect(anotherUser).DepositCoin(fee)
        await broker.connect(anotherUser).RegisterMiner()

        await broker.connect(anotherUser).SetMinerUrl(urlBytes)
    });
});


describe("Broker_IsMinerRegistered", function () {
    it("should return true for registered and false by default", async function () {
        const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

        const isregistered1 = await broker.connect(anotherUser).IsMinerRegistered(anotherUser.address)
        expect(isregistered1).to.equal(false)

        const fee = await broker.MINER_REGISTRATION_FEE()
        await token.connect(admin).transfer(anotherUser.address, fee)
        await token.connect(anotherUser).increaseAllowance(broker.address, fee)
        await broker.connect(anotherUser).DepositCoin(fee)
        await broker.connect(anotherUser).RegisterMiner()

        const isregistered2 = await broker.connect(anotherUser).IsMinerRegistered(anotherUser.address)
        expect(isregistered2).to.equal(true)
    })
})


describe("Broker_RegisterMiner", function () {
    it("should use exactly MINER_REGISTRATION_FEE coins", async function () {
        const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

        const fee = await broker.MINER_REGISTRATION_FEE()
        await token.connect(admin).transfer(anotherUser.address, fee.add(123))
        await token.connect(anotherUser).increaseAllowance(broker.address, fee.add(123))
        await broker.connect(anotherUser).DepositCoin(fee.add(123))


        const balance1 = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
        expect(balance1[0]).to.equal(fee.add(123))
        expect(balance1[1]).to.equal(0)

        await broker.connect(anotherUser).RegisterMiner()

        const balance2 = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
        expect(balance2[0]).to.equal(123)
        expect(balance2[1]).to.equal(0)
    });

    it("should reject if not enough coins", async function () {
        const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

        const fee = await broker.MINER_REGISTRATION_FEE()
        await token.connect(admin).transfer(anotherUser.address, fee.sub(1))
        await token.connect(anotherUser).increaseAllowance(broker.address, fee.sub(1))
        await broker.connect(anotherUser).DepositCoin(fee.sub(1))

        await expect(broker.connect(anotherUser).RegisterMiner()).to.be.rejected
    });
});
