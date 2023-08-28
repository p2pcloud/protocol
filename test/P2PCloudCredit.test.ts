import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployFiatMarketplaceV3Fixture } from './fixtures'
import { randomBytes12 } from "./lib";
import { ethers } from "ethers";


describe("P2PCloudCredit", () => {
    describe("setTrustedMinter", () => {
        it("should set trusted minter", async () => {
            const { admin, creditToken, anotherUser } = await loadFixture(deployFiatMarketplaceV3Fixture);

            await creditToken.connect(admin).setTrustedMinter(anotherUser.address)

            expect(await creditToken.trustedMinter()).to.equal(anotherUser.address)
        })

        it("should not set voucher signer if not owner", async () => {
            const { user, creditToken, anotherUser } = await loadFixture(deployFiatMarketplaceV3Fixture);

            await expect(creditToken.connect(user).setTrustedMinter(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")
        })
    })
    describe("setAllowedRecipient", () => {
        it("should set allowed recipient", async () => {
            const { admin, creditToken, anotherUser } = await loadFixture(deployFiatMarketplaceV3Fixture);

            await creditToken.connect(admin).setAllowedRecipient(anotherUser.address)

            expect(await creditToken.allowedRecipient()).to.equal(anotherUser.address)
        })

        it("should not set recipient if not owner", async () => {
            const { user, creditToken, anotherUser } = await loadFixture(deployFiatMarketplaceV3Fixture);

            await expect(creditToken.connect(user).setAllowedRecipient(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")
        })

        it("should not allow sending to recipient if not allowed", async () => {
            const { user, creditToken, anotherUser, admin, trustedMinter } = await loadFixture(deployFiatMarketplaceV3Fixture);

            await creditToken.connect(trustedMinter).idemopotentMint(user.address, 10000, randomBytes12())

            await expect(creditToken.connect(user).transfer(anotherUser.address, 100)).to.be.revertedWith("P2PCloudCredit: Recipient not allowed")

            await creditToken.connect(admin).setAllowedRecipient(anotherUser.address)

            await creditToken.connect(user).transfer(anotherUser.address, 100)
        })
    })
    describe("idemopotentMint", () => {
        it("should mint coins if authorized minter", async () => {
            const { creditToken, anotherUser, trustedMinter } = await loadFixture(deployFiatMarketplaceV3Fixture);

            expect(await creditToken.balanceOf(anotherUser.address)).to.equal(0)

            await creditToken.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, randomBytes12())

            expect(await creditToken.balanceOf(anotherUser.address)).to.equal(100)
        })
        it("should not mint coins if not authorized minter", async () => {
            const { user, creditToken, anotherUser } = await loadFixture(deployFiatMarketplaceV3Fixture);

            await expect(creditToken.connect(user).idemopotentMint(anotherUser.address, 100, randomBytes12())).to.be.revertedWith("P2PCloudCredit: Not authorized minter")
        })
        it("should not mint coins if already minted", async () => {
            const { creditToken, anotherUser, trustedMinter } = await loadFixture(deployFiatMarketplaceV3Fixture);

            const mintId = randomBytes12()
            await creditToken.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, mintId)

            await expect(creditToken.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, mintId)).to.be.revertedWith("P2PCloudCredit: Already minted")
        })
        it("should transfer value to recipient", async () => {
            const { creditToken, anotherUser, trustedMinter } = await loadFixture(deployFiatMarketplaceV3Fixture);

            //check creditToken balance 
            expect(await anotherUser.getBalance()).to.equal(ethers.utils.parseEther("10000"))
            const initialMinerBalance = await trustedMinter.getBalance()

            const mintId = randomBytes12()
            const tx = await creditToken.connect(trustedMinter).idemopotentMint(anotherUser.address, 7777, mintId, {
                value: ethers.utils.parseEther("123")
            })

            const gasUsed = await tx.wait().then((receipt) => receipt.gasUsed)
            const gasPrice = await tx.wait().then((receipt) => receipt.effectiveGasPrice)

            const expectedGasCost = gasPrice.mul(gasUsed)

            expect(await creditToken.balanceOf(anotherUser.address)).to.equal(7777)
            expect(await anotherUser.getBalance()).to.equal(ethers.utils.parseEther("10123"))

            const finalMinerBalance = await trustedMinter.getBalance()
            expect(initialMinerBalance.sub(finalMinerBalance).sub(expectedGasCost)).to.equal(ethers.utils.parseEther("123"))
        })
    })

    describe("burn (OpenZeppelin's)", () => {
        it("should burn tokens from the total supply", async () => {
            const { creditToken, trustedMinter, anotherUser } = await loadFixture(deployFiatMarketplaceV3Fixture);

            // Mint some tokens first
            await creditToken.connect(trustedMinter).idemopotentMint(anotherUser.address, 100, randomBytes12());

            const initialTotalSupply = await creditToken.totalSupply();
            await creditToken.connect(anotherUser).burn(30);
            const finalTotalSupply = await creditToken.totalSupply();

            expect(finalTotalSupply).to.equal(initialTotalSupply.sub(30));

            expect(await creditToken.balanceOf(anotherUser.address)).to.equal(70);
        });
    });

})