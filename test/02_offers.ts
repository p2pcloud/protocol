import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture, exampleSpecBytes, brokerWithOfferAndUserBalance } from './fixtures'

describe("Broker_offers", function () {
    describe("AddOffers", function () {
        it("should add offers", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            let offersRaw = await broker.GetAvailableOffers();
            const lengthBefore = offersRaw.length

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            offersRaw = await broker.GetAvailableOffers();

            expect(offersRaw.length).is.equal(lengthBefore + 1)
        });

        it("should not add offers for unregistered provider", async function () {
            const { broker, provider, anotherUser } = await loadFixture(deployBrokerFixture);
            await expect(broker.connect(anotherUser).AddOffer(2, 10, exampleSpecBytes)).to.be.reverted
        });
    })
    describe("UpdateOffer", function () {
        it("should update offer", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            await broker.connect(provider).UpdateOffer(0, 11, 3)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw
            const updatedOffer = offersObj.find(({ index }) => index === 0)

            expect(updatedOffer!.machinesTotal).is.equal(11)
            expect(updatedOffer!.pricePerSecond).is.equal(3)
        });

        it("should revert if offer belongs to another account", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            await expect(broker.connect(user).UpdateOffer(0, 0, 1)).to.be.reverted
        });

        it("should revert if total is less than booked", async function () {
            const { broker, provider, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            let offerId = 0
            let machines = 10
            let pps = 1
            await broker.connect(provider).UpdateOffer(offerId, machines, pps)

            //book 3 machines
            await broker.connect(user).Book(0)
            await broker.connect(user).Book(0)
            await broker.connect(user).Book(0)

            //update to 2 machines: fail
            machines = 2
            await expect(broker.connect(provider).UpdateOffer(offerId, machines, pps)).to.be.reverted

            //update to 3 machines: ok
            machines = 3
            await broker.connect(provider).UpdateOffer(offerId, machines, pps)

            //terminate 1 machine and update to 2 machines: ok
            await broker.connect(user).Terminate(0, 0)
            machines = 2
            await broker.connect(provider).UpdateOffer(offerId, machines, pps)

        });

    })
    describe("RemoveOffer", function () {
        it("should remove offer", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)
            await broker.connect(provider).AddOffer(3, 10, exampleSpecBytes)

            await broker.connect(provider).RemoveOffer(1)

            const offersRaw = await broker.GetAvailableOffers();
            const offersObj = offersRaw
            const deletedOffer = offersObj.find(({ index }) => index === 1)

            expect(deletedOffer).undefined
        });
        it("should revert if offer still has machines booked", async function () {
            const { broker, provider, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            await broker.connect(user).Book(0)

            //revert
            await expect(broker.connect(provider).RemoveOffer(0)).to.be.reverted

            //remove booking and try again
            await broker.connect(user).Terminate(0, 0)
            await broker.connect(provider).RemoveOffer(0)
        });
        it("should revert if offer belongs to another account", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            await expect(broker.connect(user).RemoveOffer(0)).to.be.reverted
        });
    })
    describe("GetAvailableOffers", function () {
        it("should return array of offers", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)
            await broker.connect(provider).AddOffer(3, 10, exampleSpecBytes)

            const offers = (await broker.GetAvailableOffers())

            expect(offers.length).to.equal(2)
        });
        it("should return only offers with available machines", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)
            await broker.connect(provider).AddOffer(3, 0, exampleSpecBytes)

            const offers = (await broker.GetAvailableOffers())

            expect(offers.length).to.equal(1)
        });
    })
    describe("GetProviderOffers", function () {
        it("should return array of offers", async function () {
            const { broker, provider, admin, token } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)

            await token.connect(admin).approve(broker.address, broker.PROVIDER_REGISTRATION_FEE())
            await broker.connect(admin).DepositCoin(broker.PROVIDER_REGISTRATION_FEE())
            await broker.connect(admin).RegisterProvider()
            await broker.connect(admin).AddOffer(3, 10, exampleSpecBytes)

            const offers = (await broker.GetProvidersOffers(provider.address))


            expect(offers.length).to.equal(1)
        });
        it("should return offers including ones with availability zero", async function () {
            const { broker, provider, user } = await loadFixture(deployBrokerFixture);

            await broker.connect(provider).AddOffer(2, 10, exampleSpecBytes)
            await broker.connect(provider).AddOffer(2, 0, exampleSpecBytes)

            const offers = (await broker.GetProvidersOffers(provider.address))


            expect(offers.length).to.equal(2)
        });
    })
})
