import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { deployBrokerFixture } from './fixtures'

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
        it("should set community fee");
        it("should change amount of fee paid in ClaimPament");
        it("should revert if not community contract");
        it("should revert if fee is not between 0 (0%) and 2500 (25%)");
    })
})