import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceV3Fixture } from './fixtures'
import { NZ_HEX, US_HEX, signKYC } from "./lib";

describe("VerifiableKYC", function () {
    it("should return KYC status", async function () {
        const { marketplace, user, providersSigner, anotherUser, kycSigner, provider } = await loadFixture(deployMarketplaceV3Fixture);

        expect(await marketplace.KYCStatus(provider.address)).to.equal(US_HEX)
        expect(await marketplace.KYCStatus(anotherUser.address)).to.equal("0x0000")
    })
    it("should not accept if receiver address or country is rigged", async function () {
        const { marketplace, user, providersSigner, anotherUser, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

        //wrong signer
        await expect(marketplace.connect(anotherUser).submitKYC(
            anotherUser.address, US_HEX,
            await signKYC(anotherUser.address, US_HEX, user)
        )).to.be.revertedWithCustomError(marketplace, "InvalidKYCSigner")

        //rig country
        await expect(marketplace.connect(anotherUser).submitKYC(
            anotherUser.address, US_HEX,
            await signKYC(anotherUser.address, NZ_HEX, kycSigner)
        )).to.be.revertedWithCustomError(marketplace, "InvalidKYCSigner")

        //rig address
        await expect(marketplace.connect(anotherUser).submitKYC(
            anotherUser.address, US_HEX,
            await signKYC(user.address, US_HEX, kycSigner)
        )).to.be.revertedWithCustomError(marketplace, "InvalidKYCSigner")

        await marketplace.connect(anotherUser).submitKYC(
            anotherUser.address, US_HEX,
            await signKYC(anotherUser.address, US_HEX, kycSigner)
        )
    })
}) 