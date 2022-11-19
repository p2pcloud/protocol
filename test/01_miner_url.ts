import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("BrokerV1_getMinerUrl", function () {
    async function deployBrokerFixture() {
        const [owner, otherAccount] = await ethers.getSigners();

        const Broker = await ethers.getContractFactory("BrokerV1");
        const broker = await upgrades.deployProxy(Broker);

        return { broker, owner, otherAccount };
    }

    it("should set miner url", async function () {
        const { broker, owner } = await loadFixture(deployBrokerFixture);
        const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
        await broker.SetMinerUrl(urlBytes)
        const url = await broker.GetMinerUrl(owner.address);
        expect(url).to.equal(urlBytes);
    });
});
