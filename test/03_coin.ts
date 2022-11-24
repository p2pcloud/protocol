import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'

import type { Contract } from "ethers";
import type { Token } from "../typechain-types";
import type { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

describe("BrokerV1_coin", function () {
    let broker: Contract;
    let token: Token;
    let miner: SignerWithAddress;
    let user: SignerWithAddress;

    before(async () => {
        [broker, miner, user, token] = await loadFixture(deployBrokerFixture);
    });

    describe("DepositCoin", function () {
        it("should incrase balance", async function () {
            const amt = ethers.utils.parseUnits('5', 6)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)
            await token.transfer(user.address, allowance)
            console.log(amt, allowance, await token.balanceOf(user.address), await user.getBalance())

            await broker.connect(token.connect(user).signer).DepositCoin(allowance)
        });
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
        it("should return locked and free balance", async function () {
            const balance = await broker.GetCoinBalance(user.address)
            console.log(balance)
        });
    })
    describe("SetCoinAddress", function () {
        it("should set coin address");
        it("should revert if not owner");
    })
});
