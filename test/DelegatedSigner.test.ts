import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceFixture } from './fixtures'

describe("DelegatedSigner", function () {
    it("should set trusted address", async function () {
        const { marketplace, user, anotherUser } = await loadFixture(deployMarketplaceFixture);
        const signerAddress = "0x2352D20fC81225c8ECD8f6FaA1B37F24FEd450c9"

        await marketplace.connect(user).setSigner(signerAddress);
        expect(await marketplace.connect(anotherUser).getSigner(user.address)).to.equal(signerAddress);
        expect(await marketplace.connect(anotherUser).resolveSigner(signerAddress)).to.equal(user.address);
    });

    it("should not allow reusing signerAddress", async function () {
        const { marketplace, user, anotherUser } = await loadFixture(deployMarketplaceFixture);
        const trustedAddress = "0x2352D20fC81225c8ECD8f6FaA1B37F24FEd450c9"

        await marketplace.connect(user).setSigner(trustedAddress);
        await expect(marketplace.connect(anotherUser).setSigner(trustedAddress)).to.be.revertedWith("Signer already in use");
    });

    it("should allow changing signing key and back", async function () {
        const { marketplace, user, anotherUser } = await loadFixture(deployMarketplaceFixture);
        const signerAddress1 = "0x1111111111111111111111111111111111111111"
        const signerAddress2 = "0x2222222222222222222222222222222222222222"

        await marketplace.connect(user).setSigner(signerAddress1);
        expect(await marketplace.connect(anotherUser).getSigner(user.address)).to.equal(signerAddress1);

        await marketplace.connect(user).setSigner(signerAddress2);
        expect(await marketplace.connect(anotherUser).getSigner(user.address)).to.equal(signerAddress2);

        await marketplace.connect(user).setSigner(signerAddress1);
        expect(await marketplace.connect(anotherUser).getSigner(user.address)).to.equal(signerAddress1);
    });
});