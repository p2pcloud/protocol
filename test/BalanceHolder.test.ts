import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployMarketplaceV3Fixture } from './fixtures'
import { NZ_HEX, signKYC } from "./lib";

describe("BalanceHolder", function () {
    describe("depositCoin", function () {
        it("should incrase balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            const [userBalance1] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance1.toString()).is.equal('0')

            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123')

            const [userBalance2] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance2.toString()).is.equal('123')
        });

        it("should revert if transfer fails", async function () {
            const { marketplace, user } = await loadFixture(deployMarketplaceV3Fixture);
            await expect(marketplace.connect(user).depositCoin('123')).to.be.reverted
        });

        it("should require user's KYC", async function () {
            const { marketplace, admin, token, anotherUser, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await token.connect(admin).transfer(anotherUser.address, '123456')
            await token.connect(anotherUser).approve(marketplace.address, '123456')

            await expect(marketplace.connect(anotherUser).depositCoin('123456')).to.be.revertedWith('No KYC or country is not allowed')

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, NZ_HEX,
                await signKYC(anotherUser.address, NZ_HEX, kycSigner)
            )

            await expect(marketplace.connect(anotherUser).depositCoin('123456')).to.be.revertedWith('No KYC or country is not allowed')

            await marketplace.connect(admin).allowUserCountry(NZ_HEX)

            await marketplace.connect(anotherUser).depositCoin('123456')
        });
    })
    describe("WithdrawCoin", function () {
        it("should not withdraw locked balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123456 tokens
            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '1')
            const [free1, locked1] = await marketplace.connect(user).getBalance(user.address)
            expect(locked1.toString()).is.equal('10080')
            expect(free1.toString()).is.equal(String(123456 - 10080))

            //fail to withdraw free+1
            await expect(marketplace.connect(user).withdrawCoin(123456 - 10080 + 1)).to.be.reverted

            //withdraw excatly free
            await marketplace.connect(user).withdrawCoin(123456 - 10080)
        });
        it("should revert if transfer fails", async function () {
            //TODO: modify coin contract to make this test pass
            const { marketplace, token, user } = await loadFixture(deployMarketplaceV3Fixture);

            const [userBalance1] = await marketplace.connect(user).getBalance(user.address)

            await token.transferShouldFail(true)
            await expect(marketplace.connect(user).withdrawCoin(1)).to.be.reverted

            const [userBalance2] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance2.toString()).is.equal(userBalance1.toString())

        });
    })
    describe("getBalance", function () {
        it("should return locked and free balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const defaultBalance = await marketplace.connect(user).getFreeBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123 tokens
            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '1')

            const [free, locked] = await marketplace.connect(user).getBalance(user.address)
            expect(free.toString()).is.equal(String(123456 - 60 * 24 * 7))
            expect(locked.toString()).is.equal(String(60 * 24 * 7))
        });
    })
});