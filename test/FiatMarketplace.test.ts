import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployFiatMarketplaceFixture, deployMarketplaceFixture } from './fixtures'


describe.only("FiatMarketplace", () => {
    describe("setVoucherSigner", () => {
        it("should set voucher signer", async () => {
            const { marketplace, provider, admin } = await loadFixture(deployFiatMarketplaceFixture);
            const voucherSigner = "0x21539334f45Ac41Bd10789942b744a18a4775d6d"

            await marketplace.connect(admin).setVoucherSigner(voucherSigner)

            expect(await marketplace.voucherSigner()).to.equal(voucherSigner)
        })
        it("should not set voucher signer if not owner", async () => {
            const { marketplace, provider, admin, user } = await loadFixture(deployFiatMarketplaceFixture);
            const voucherSigner = "0x21539334f45Ac41Bd10789942b744a18a4775d6d"

            await expect(marketplace.connect(user).setVoucherSigner(voucherSigner)).to.be.revertedWith("Ownable: caller is not the owner")
        })
    })
    describe("claimVoucher", () => {
        it("should add voucher to balance")
        it("should not allow vouche reuse")
        it("should refuse if signature is invalid")
    })
    describe("burnCoin", () => {
        it("should burn coin")
        it("should not burn coin if not owner")
    })
    describe("withdrawCoin", () => {
        it("should withdraw coin if not a provider")
    })
    describe("registerProvider", () => {
        it("should not allow anybody to register as a provider")
    })
    describe("registerProvider", () => {
        it("should not work without params")
        it("should register if owner")
    })
})