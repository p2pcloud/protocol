import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("BrokerV1_stablecoin", function () {
    describe("DepositStablecoin", function () {
        it("should incrase balance");
        it("should revert if transfer fails");
    })
    describe("GetLockedStablecoinBalance", function () {
        it("should increase with vm booking");
        it("should decrease with vm termination");
    })
    describe("WithdrawStablecoin", function () {
        it("should withdraw only free balance");
        it("should revert if transfer fails");
        it("should revert if not enough balance");
    })
    describe("GetStablecoinBalance", function () {
        it("should return locked and free balance");
    })
    describe("SetStablecoinAddress", function () {
        it("should set stablecoin address");
        it("should revert if not owner");
    })
});
