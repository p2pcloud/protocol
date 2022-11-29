import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'

describe("BrokerV1_getMinerUrl", function () {
    it("should set miner url", async function () {
        const { broker, miner } = await loadFixture(deployBrokerFixture);
        const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
        await broker.SetMinerUrl(urlBytes)
        const url = await broker.GetMinerUrl(miner.address);
        expect(url).to.equal(urlBytes);
    });
});
