import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { DEFAULT_USER_BALANCE, deployFiatMarketplaceFixture, deployMarketplaceFixture } from './fixtures'
import { UnsignedVoucher, setUserCoinBalance, signVoucher } from "./lib";
import { ethers, providers } from "ethers";


describe("FiatMarketplace", () => {
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
        it("should add voucher to balance", async () => {
            const { marketplace, provider, admin, user, voucherSigner } = await loadFixture(deployFiatMarketplaceFixture);

            await marketplace.connect(admin).setVoucherSigner(voucherSigner.address)

            const voucher: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature = await signVoucher(voucherSigner, voucher, marketplace.address)

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE)
            await marketplace.connect(user).claimVoucher(voucher, signature)
            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE + 100)
        })
        it("should not allow voucher reuse", async () => {
            const { marketplace, admin, user, voucherSigner } = await loadFixture(deployFiatMarketplaceFixture);

            await marketplace.connect(admin).setVoucherSigner(voucherSigner.address)

            //voucher 1 
            const voucher1: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature1 = await signVoucher(voucherSigner, voucher1, marketplace.address)

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE)
            await marketplace.connect(user).claimVoucher(voucher1, signature1)
            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE + 100)

            //voucher 2 
            const voucher2: UnsignedVoucher = {
                amount: 123456789,
                paymentId: ethers.utils.formatBytes32String("two"),
            }
            const signature2 = await signVoucher(voucherSigner, voucher2, marketplace.address)

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE + 100)
            await marketplace.connect(user).claimVoucher(voucher2, signature2)
            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE + 100 + 123456789)

            //voucher 2 reuse
            await expect(marketplace.connect(user).claimVoucher(voucher2, signature2)).to.be.revertedWith("Voucher already used")
        })
        it("should evert if signature is invalid", async () => {
            const { marketplace, admin, user, voucherSigner } = await loadFixture(deployFiatMarketplaceFixture);

            await marketplace.connect(admin).setVoucherSigner(voucherSigner.address)

            //wrong signer
            const voucher1: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("one"),
            }
            const signature1 = await signVoucher(admin, voucher1, marketplace.address)

            await expect(marketplace.connect(user).claimVoucher(voucher1, signature1)).to.be.revertedWith("Invalid signature")

            //wrong amount
            const voucher2: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("two"),
            }
            const signature2 = await signVoucher(voucherSigner, voucher2, marketplace.address)
            voucher2.amount = 200

            await expect(marketplace.connect(user).claimVoucher(voucher2, signature2)).to.be.revertedWith("Invalid signature")

            //wrong paymentId
            const voucher3: UnsignedVoucher = {
                amount: 100,
                paymentId: ethers.utils.formatBytes32String("three"),
            }
            const signature3 = await signVoucher(voucherSigner, voucher3, marketplace.address)
            voucher3.paymentId = ethers.utils.formatBytes32String("four")

            await expect(marketplace.connect(user).claimVoucher(voucher3, signature3)).to.be.revertedWith("Invalid signature")
        })
    })
    describe("burnCoin", () => {
        it("should burn coin", async () => {
            const { marketplace, admin, user, voucherSigner } = await loadFixture(deployFiatMarketplaceFixture);

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE)

            await marketplace.connect(admin).burnCoin(1234, user.address)

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE - 1234)

        })
        it("should not burn coin if not owner", async () => {
            const { marketplace, admin, user, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE)

            await expect(marketplace.connect(anotherUser).burnCoin(1234, user.address)).to.be.revertedWith("Ownable: caller is not the owner")

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE)
        })
    })
    describe("withdrawCoin", () => {
        it("should withdraw coin only for provider", async () => {
            const { marketplace, admin, user, token, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            await token.connect(admin).transfer(marketplace.address, 200000000)

            //top up balance
            const fee = (await marketplace.PROVIDER_REGISTRATION_FEE()).toNumber()
            const voucher = { amount: fee, paymentId: ethers.utils.formatBytes32String("three") }
            const signature = await signVoucher(voucherSigner, voucher, marketplace.address)
            await marketplace.connect(user).claimVoucher(voucher, signature)

            //try to withdraw
            expect(await marketplace.getTotalBalance(user.address)).to.equal(fee + DEFAULT_USER_BALANCE)
            await expect(marketplace.connect(user).withdrawCoin(1234)).to.be.revertedWith("Only provider can withdraw")

            //register as provider and try again
            await marketplace.connect(admin).registerFiatProvider(user.address)
            await marketplace.connect(user).withdrawCoin(1234)

            expect(await marketplace.getTotalBalance(user.address)).to.equal(DEFAULT_USER_BALANCE - 1234)
        })
    })
    describe("depositCoin", () => {
        it("should not work at all", async () => {
            const { marketplace, admin, user, provider, token, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            for (let wallet of [admin, user, provider, anotherUser]) {
                await token.connect(admin).transfer(wallet.address, 12345)
                await token.connect(wallet).approve(marketplace.address, 12345)

                await expect(marketplace.connect(wallet).depositCoin(1234)).to.be.revertedWith("Not supported")
            }
        })
    })
    describe("registerProvider", () => {
        it("should be disabled", async () => {
            const { marketplace, admin, user, provider, token, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);


            for (let wallet of [admin, user, anotherUser]) {
                const fee = (await marketplace.PROVIDER_REGISTRATION_FEE()).toNumber()
                const voucher = { amount: fee, paymentId: ethers.utils.formatBytes32String(wallet.address.slice(-10)) }
                const signature = await signVoucher(voucherSigner, voucher, marketplace.address)
                await marketplace.connect(wallet).claimVoucher(voucher, signature)

                await expect(marketplace.connect(wallet).registerProvider()).to.be.revertedWith("Not supported")
            }
        })
    })

    describe("registerFiatProvider", () => {
        it("should register if owner", async () => {
            const { marketplace, admin, user, token, provider, voucherSigner, anotherUser } = await loadFixture(deployFiatMarketplaceFixture);

            //top up balance
            const fee = (await marketplace.PROVIDER_REGISTRATION_FEE()).toNumber()
            const voucher = { amount: fee, paymentId: ethers.utils.formatBytes32String("three") }
            const signature = await signVoucher(voucherSigner, voucher, marketplace.address)
            await marketplace.connect(anotherUser).claimVoucher(voucher, signature)

            expect(await marketplace.isProviderRegistered(anotherUser.address)).to.equal(false)

            await expect(marketplace.connect(anotherUser).registerFiatProvider(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")
            await expect(marketplace.connect(provider).registerFiatProvider(anotherUser.address)).to.be.revertedWith("Ownable: caller is not the owner")

            expect(await marketplace.isProviderRegistered(anotherUser.address)).to.equal(false)

            await marketplace.connect(admin).registerFiatProvider(anotherUser.address)

            expect(await marketplace.isProviderRegistered(anotherUser.address)).to.equal(true)
        })
    })
})