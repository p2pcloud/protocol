import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture, deployOffersFixture } from './fixtures'

describe("BrokerV1_coin", function () {
    describe("DepositCoin", function () {
        it("should incrase balance", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('5', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await broker.connect(user).DepositCoin(allowance)

            const [free, locked] = await broker.GetCoinBalance(user.address)
            const total = free.add(locked)
            const balance = ethers.utils.formatUnits(total, 'mwei')
            const equal = ethers.utils.formatUnits(amt, 'mwei')

            expect(balance).is.equal(equal)
        });
        it("should revert if transfer fails", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('5', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await expect(broker.connect(miner).DepositCoin(allowance)).to.be.reverted
        });
    })
    describe("GetLockedCoinBalance", function () {
        it("should increase with vm booking", async function () {
            const { broker, token, miner, user } = await loadFixture(deployOffersFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('5', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await broker.connect(user).DepositCoin(allowance)

            const booked = await broker.connect(user).Book(0)
            const [free, locked] = await broker.connect(user).GetCoinBalance(user.address)

            // const freeBalance = ethers.utils.formatUnits(free, 'mwei')
            const lockedBalance = ethers.utils.formatUnits(locked, 'mwei')

            expect(Number(lockedBalance)).is.not.equal(0)
        });
        it("should decrease with vm termination", async function () {
            const { broker, token, miner, user } = await loadFixture(deployOffersFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('5', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)
            await broker.connect(user).Terminate(0, 0)

            const [free, locked] = await broker.connect(user).GetCoinBalance(user.address)
            const lockedBalance = ethers.utils.formatUnits(locked, 'mwei')

            expect(Number(lockedBalance)).is.equal(0)
        });
    })
    describe("WithdrawCoin", function () {
        it("should withdraw only free balance", async function () {
            const { broker, token, miner, user } = await loadFixture(deployOffersFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            const [free, locked] = await broker.connect(user).GetCoinBalance(user.address)
            const freeBalance = ethers.utils.formatUnits(free, 'mwei')
            const lockedBalance = ethers.utils.formatUnits(locked, 'mwei')

            await expect(broker.connect(user).WithdrawCoin(amt)).to.be.reverted
        });
        it("should revert if transfer fails", async function () {
            const { broker, token, miner, user } = await loadFixture(deployOffersFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await broker.connect(user).DepositCoin(amt)
            await broker.connect(user).Book(0)

            await expect(broker.connect(user).WithdrawCoin(amt)).to.be.reverted

        });
        it("should revert if not enough balance", async function () {
            const { broker, token, miner, user } = await loadFixture(deployOffersFixture);
            const amt = ethers.utils.parseUnits('1', 'mwei')
            await expect(broker.connect(user).WithdrawCoin(amt)).to.be.reverted
        });
    })
    describe("GetCoinBalance", function () {
        it("should return locked and free balance", async function () {
            const { broker, token, miner, user } = await loadFixture(deployOffersFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            const [free, locked] = await broker.GetCoinBalance(user.address)
            const total = free.add(locked)
            const balance = ethers.utils.formatUnits(total, 'mwei')
            const equal = ethers.utils.formatUnits(amt, 'mwei')

            expect(balance).is.equal(equal)
        });
    })
    describe("SetCoinAddress", function () {
        it("should set coin address", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await expect(broker.SetCoinAddress(token.address)).not.to.be.reverted
        });
        it("should revert if not owner", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(user.address)
            await expect(broker.SetCoinAddress(token.address)).to.be.reverted
        });
    })
});
