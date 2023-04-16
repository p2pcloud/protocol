import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployBrokerFixture } from './fixtures'

describe("BalanceHolder", function () {
    describe("depositCoin", function () {
        it("should incrase balance", async function () {
            const { broker, token, provider, user, admin } = await loadFixture(deployBrokerFixture);

            await token.connect(admin).transfer(user.address, '123456')

            const [userBalance1] = await broker.connect(user).getCoinBalance(user.address)
            expect(userBalance1.toString()).is.equal('0')

            await token.connect(user).approve(broker.address, '123456')
            await broker.connect(user).depositCoin('123')

            const [userBalance2] = await broker.connect(user).getCoinBalance(user.address)
            expect(userBalance2.toString()).is.equal('123')
        });

        it("should revert if transfer fails", async function () {
            const { broker, token, provider, user, admin } = await loadFixture(deployBrokerFixture);
            await expect(broker.connect(user).depositCoin('123')).to.be.reverted
        });
    })
    describe("WithdrawCoin", function () {
        it("should not withdraw locked balance", async function () {
            const { broker, token, provider, user, admin } = await loadFixture(deployBrokerFixture);

            //deposit 123 tokens
            await token.connect(admin).transfer(user.address, '123456')
            await token.connect(user).approve(broker.address, '123456')
            await broker.connect(user).depositCoin('123')

            //lock 23 tokens
            await broker.test__increaseLockedBalance(user.address, '23')

            //fail to withdraw 101 tokens
            await expect(broker.connect(user).withdrawCoin('101')).to.be.reverted

            //withdraw 100 tokens
            await broker.connect(user).withdrawCoin('100')
        });
        it("should revert if transfer fails", async function () {
            //TODO: modify coin contract to make this test pass
            const { broker, token, provider, user, admin } = await loadFixture(deployBrokerFixture);

            const [userBalance1] = await broker.connect(user).getCoinBalance(user.address)

            await token.transferShouldFail(true)
            await expect(broker.connect(user).withdrawCoin(1)).to.be.reverted

            const [userBalance2] = await broker.connect(user).getCoinBalance(user.address)
            expect(userBalance2.toString()).is.equal(userBalance1.toString())

        });
    })
    describe("getCoinBalance", function () {
        it("should return locked and free balance", async function () {
            const { broker, token, provider, user, admin } = await loadFixture(deployBrokerFixture);

            //deposit 123 tokens
            await token.connect(admin).transfer(user.address, '123456')
            await token.connect(user).approve(broker.address, '123456')
            await broker.connect(user).depositCoin('123')

            //lock 23 tokens
            await broker.test__increaseLockedBalance(user.address, '23')

            const [free, locked] = await broker.connect(user).getCoinBalance(user.address)
            expect(free.toString()).is.equal('100')
            expect(locked.toString()).is.equal('23')
        });
    })
    describe("SetCoinAddress", function () {
        it("should set coin address", async function () {
            const { broker, token, admin } = await loadFixture(deployBrokerFixture);

            await broker.connect(admin).setCoin(token.address)
            expect(await broker.coin()).is.equal(token.address)
        });
        it("should revert if not owner", async function () {
            const { broker, token, provider, user } = await loadFixture(deployBrokerFixture);

            await expect(broker.connect(user).setCoin(token.address)).to.be.reverted
        });
    })
});