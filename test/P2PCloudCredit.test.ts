import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployFiatMarketplaceFixture } from './fixtures'
import { UnsignedVoucher, signVoucher } from "./lib";
import { BigNumber, ethers } from "ethers";


describe("P2PCloudCredit", () => {
    describe("setVoucherSigner", () => {
        it("should set voucher signer", async () => {
            const { admin, token } = await loadFixture(deployFiatMarketplaceFixture);
            const voucherSigner = "0x21539334f45Ac41Bd10789942b744a18a4775d6d"

            await token.connect(admin).setVoucherSigner(voucherSigner)

            expect(await token.voucherSigner()).to.equal(voucherSigner)
        })
        it("should not set voucher signer if not owner", async () => {
            const { user, token } = await loadFixture(deployFiatMarketplaceFixture);
            const voucherSigner = "0x21539334f45Ac41Bd10789942b744a18a4775d6d"

            await expect(token.connect(user).setVoucherSigner(voucherSigner)).to.be.revertedWith("Ownable: caller is not the owner")
        })
    })
    describe("claimVoucher", () => {
        it("should add voucher to balance", async () => {
            const { admin, user, anotherUser, voucherSigner, token } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setVoucherSigner(voucherSigner.address)

            const voucher: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature = await signVoucher(voucherSigner, voucher, token.address)

            expect(await token.balanceOf(user.address)).to.equal(0)
            await token.connect(anotherUser).claimVoucher(voucher, signature, user.address)
            expect(await token.balanceOf(user.address)).to.equal(100)
        })
        it("should not allow voucher reuse", async () => {
            const { token, admin, user, anotherUser, voucherSigner } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setVoucherSigner(voucherSigner.address)

            //voucher 1 
            const voucher1: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature1 = await signVoucher(voucherSigner, voucher1, token.address)

            expect(await token.balanceOf(user.address)).to.equal(0)
            await token.connect(anotherUser).claimVoucher(voucher1, signature1, user.address)
            expect(await token.balanceOf(user.address)).to.equal(100)

            //voucher 2 
            const voucher2: UnsignedVoucher = {
                amount: 123456789,
                paymentId: ethers.utils.formatBytes32String("two"),
            }
            const signature2 = await signVoucher(voucherSigner, voucher2, token.address)

            expect(await token.balanceOf(user.address)).to.equal(100)
            await token.connect(anotherUser).claimVoucher(voucher2, signature2, user.address)
            expect(await token.balanceOf(user.address)).to.equal(100 + 123456789)

            //voucher 2 reuse
            await expect(token.connect(anotherUser).claimVoucher(voucher2, signature2, user.address)).to.be.revertedWith("Voucher already used")
        })
        it("should evert if signature is invalid", async () => {
            const { token, admin, user, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setVoucherSigner(voucherSigner.address)

            //wrong signer
            const voucher1: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature1 = await signVoucher(admin, voucher1, token.address)

            await expect(token.connect(anotherUser).claimVoucher(voucher1, signature1, user.address)).to.be.revertedWith("Invalid signature")

            //wrong amount
            const voucher2: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("two"),
            }
            const signature2 = await signVoucher(voucherSigner, voucher2, token.address)
            voucher2.amount = 200

            await expect(token.connect(anotherUser).claimVoucher(voucher2, signature2, user.address)).to.be.revertedWith("Invalid signature")

            //wrong paymentId
            const voucher3: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("three"),
            }
            const signature3 = await signVoucher(voucherSigner, voucher3, token.address)
            voucher3.paymentId = ethers.utils.formatBytes32String("four")

            await expect(token.connect(anotherUser).claimVoucher(voucher3, signature3, user.address)).to.be.revertedWith("Invalid signature")
        })

        it("should emit VoucherClaimed event", async () => {
            const { token, admin, user, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setVoucherSigner(voucherSigner.address)

            //voucher 1 
            const voucher1: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature1 = await signVoucher(voucherSigner, voucher1, token.address)

            await expect(await token.connect(anotherUser).claimVoucher(voucher1, signature1, user.address)).to.emit(token, "VoucherClaimed").withArgs(user.address, voucher1.amount, voucher1.paymentId)
        })


        it("should proxy value if provided", async () => {
            const { token, admin, user, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).setVoucherSigner(voucherSigner.address)

            //voucher 1 
            const voucher1: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature1 = await signVoucher(voucherSigner, voucher1, token.address)

            const balanceBefore = await user.getBalance()
            const balanceAdded = BigNumber.from("300000")

            await token.connect(anotherUser).claimVoucher(voucher1, signature1, user.address, {
                value: balanceAdded,
            })

            const balanceAfter = await user.getBalance()
            expect(balanceAfter.sub(balanceBefore)).to.equal(balanceAdded)
        })
    })
})
