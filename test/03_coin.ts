import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("BrokerV1_coin", function () {
    describe("DepositCoin", function () {
        it("should incrase balance");
        it("should revert if transfer fails");
    })
    describe("GetLockedCoinBalance", function () {
        it("should increase with vm booking");
        it("should decrease with vm termination");
    })
    describe("WithdrawCoin", function () {
        it("should withdraw only free balance");
        it("should revert if transfer fails");
        it("should revert if not enough balance");
    })
    describe("GetCoinBalance", function () {
        it("should return locked and free balance");
    })
    describe("SetCoinAddress", function () {
        it("should set coin address");
        it("should revert if not owner");
    })
});
