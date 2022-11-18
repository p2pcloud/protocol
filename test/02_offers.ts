import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("BrokerV1_offers", function () {
    describe("UpdateOffer", function () {
        it("should update offer");
        it("should revert if offer belongs to another account");
    })
    describe("RemoveOffer", function () {
        it("should remove offer");
        it("should revert if offer belongs to another account");
    })
    describe("GetAvailableOffers", function () {
        it("should return array of offers");
        it("should return only offers with available machines");
    })
    describe("GetMinersOffers", function () {
        it("should return array of offers");
        it("should return offers including ones with availability zero");
    })
});
