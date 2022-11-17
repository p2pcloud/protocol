import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("Broker", function () {
    // We define a fixture to reuse the same setup in every test.
    // We use loadFixture to run this setup once, snapshot that state,
    // and reset Hardhat Network to that snapshot in every test.
    async function deployBrokerFixture() {
        const [owner, otherAccount] = await ethers.getSigners();

        const Broker = await ethers.getContractFactory("Broker");
        const broker = await Broker.deploy();

        return { broker, owner, otherAccount };
    }

    describe("Miner", function () {
        describe("URL", function () {
            it("Should set miner url", async function () {
                const { broker } = await loadFixture(deployBrokerFixture);
                const urlBytes = ethers.utils.formatBytes32String("example.com");
                await expect(broker.setMunerUrl(urlBytes)).not.to.be.reverted;
            });

            it("Should return miner url", async function () {
                const { broker, owner } = await loadFixture(deployBrokerFixture);
                const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
                await broker.setMunerUrl(urlBytes)
                const url = await broker.getMinerUrl(owner.address);
                expect(url).to.equal(urlBytes);
            });
        });

    });
});
