import { expect } from "chai";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { brokerWithFiveOffers, brokerWithOfferAndUserBalance, deployBrokerFixture } from './fixtures'
import { BigNumber, ethers } from "ethers";

describe("Broker_trials", function () {
    describe("AddTrial", function () {
        it("should set trial allowance", async () => {
            const { broker, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            const { free, locked } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free).to.equal(0)
            expect(locked).to.equal(0)

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("123456789"))

            const { free: free2, locked: locked2 } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free2).to.equal("123456789")
            expect(locked2).to.equal(0)
        });
        it("should block withdrawal", async () => {
            const { broker, token, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await token.connect(admin).transfer(anotherUser.address, BigNumber.from("123456789"))
            await token.connect(anotherUser).approve(broker.address, BigNumber.from("123456789"))
            await broker.connect(anotherUser).DepositCoin(BigNumber.from("123456789"))

            const { free, locked } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free).to.equal("123456789")
            expect(locked).to.equal(0)

            await broker.connect(anotherUser).WithdrawCoin(BigNumber.from("789"))
            const { free: free2, locked: locked2 } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free2).to.equal("123456000")
            expect(locked2).to.equal(0)

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("777"))
            const { free: free3, locked: locked3 } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free3).to.equal("123456777")
            expect(locked3).to.equal(0)

            //reverted
            await expect(broker.connect(anotherUser).WithdrawCoin(BigNumber.from("789"))).to.be.reverted

            //disable trial
            await broker.connect(anotherUser).RemoveTrial(anotherUser.address)

            //not reverted
            await broker.connect(anotherUser).WithdrawCoin(BigNumber.from("789"))
        });
        it("should still allow deposits", async () => {
            const { broker, token, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await token.connect(admin).transfer(anotherUser.address, BigNumber.from("123456789"))
            await token.connect(anotherUser).approve(broker.address, BigNumber.from("123456789"))

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("777"))

            await broker.connect(anotherUser).DepositCoin(BigNumber.from("1000"))

            const { free, locked } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free).to.equal("1777")
            expect(locked).to.equal(0)

        });
        it("should fail if not community contract", async () => {
            const { broker, token, user, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await expect(broker.connect(user).AddTrial(anotherUser.address, BigNumber.from("777"))).to.be.reverted
            await expect(broker.connect(anotherUser).AddTrial(anotherUser.address, BigNumber.from("777"))).to.be.reverted
            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("777"))
        });
        it("should top up balance", async () => {
            const { broker, token, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await token.connect(admin).transfer(anotherUser.address, BigNumber.from("123456789"))
            await token.connect(anotherUser).approve(broker.address, BigNumber.from("123456789"))

            await broker.connect(anotherUser).DepositCoin(BigNumber.from("1000"))
            const { free, locked } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free).to.equal("1000")
            expect(locked).to.equal(0)


            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("777"))

            const { free: free2, locked: locked2 } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)

            expect(free2).to.equal("1777")
            expect(locked2).to.equal(0)
        })
        it("should top up if called twice", async () => {
            const { broker, token, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("300"))
            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("500"))

            const { free, locked } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free).to.equal("800")
            expect(locked).to.equal(0)
        })
    })
    describe("DisableTrial", function () {
        it("should decrease balance and remove trial mark", async () => {
            const { broker, token, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("300"))

            const { free, locked } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free).to.equal("300")
            expect(locked).to.equal(0)

            await broker.connect(anotherUser).RemoveTrial(anotherUser.address)

            const { free: free2, locked: locked2 } = await broker.connect(anotherUser).GetCoinBalance(anotherUser.address)
            expect(free2).to.equal(0)
            expect(locked2).to.equal(0)
        });
        it("should be only called by user or community contract", async () => {
            const { broker, token, user, admin, anotherUser } = await loadFixture(deployBrokerFixture);

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("300"))
            await expect(broker.connect(user).RemoveTrial(anotherUser.address)).to.be.reverted
            await broker.connect(anotherUser).RemoveTrial(anotherUser.address)

            await broker.connect(admin).AddTrial(anotherUser.address, BigNumber.from("300"))
            await expect(broker.connect(user).RemoveTrial(anotherUser.address)).to.be.reverted
            await broker.connect(admin).RemoveTrial(anotherUser.address)
        });
    })
})
