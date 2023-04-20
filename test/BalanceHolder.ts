import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployMarketplaceFixture } from './fixtures'

describe("BalanceHolder", function () {
    describe("depositCoin", function () {
        it("should incrase balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceFixture);

            await token.connect(admin).transfer(user.address, '123456')

            const [userBalance1] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(userBalance1.toString()).is.equal('0')

            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123')

            const [userBalance2] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(userBalance2.toString()).is.equal('123')
        });

        it("should revert if transfer fails", async function () {
            const { marketplace, user } = await loadFixture(deployMarketplaceFixture);
            await expect(marketplace.connect(user).depositCoin('123')).to.be.reverted
        });
    })
    describe("WithdrawCoin", function () {
        it("should not withdraw locked balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceFixture);

            //deposit 123 tokens
            await token.connect(admin).transfer(user.address, '123456')
            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123')

            //lock 23 tokens
            await marketplace.test__increaseLockedBalance(user.address, '23')

            //fail to withdraw 101 tokens
            await expect(marketplace.connect(user).withdrawCoin('101')).to.be.reverted

            //withdraw 100 tokens
            await marketplace.connect(user).withdrawCoin('100')
        });
        it("should revert if transfer fails", async function () {
            //TODO: modify coin contract to make this test pass
            const { marketplace, token, user } = await loadFixture(deployMarketplaceFixture);

            const [userBalance1] = await marketplace.connect(user).getCoinBalance(user.address)

            await token.transferShouldFail(true)
            await expect(marketplace.connect(user).withdrawCoin(1)).to.be.reverted

            const [userBalance2] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(userBalance2.toString()).is.equal(userBalance1.toString())

        });
    })
    describe("getCoinBalance", function () {
        it("should return locked and free balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceFixture);

            //deposit 123 tokens
            await token.connect(admin).transfer(user.address, '123456')
            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123')

            //lock 23 tokens
            await marketplace.test__increaseLockedBalance(user.address, '23')

            const [free, locked] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(free.toString()).is.equal('100')
            expect(locked.toString()).is.equal('23')
        });
    })
    describe("SetCoinAddress", function () {
        it("should set coin address", async function () {
            const { marketplace, token, admin } = await loadFixture(deployMarketplaceFixture);

            await marketplace.connect(admin).setCoin(token.address)
            expect(await marketplace.coin()).is.equal(token.address)
        });
        it("should revert if not owner", async function () {
            const { marketplace, token, user } = await loadFixture(deployMarketplaceFixture);

            await expect(marketplace.connect(user).setCoin(token.address)).to.be.reverted
        });
    })
});