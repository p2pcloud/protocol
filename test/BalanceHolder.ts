import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployMarketplaceFixture } from './fixtures'

describe("BalanceHolder", function () {
    describe("depositCoin", function () {
        it("should incrase balance", async function () {
            const { marketplace, token, user, admin } = await loadFixture(deployMarketplaceFixture);

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getCoinBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

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

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getCoinBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123456 tokens
            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '1')
            const [free1, locked1] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(locked1.toString()).is.equal('10080')
            expect(free1.toString()).is.equal(String(123456 - 10080))

            //fail to withdraw free+1
            await expect(marketplace.connect(user).withdrawCoin(123456 - 10080 + 1)).to.be.reverted

            //withdraw excatly free
            await marketplace.connect(user).withdrawCoin(123456 - 10080)
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

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getCoinBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123 tokens
            await token.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '1')

            const [free, locked] = await marketplace.connect(user).getCoinBalance(user.address)
            expect(free.toString()).is.equal(String(123456 - 60 * 24 * 7))
            expect(locked.toString()).is.equal(String(60 * 24 * 7))
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