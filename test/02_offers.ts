import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployBrokerFixture } from './fixtures'

import type { Contract } from "ethers";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

type Offer = {
    Index: number,
    Miner: string,
    PPS: number,
    Availablility: number,
    VmTypeId: number
}

describe("BrokerV1_offers", function () {
    let fixture: {
        broker: Contract;
        owner: SignerWithAddress;
        otherAccount: SignerWithAddress;
    }

    before(async () => {
        fixture = await loadFixture(deployBrokerFixture);
    });

    describe("AddOffers", function () {
        it("should add offers", async function () {
            const { broker, owner, otherAccount } = fixture;
            const offers = [
                [1, 1, 1],
                [2, 2, 2],
                [3, 3, 3],
                [4, 4, 4],
                [5, 5, 5],
            ]
            let count = 0
            await Promise.all(offers.map(async (offer) => {
                count++
                return await broker.AddOffer(...offer);
            }))
            expect(count).is.equal(5)
        });
    })
    describe("UpdateOffer", function () {
        it("should update offer", async function () {
            const { broker, owner, otherAccount } = fixture;
            await broker.UpdateOffer(0, 0, 1)
        });
        it("should revert if offer belongs to another account", async function () {
            const { broker, owner, otherAccount } = fixture;
            await expect(broker.connect(otherAccount).UpdateOffer(0, 0, 1)).to.be.reverted
        });
    })
    describe("RemoveOffer", function () {
        it("should remove offer", async function () {
            const { broker, owner, otherAccount } = fixture;
            await broker.RemoveOffer(1)
            const offersRaw = await broker.GetAvailableOffers();
            const offers = offersRaw.map(offerFromRaw)
            expect(offers.length).to.equal(3)
        });
        it("should revert if offer belongs to another account", async function () {
            const { broker, owner, otherAccount } = fixture;
            await expect(broker.connect(otherAccount).RemoveOffer(1)).to.be.reverted
        });
    })
    describe("GetAvailableOffers", function () {
        it("should return array of offers", async function () {
            const { broker, owner, otherAccount } = fixture;
            const offersRaw = await broker.GetAvailableOffers();
            const offers = offersRaw.map(offerFromRaw)
            expect(offers.length).to.equal(3)
        });
        it("should return only offers with available machines", async function () {
            const { broker, owner, otherAccount } = fixture;
            const offersRaw = await broker.GetAvailableOffers();
            const offers = offersRaw.map(offerFromRaw)
            const availability = offers.every(({ Availablility }: { Availablility: number }) => Availablility)
            expect(availability).true
        });
        it("should return only offers with type 5", async function () {
            const { broker, owner, otherAccount } = fixture;
            const offersRaw = await broker.GetAvailableOffersByType(5);
            const offers: Offer[] = offersRaw.map(offerFromRaw)
            const typed = offers.every(({ VmTypeId }) => VmTypeId === 5)
            expect(typed).true
        });
    })
    describe("GetMinersOffers", function () {
        it("should return array of offers", async function () {
            const { broker, owner, otherAccount } = fixture;
            const offersRaw = await broker.GetMinersOffers(owner.address);
            const offers = offersRaw.map(offerFromRaw)
            expect(offers.length).to.equal(4)
        });
        it("should return offers including ones with availability zero", async function () {
            const { broker, owner, otherAccount } = fixture;
            const offersRaw = await broker.GetMinersOffers(owner.address);
            const offers = offersRaw.map(offerFromRaw)
            const availability = offers.every(({ Availablility }: { Availablility: number }) => Availablility)
            expect(availability).false
        });
    })
})

function offerFromRaw(offerRaw: any[]) {
    const [Index, Miner, PPS, Availablility, VmTypeId] = offerRaw

    return {
        VmTypeId: VmTypeId.toNumber(),
        PPS: PPS.toNumber(),
        Availablility: Availablility.toNumber(),
        Miner: Miner,
        Index: Index.toNumber(),
    }
}
