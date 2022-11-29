import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";

import { expect } from "chai";
import { ethers } from "hardhat";
import { deployBrokerFixture, offers, OffersItem } from './fixtures'

describe("BrokerV1_community", function () {

    describe("SetCommunityAddress", function () {
        it("should set community address if address is address(0)", async function () {
            const { broker, user } = await loadFixture(deployBrokerFixture);

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(user.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(user.address);
        });
        it("should set community address if community is the sender", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(miner.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(miner.address);

            expect(await broker.SetCommunityContract(user.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(user.address);
        });
        it("should revert if not owner", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(miner.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(miner.address);

            await expect(broker.connect(user).SetCommunityContract(miner.address)).to.be.reverted
        });
    })
    describe("SetCommunityFee", function () {
        it("should set community fee", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);
            await broker.SetCommunityContract(miner.address)
            await expect(broker.SetCommunityFee(2000)).to.not.be.reverted
        });
        it("should change amount of fee paid in ClaimPament", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);
            const FEE = 2000
            const week = 86400

            // Contract settings
            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)
            await broker.SetCommunityFee(FEE)

            // Add User offers
            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.connect(user).AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            // Allow broker deposit User tokens
            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            // Deposit & book
            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            // Waiting week
            await time.increase(week);

            // Get booking info
            const [[index, pricePerSecond]] = await broker.FindBookingsByUser(user.address)
            const pricePerWeek = ethers.utils.formatUnits(pricePerSecond.mul(week), 'mwei')

            // Claim payment
            expect(await broker.connect(user).ClaimPayment(index)).not.to.be.reverted

            // Get balance & calc percents
            const [free, locked] = await broker.GetCoinBalance(miner.address)
            const balance = ethers.utils.formatUnits(free, 'mwei')
            const percents = (+pricePerWeek * FEE) / 10000

            expect(Number(balance)).is.equal(percents)
        });
        it("should revert if not community contract", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);
            await broker.SetCommunityContract(user.address)
            await expect(broker.SetCommunityFee(2000)).to.be.reverted
        });
        it("should revert if fee is not between 0 (0%) and 2500 (25%)", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);
            await broker.SetCommunityContract(miner.address)
            await expect(broker.SetCommunityFee(2500)).to.be.reverted
        });
    })
})