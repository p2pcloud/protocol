import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployFiatMarketplaceFixture } from './fixtures'
import { UnsignedVoucher, signVoucher } from "./lib";
import { BigNumber, ethers } from "ethers";


describe.only("P2PCloudCredit", () => {
    describe("setTrustedMinter", () => {
        it("should set trusted minter", async () => {
            const { admin, token, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setTrustedMinter(anotherUser.address)

            expect(await token.trustedMinter()).to.equal(anotherUser.address)
        })

        it("should not set voucher signer if not owner", async () => {
            const { user, token, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await expect(token.connect(user).setTrustedMinter(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")
        })
    })
    describe("setAllowedRecipient", () => {
        it("should set allowed recipient", async () => {
            const { admin, token, anotherUser, marketplace } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setAllowedRecipient(anotherUser.address)

            expect(await token.allowedRecipient()).to.equal(anotherUser.address)
        })

        it("should not set recipient if not owner", async () => {
            const { user, token, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await expect(token.connect(user).setAllowedRecipient(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")
        })

        it("should not allow sending to recipient if not allowed", async () => {
            const { user, token, anotherUser, admin } = await loadFixture(deployFiatMarketplaceFixture);

            await expect(token.connect(user).transfer(anotherUser.address, 100)).to.be.revertedWith("ERC20: Recipient not allowed")

            await token.connect(admin).setAllowedRecipient(anotherUser.address)

            await token.connect(admin).transfer(anotherUser.address, 100)
        })
    })
})
