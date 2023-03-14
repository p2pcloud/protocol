import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'

describe("Broker_addTrustedKey", function () {
    it("should set trusted key", async function () {
        const { broker, user } = await loadFixture(deployBrokerFixture);
        const trustedAddress = "0x11111AC4a2f2f49191D9a5E6D503dA27724b1Bd4"

        await broker.connect(user).setTrustedAddress(trustedAddress);
        expect(await broker.getTrustedAddress(user.address)).to.equal(trustedAddress);
    });
});