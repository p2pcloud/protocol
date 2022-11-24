import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture, offerFromRaw, offers, type Offer, OffersItem } from './fixtures'

describe("BrokerV1_offers", function () {

    describe("AddOffers", function () {
        it("should add offers", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            expect(txs.length).is.equal(5)
        });
    })
    describe("UpdateOffer", function () {
        it("should update offer", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.UpdateOffer(0, 10, 1)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw.map(offerFromRaw)
            const updatedOffer = offersObj.find(({ Index }) => Index === 0)

            expect(updatedOffer?.Availablility).is.equal(10)
        });
        it("should revert if offer belongs to another account", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await expect(broker.connect(user).UpdateOffer(0, 0, 1)).to.be.reverted
        });
    })
    describe("RemoveOffer", function () {
        it("should remove offer", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.RemoveOffer(1)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw.map(offerFromRaw)
            const deletedOffer = offersObj.find(({ Index }) => Index === 1)

            expect(deletedOffer).undefined
        });
        it("should revert if offer belongs to another account", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await expect(broker.connect(user).RemoveOffer(1)).to.be.reverted
        });
    })
    describe("GetAvailableOffers", function () {
        it("should return array of offers", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw.map(offerFromRaw)

            expect(offersObj.length).to.equal(txs.length)
        });
        it("should return only offers with available machines", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.UpdateOffer(0, 0, 1)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw.map(offerFromRaw)
            const availability = offersObj.every(({ Availablility }: { Availablility: number }) => Availablility !== 0)

            expect(availability).true
        });
        it("should return only offers with type 5", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            const offersRaw = await broker.GetAvailableOffersByType(5);
            const offersObj: Offer[] = offersRaw.map(offerFromRaw)
            const typed = offersObj.every(({ VmTypeId }) => VmTypeId === 5)

            expect(typed).true
        });
    })
    describe("GetMinerOffers", function () {
        it("should return array of offers", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            const offersRaw = await broker.GetMinersOffers(miner.address);
            const offersObj = offersRaw.map(offerFromRaw)

            expect(offersObj.length).to.equal(txs.length)
        });
        it("should return offers including ones with availability zero", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);

            const txs = await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.UpdateOffer(0, 0, 1)

            const offersRaw = await broker.GetMinersOffers(miner.address);
            const offersObj = offersRaw.map(offerFromRaw)

            expect(offersObj.length).to.equal(txs.length)
        });
    })
})
