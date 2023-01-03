import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture, brokerWithFiveOffers, brokerWithOfferAndUserBalance } from './fixtures'
import { BigNumber } from "ethers";

describe("BrokerV1_coin", function () {
    describe("DepositCoin", function () {
        it("should incrase balance", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(deployBrokerFixture);

            await token.connect(admin).transfer(user.address, '123456')

            const [userBalance1] = await broker.connect(user).GetCoinBalance(user.address)
            expect(userBalance1.toString()).is.equal('0')

            await token.connect(user).approve(broker.address, '123456')
            await broker.connect(user).DepositCoin('123')

            const [userBalance2] = await broker.connect(user).GetCoinBalance(user.address)
            expect(userBalance2.toString()).is.equal('123')
        });

        it("should revert if transfer fails", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(deployBrokerFixture);
            await expect(broker.connect(user).DepositCoin('123')).to.be.reverted
        });
    })
    describe("GetLockedCoinBalance", function () {
        it("should increase with vm booking", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            const [free1, locked1] = await broker.connect(user).GetCoinBalance(user.address)

            const PPS = (await broker.GetOffer(0)).pricePerSecond
            await broker.connect(user).Book(0)

            const expectedIncrease = BigNumber.from(PPS).mul(60 * 60 * 24 * 7)

            const [free2, locked2] = await broker.connect(user).GetCoinBalance(user.address)
            expect(locked2.sub(locked1)).is.equal(expectedIncrease)
            expect(free1.sub(free2)).is.equal(expectedIncrease)
        });
        it("should decrease with vm termination", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            const [free1, locked1] = await broker.connect(user).GetCoinBalance(user.address)
            expect(locked1.toString()).is.equal('0')

            await broker.connect(user).Book(0)

            const [free2, locked2] = await broker.connect(user).GetCoinBalance(user.address)
            expect(locked2.toString()).is.not.equal('0')

            await broker.connect(user).Terminate(0, 0)

            const [free3, locked3] = await broker.connect(user).GetCoinBalance(user.address)
            expect(locked3.toString()).is.equal('0')
        });
    })
    describe("WithdrawCoin", function () {
        it("should withdraw only free balance", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)
            const [free, locked] = await broker.connect(user).GetCoinBalance(user.address)

            await expect(broker.connect(user).WithdrawCoin(free.add(1))).to.be.reverted
            await broker.connect(user).WithdrawCoin(free)
        });
        it("should revert if transfer fails", async function () {
            //TODO: modify coin contract to make this test pass
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await token.transferShouldFail(true)
            await expect(broker.connect(user).WithdrawCoin(1)).to.be.reverted
        });
    })
    describe("GetCoinBalance", function () {
        it("should return locked and free balance", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const [free, locked] = await broker.connect(user).GetCoinBalance(user.address)
            expect(free.toString()).is.not.equal('0')
            expect(locked.toString()).is.not.equal('0')
        });
    })
    describe("SetCoinAddress", function () {
        it("should set coin address", async function () {
            const { broker, token, admin } = await loadFixture(deployBrokerFixture);

            await broker.connect(admin).SetCoinAddress(token.address)
            expect(await broker.coin()).is.equal(token.address)
        });
        it("should revert if not owner", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await expect(broker.connect(user).SetCoinAddress(token.address)).to.be.reverted
        });
    })
});
