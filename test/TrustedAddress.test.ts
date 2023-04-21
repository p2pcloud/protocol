import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceFixture } from './fixtures'

describe("TrustedAddress", function () {
    it("should set trusted address", async function () {
        const { marketplace, user, anotherUser } = await loadFixture(deployMarketplaceFixture);
        const trustedAddress = "0x2352D20fC81225c8ECD8f6FaA1B37F24FEd450c9"

        await marketplace.connect(user).setTrustedAddress(trustedAddress);
        expect(await marketplace.connect(anotherUser).getTrustedAddress(user.address)).to.equal(trustedAddress);
    });
});