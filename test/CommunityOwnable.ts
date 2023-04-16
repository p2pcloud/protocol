import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployBrokerFixture } from './fixtures'

describe("CommunityOwnable", function () {
    describe("transferOwnership", function () {
        it("should set community address if community is the sender", async function () {
            const { broker, provider, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            await broker.connect(admin).transferOwnership(anotherUser.address)

            expect(await broker.connect(anotherUser).transferOwnership(provider.address)).to.not.be.reverted
            expect(await broker.owner()).to.equal(provider.address);
        });
        it("should revert if not owner", async function () {
            const { broker, provider, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            //FIXME: this is a hack to not make a separate fixture
            await broker.connect(admin).transferOwnership(provider.address)

            expect(broker.connect(anotherUser).transferOwnership(admin.address)).to.be.reverted
            expect(await broker.owner()).to.equal(provider.address);
        });
    })
    describe("SetCommunityFee", function () {
        it("should set community fee", async function () {
            const { broker, admin } = await loadFixture(deployBrokerFixture);

            await expect(broker.connect(admin).SetCommunityFee(2000)).to.not.be.reverted
            expect(await broker.communityFee()).to.equal(2000);
        });
        it("should revert if not community contract", async function () {
            const { broker, user } = await loadFixture(deployBrokerFixture);
            await broker.transferOwnership(user.address)
            await expect(broker.SetCommunityFee(2000)).to.be.reverted
        });
        it("should revert if fee is not between 0 (0%) and 2500 (25%)", async function () {
            const { broker, admin } = await loadFixture(deployBrokerFixture);
            await expect(broker.connect(admin).SetCommunityFee(0)).to.not.be.reverted
            await expect(broker.connect(admin).SetCommunityFee(2499)).to.not.be.reverted
            await expect(broker.connect(admin).SetCommunityFee(2500)).to.be.reverted
        });

        it.skip("should change amount of fee paid in ClaimPament ", async function () {
            //TODO: implement
        });
    })
})
