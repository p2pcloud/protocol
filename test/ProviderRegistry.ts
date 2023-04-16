import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'

describe("ProviderRegistry", function () {
    describe("getProviderUrl", function () {
        it("should set provider url", async function () {
            const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            const fee = await broker.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(broker.address, fee)
            await broker.connect(anotherUser).depositCoin(fee)
            await broker.connect(anotherUser).registerProvider()

            const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
            await broker.connect(anotherUser).setProviderUrl(urlBytes)
            const url = await broker.connect(anotherUser).getProviderUrl(anotherUser.address);
            expect(url).to.equal(urlBytes);
        });

        it("should not set provider url for non-registered provider", async function () {
            const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);


            const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
            expect(broker.connect(anotherUser).setProviderUrl(urlBytes)).to.be.revertedWith("Provider not registered")

            const fee = await broker.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(broker.address, fee)
            await broker.connect(anotherUser).depositCoin(fee)
            await broker.connect(anotherUser).registerProvider()

            await broker.connect(anotherUser).setProviderUrl(urlBytes)
        });
    })


    describe("isProviderRegistered", function () {
        it("should return true for registered and false by default", async function () {
            const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            const isregistered1 = await broker.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isregistered1).to.equal(false)

            const fee = await broker.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(broker.address, fee)
            await broker.connect(anotherUser).depositCoin(fee)
            await broker.connect(anotherUser).registerProvider()

            const isregistered2 = await broker.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isregistered2).to.equal(true)
        })
    })


    describe("registerProvider", function () {
        it("should use exactly PROVIDER_REGISTRATION_FEE coins", async function () {
            const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            const fee = await broker.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee.add(123))
            await token.connect(anotherUser).increaseAllowance(broker.address, fee.add(123))
            await broker.connect(anotherUser).depositCoin(fee.add(123))


            const balance1 = await broker.connect(anotherUser).getCoinBalance(anotherUser.address)
            expect(balance1[0]).to.equal(fee.add(123))
            expect(balance1[1]).to.equal(0)

            await broker.connect(anotherUser).registerProvider()

            const balance2 = await broker.connect(anotherUser).getCoinBalance(anotherUser.address)
            expect(balance2[0]).to.equal(123)
            expect(balance2[1]).to.equal(0)
        });

        it("should reject if not enough coins", async function () {
            const { broker, token, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            const fee = await broker.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee.sub(1))
            await token.connect(anotherUser).increaseAllowance(broker.address, fee.sub(1))
            await broker.connect(anotherUser).depositCoin(fee.sub(1))

            await expect(broker.connect(anotherUser).registerProvider()).to.be.rejected
        });
    });

});
