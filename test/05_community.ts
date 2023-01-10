import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { BN } from "bn.js";
import { expect } from "chai";
import { ethers } from "hardhat";
import { deployBrokerFixture, brokerWithOfferAndUserBalance, brokerWithFiveOffers } from './fixtures'

describe("Broker_community", function () {
    describe("SetCommunityAddress", function () {
        it("should set community address if address is address(0)", async function () {
            const { broker, user, admin } = await loadFixture(deployBrokerFixture);

            //FIXME: this is a hack to not make a separate fixture
            await broker.connect(admin).SetCommunityContract(ethers.constants.AddressZero)

            expect(await broker.communityContract()).to.equal(ethers.constants.AddressZero);

            expect(await broker.SetCommunityContract(user.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(user.address);
        });
        it("should set community address if community is the sender", async function () {
            const { broker, miner, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            //FIXME: this is a hack to not make a separate fixture
            await broker.connect(admin).SetCommunityContract(anotherUser.address)

            expect(await broker.connect(anotherUser).SetCommunityContract(miner.address)).to.not.be.reverted
            expect(await broker.communityContract()).to.equal(miner.address);
        });
        it("should revert if not owner", async function () {
            const { broker, miner, anotherUser, admin } = await loadFixture(deployBrokerFixture);

            //FIXME: this is a hack to not make a separate fixture
            await broker.connect(admin).SetCommunityContract(miner.address)

            expect(broker.connect(anotherUser).SetCommunityContract(admin.address)).to.be.reverted
            expect(await broker.communityContract()).to.equal(miner.address);
        });
    })
    describe("SetCommunityFee", function () {
        it("should set community fee", async function () {
            const { broker, miner, admin } = await loadFixture(deployBrokerFixture);

            await expect(broker.connect(admin).SetCommunityFee(2000)).to.not.be.reverted
            expect(await broker.communityFee()).to.equal(2000);
        });
        it("should revert if not community contract", async function () {
            const { broker, miner, user } = await loadFixture(deployBrokerFixture);
            await broker.SetCommunityContract(user.address)
            await expect(broker.SetCommunityFee(2000)).to.be.reverted
        });
        it("should revert if fee is not between 0 (0%) and 2500 (25%)", async function () {
            const { broker, miner, admin } = await loadFixture(deployBrokerFixture);
            await expect(broker.connect(admin).SetCommunityFee(0)).to.not.be.reverted
            await expect(broker.connect(admin).SetCommunityFee(2499)).to.not.be.reverted
            await expect(broker.connect(admin).SetCommunityFee(2500)).to.be.reverted
        });
        //FIXME: by some reason this test fails without 
        //it says: This might be caused by using nested loadFixture calls in a test, for example by using multiple beforeEach calls. This is not supported yet.
        it.skip("should change amount of fee paid in ClaimPament ", async function () {
            for (let FEE of [0, 10, 100]) {
                const { broker, miner, admin, user } = await loadFixture(brokerWithOfferAndUserBalance);

                const SECONDS = 3600 * 24
                const OFFER_ID = 4
                const PPS = (await broker.GetOffer(OFFER_ID)).pricePerSecond

                // Contract settings
                await broker.SetCommunityFee(FEE)

                //initial balances
                await broker.connect(user).Book(OFFER_ID)

                const [initialUserBalance] = await broker.GetCoinBalance(user.address)
                const [initialMinerBalance] = await broker.GetCoinBalance(miner.address)
                const [initialCommunityBalance] = await broker.GetCoinBalance(admin.address)

                expect(initialCommunityBalance.toString()).to.equal('0')
                expect(initialMinerBalance.toString()).to.equal('0')

                await time.increase(SECONDS);
                await broker.connect(miner).ClaimPayment(0)

                const [userBalance] = await broker.GetCoinBalance(user.address)
                const [minerBalance] = await broker.GetCoinBalance(miner.address)
                const [communityBalance] = await broker.GetCoinBalance(admin.address)

                const totalCost = PPS * (SECONDS + 1)//TODO: fix this correction

                expect(userBalance.toString()).to.equal(initialUserBalance.sub(totalCost).toString(), 'user balance, fee=' + FEE)

                const communityPayment = new BN(totalCost).mul(new BN(FEE)).div(new BN(10000))
                expect(communityBalance.toString()).to.equal(communityPayment.toString(), 'community balance, fee=' + FEE)

                const minerPayment = new BN(totalCost).sub(communityPayment)
                expect(minerBalance.toString()).to.equal(minerPayment.toString(), 'miner balance, fee=' + FEE)//TODO: fix this correction
            }
        });
    })
})
