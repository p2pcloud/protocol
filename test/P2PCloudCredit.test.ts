import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployFiatMarketplaceFixture, randomBytes12 } from './fixtures'
import { UnsignedVoucher, signVoucher } from "./lib";
import { BigNumber, ethers } from "ethers";


describe("P2PCloudCredit", () => {
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
            const { admin, token, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setAllowedRecipient(anotherUser.address)

            expect(await token.allowedRecipient()).to.equal(anotherUser.address)
        })

        it("should not set recipient if not owner", async () => {
            const { user, token, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await expect(token.connect(user).setAllowedRecipient(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")
        })

        it("should not allow sending to recipient if not allowed", async () => {
            const { user, token, anotherUser, admin } = await loadFixture(deployFiatMarketplaceFixture);

            await expect(token.connect(user).transfer(anotherUser.address, 100)).to.be.revertedWith("P2PCloudCredit: Recipient not allowed")

            await token.connect(admin).setAllowedRecipient(anotherUser.address)

            await token.connect(admin).transfer(anotherUser.address, 100)
        })
    })
    describe("idemopotentMint", () => {
        it("should mint coins if authorized minter", async () => {
            const { token, anotherUser, trustedMinter } = await loadFixture(deployFiatMarketplaceFixture);

            expect(await token.balanceOf(anotherUser.address)).to.equal(0)

            await token.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, randomBytes12())

            expect(await token.balanceOf(anotherUser.address)).to.equal(100)
        })
        it("should not mint coins if not authorized minter", async () => {
            const { user, token, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await expect(token.connect(user).idemopotentMint(anotherUser.address, 100, randomBytes12())).to.be.revertedWith("P2PCloudCredit: Not authorized minter")
        })
        it("should not mint coins if already minted", async () => {
            const { token, anotherUser, trustedMinter } = await loadFixture(deployFiatMarketplaceFixture);

            const mintId = randomBytes12()
            await token.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, mintId)

            await expect(token.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, mintId)).to.be.revertedWith("P2PCloudCredit: Already minted")
        })
        it("should transfer value to recipient", async () => {
            const { token, anotherUser, trustedMinter } = await loadFixture(deployFiatMarketplaceFixture);

            //check token balance 
            expect(await anotherUser.getBalance()).to.equal(ethers.utils.parseEther("10000"))
            const initialMinerBalance = await trustedMinter.getBalance()

            const mintId = randomBytes12()
            const tx = await token.connect(trustedMinter).idemopotentMint(anotherUser.address, 7777, mintId, {
                value: ethers.utils.parseEther("123")
            })

            const gasUsed = await tx.wait().then((receipt) => receipt.gasUsed)
            const gasPrice = await tx.wait().then((receipt) => receipt.effectiveGasPrice)

            const expectedGasCost = gasPrice.mul(gasUsed)

            expect(await token.balanceOf(anotherUser.address)).to.equal(7777)
            expect(await anotherUser.getBalance()).to.equal(ethers.utils.parseEther("10123"))

            const finalMinerBalance = await trustedMinter.getBalance()
            expect(initialMinerBalance.sub(finalMinerBalance).sub(expectedGasCost)).to.equal(ethers.utils.parseEther("123"))
        })
    })

    describe("burn (OpenZeppelin's)", () => {
        it("should burn tokens from the total supply", async () => {
            const { token, trustedMinter, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            // Mint some tokens first
            await token.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, randomBytes12());

            const initialTotalSupply = await token.totalSupply();
            await token.connect(anotherUser).burn(30);
            const finalTotalSupply = await token.totalSupply();

            expect(finalTotalSupply).to.equal(initialTotalSupply.sub(30));

            expect(await token.balanceOf(anotherUser.address)).to.equal(70);
        });
    });

})
