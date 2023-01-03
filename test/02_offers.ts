import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture, exampleSpecBytes } from './fixtures'

describe("BrokerV1_offers", function () {
    describe("AddOffers", function () {
        it("should add offers", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            let offersRaw = await broker.GetAvailableOffers();
            const lengthBefore = offersRaw.length

            await broker.AddOffer(2, 10, exampleSpecBytes)

            offersRaw = await broker.GetAvailableOffers();

            expect(offersRaw.length).is.equal(lengthBefore + 1)
        });
    })
    describe("UpdateOffer", function () {
        it("should update offer", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.AddOffer(2, 10, exampleSpecBytes)

            await broker.UpdateOffer(0, 11, 3)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw
            const updatedOffer = offersObj.find(({ index }) => index === 0)

            expect(updatedOffer?.machinesAvailable).is.equal(11)
            expect(updatedOffer?.pricePerSecond).is.equal(3)
        });

        it("should revert if offer belongs to another account", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.AddOffer(2, 10, exampleSpecBytes)

            await expect(broker.connect(user).UpdateOffer(0, 0, 1)).to.be.reverted
        });
    })
    describe("RemoveOffer", function () {
        it("should remove offer", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.AddOffer(2, 10, exampleSpecBytes)
            await broker.AddOffer(3, 10, exampleSpecBytes)

            await broker.RemoveOffer(1)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw
            const deletedOffer = offersObj.find(({ index }) => index === 1)

            expect(deletedOffer).undefined
        });
        it("should revert if offer belongs to another account", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.AddOffer(2, 10, exampleSpecBytes)

            await expect(broker.connect(user).RemoveOffer(0)).to.be.reverted
        });
    })
    describe("GetAvailableOffers", function () {
        it("should return array of offers", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.AddOffer(2, 10, exampleSpecBytes)
            await broker.AddOffer(3, 10, exampleSpecBytes)

            const offers = (await broker.GetAvailableOffers())

            expect(offers.length).to.equal(2)
        });
        it("should return only offers with available machines", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.AddOffer(2, 10, exampleSpecBytes)
            await broker.AddOffer(3, 0, exampleSpecBytes)

            const offers = (await broker.GetAvailableOffers())

            expect(offers.length).to.equal(1)
        });
    })
    describe("GetMinerOffers", function () {
        it("should return array of offers", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(miner).AddOffer(2, 10, exampleSpecBytes)
            await broker.connect(user).AddOffer(3, 10, exampleSpecBytes)

            const offers = (await broker.GetMinersOffers(miner.address))


            expect(offers.length).to.equal(1)
        });
        it("should return offers including ones with availability zero", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(miner).AddOffer(2, 10, exampleSpecBytes)
            await broker.connect(miner).AddOffer(2, 0, exampleSpecBytes)

            const offers = (await broker.GetMinersOffers(miner.address))


            expect(offers.length).to.equal(2)
        });
    })
})
