import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("BrokerV1_getMinerUrl", function () {

    it("should keep data while upgrading the contract", async function () {
        const [owner, otherAccount] = await ethers.getSigners();

        //deploy broker
        const BrokerV1 = await ethers.getContractFactory("BrokerV1");
        const broker = await upgrades.deployProxy(BrokerV1);

        //set url
        const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
        await broker.SetMinerUrl(urlBytes)
        const url = await broker.GetMinerUrl(owner.address);
        expect(ethers.utils.parseBytes32String(url)).to.equal("woop.woop/woop");

        //upgrade broker
        const BrokerTestChange = await ethers.getContractFactory("BrokerTestChange");
        const upgraded = await upgrades.upgradeProxy(broker.address, BrokerTestChange);

        const urlAfterUpgrade = await upgraded.RenamedGetMinerUrl(owner.address);
        expect(ethers.utils.parseBytes32String(urlAfterUpgrade)).to.equal("woop.woop/woop");
    });
});
