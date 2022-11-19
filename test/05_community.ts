import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("BrokerV1_community", function () {
    async function deployBrokerFixture() {
        const [owner, otherAccount] = await ethers.getSigners();

        const Broker = await ethers.getContractFactory("BrokerV1");
        const broker = await upgrades.deployProxy(Broker);

        return { broker, owner, otherAccount };
    }

    describe("SetCommunityAddress", function () {
        it("should set community address if address is address(0)", async function () {
            const { broker, owner, otherAccount } = await loadFixture(deployBrokerFixture);

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(otherAccount.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(otherAccount.address);
        });
        it("should set community address if community is the sender", async function () {
            const { broker, owner, otherAccount } = await loadFixture(deployBrokerFixture);

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(owner.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(owner.address);

            expect(await broker.SetCommunityContract(otherAccount.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(otherAccount.address);
        });
        it("should revert if not owner", async function () {
            const { broker, owner, otherAccount } = await loadFixture(deployBrokerFixture);

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(owner.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(owner.address);

            await expect(broker.connect(otherAccount).SetCommunityContract(owner.address)).to.be.reverted
        });
    })
    describe("SetCommunityFee", function () {
        it("should set community fee");
        it("should change amount of fee paid in ClaimPament");
        it("should revert if not community contract");
        it("should revert if fee is not between 0 (0%) and 2500 (25%)");
    })
})