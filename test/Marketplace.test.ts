import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceV3Fixture } from './fixtures'
import { ethers, upgrades } from "hardhat";


describe("Markeplace", function () {
    it("should reject initialize call")
    it("should accept owner change only from owner", async function () {
        const { marketplace, user, anotherUser, admin } = await loadFixture(deployMarketplaceV3Fixture);

        expect(await marketplace.owner()).to.equal(admin.address);

        await expect(marketplace.connect(user).transferOwnership(anotherUser.address))
            .to.be.revertedWith("Ownable: caller is not the owner")

        await marketplace.connect(admin).transferOwnership(anotherUser.address)

        expect(await marketplace.owner()).to.equal(anotherUser.address);
    })
    it("should not fail upgrade", async function () {
        const { marketplace, user, anotherUser, admin } = await loadFixture(deployMarketplaceV3Fixture);
        const owner = await marketplace.owner();

        const TestableMarketplaceV3Upgraded = await ethers.getContractFactory("TestableMarketplaceV3Upgraded");
        await upgrades.upgradeProxy(marketplace.address, TestableMarketplaceV3Upgraded);
    });
})