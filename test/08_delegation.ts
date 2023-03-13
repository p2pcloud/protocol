import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'

describe("Broker_addTrustedKey", function () {
    it("should set trusted key", async function () {
        const { broker, user } = await loadFixture(deployBrokerFixture);
        const publicKey = ethers.utils.formatBytes32String("encryptionKey");

        await broker.connect(user).setTrustedKey(publicKey);
        expect(await broker.getTrustedKey(user.address)).to.equal(publicKey);
    });
});