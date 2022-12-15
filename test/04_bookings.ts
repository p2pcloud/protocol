import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { brokerWithFiveOffers, brokerWithOfferAndUserBalance, deployBrokerFixture, offerFromRaw, OffersItem } from './fixtures'

describe.only("BrokerV1_bookings", function () {
    describe.only("Book", function () {
        it.only("should create a new booking", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(brokerWithOfferAndUserBalance);


            let booked = await broker.FindBookingsByUser(user.address)
            expect(booked.length).is.equal(0)


            await broker.connect(user).Book(0)

            booked = await broker.connect(user).FindBookingsByUser(user.address)
            expect(booked.length).is.equal(1)
        });

        it.only("should revert if not enough free coin balance", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(brokerWithFiveOffers);

            //add some money to user
            await token.connect(admin).transfer(user.address, '1000000000')
            await token.connect(user).approve(broker.address, '1000000000')

            const pps = 10
            await broker.connect(miner).AddOffer(pps, 1, Array(32).fill(0))
            const offers = (await broker.GetAvailableOffers()).map(offerFromRaw)
            const lastOffer = offers[offers.length - 1]

            const moneyNeeded = pps * 3600 * 24 * 7

            //less than needed
            await broker.connect(user).DepositCoin(moneyNeeded - 1)
            await expect(broker.connect(user).Book(lastOffer.Index)).to.be.reverted

            //more than needed
            await broker.connect(user).DepositCoin(2)
            await expect(broker.connect(user).Book(lastOffer.Index)).to.not.be.reverted

        });
        it.only("should increase locked coin balance", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            const pps = 10
            await broker.connect(miner).AddOffer(pps, 1, Array(32).fill(0))
            const offers = (await broker.GetAvailableOffers()).map(offerFromRaw)
            const lastOffer = offers[offers.length - 1]

            const [free, locked] = await broker.GetCoinBalance(user.address)
            expect(locked.toString()).to.equal('0')

            await broker.connect(user).Book(lastOffer.Index)

            const [free2, locked2] = await broker.GetCoinBalance(user.address)
            expect(locked2.toString()).to.equal((pps * 3600 * 24 * 7).toString())
        });
        it.only("should decrease machines available", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(miner).AddOffer(1, 10, Array(32).fill(0))

            let offers = (await broker.GetAvailableOffers()).map(offerFromRaw)
            let lastOfferId = offers[offers.length - 1].Index

            let offer = offerFromRaw(await broker.GetOffer(lastOfferId))
            expect(offer.Availablility).to.equal(10)

            await broker.connect(user).Book(offer.Index)

            offer = offerFromRaw(await broker.GetOffer(lastOfferId))
            expect(offer.Availablility).to.equal(9)
        });
        it("should revert if no machines available", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(miner).AddOffer(1, 1, Array(32).fill(0))

            let offers = (await broker.GetAvailableOffers()).map(offerFromRaw)
            let lastOfferId = offers[offers.length - 1].Index

            //first ok
            await expect(broker.connect(user).Book(lastOfferId)).to.not.be.reverted
            //second reverted
            await expect(broker.connect(user).Book(lastOfferId)).to.be.reverted
        });
    })
    describe.only("Terminate", function () {
        it.only("should revert if booking does not exist", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const REASON = 0

            await expect(broker.connect(user).Terminate(999, REASON)).to.be.reverted
            await expect(broker.connect(user).Terminate(0, REASON)).to.not.be.reverted
        });
        it.only("should revert if user is not the owner", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const REASON = 0

            await expect(broker.connect(anotherUser).Terminate(0, REASON)).to.be.reverted
            await expect(broker.connect(user).Terminate(0, REASON)).to.not.be.reverted
        });
        it.only("should delete booking", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const REASON = 0

            broker.connect(user).Book(0)

            let booked = await broker.connect(user).FindBookingsByUser(user.address)
            expect(booked.length).is.equal(1)

            await expect(broker.connect(user).Terminate(0, REASON)).not.to.be.reverted

            booked = await broker.connect(user).FindBookingsByUser(user.address)
            expect(booked.length).is.equal(0)
        });
        it.only("should increase machines available", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const REASON = 0

            await broker.connect(user).Book(0)

            let offer = offerFromRaw(await broker.GetOffer(0))
            const initialMachinesAvailable = offer.Availablility

            await broker.connect(user).Terminate(0, REASON)
            offer = offerFromRaw(await broker.GetOffer(0))

            expect(offer.Availablility).is.equal(initialMachinesAvailable + 1)
        });
        it.only("should execute payment", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const seconds = 3600 * 24
            const OFFER_ID = 3
            const PPS = offerFromRaw(await broker.GetOffer(OFFER_ID)).PPS

            const [minerBalanceFree,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree.toString()).to.equal('0')

            await broker.connect(user).Book(OFFER_ID)
            await time.increase(3600 * 24);
            await broker.connect(user).Terminate(0, 0)
            const CORRECTION = 1//TODO: right now correction is 1 second

            const [minerBalanceFree2,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree2.toString()).to.equal((PPS * (seconds + CORRECTION)).toString())
        });
    })
    describe("FindBookingsByUser", function () {
        it("should return array of bookings", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            const booked = await broker.FindBookingsByUser(user.address)
            expect(booked.length).is.equal(1)
        });
    })
    describe("FindBookingsByMiner", function () {
        it("should return array of bookings", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            const booked = await broker.FindBookingsByMiner(miner.address)
            expect(booked.length).is.equal(1)
        });
    })
    describe("GetBooking", function () {
        it("should return booking", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            const booking = await broker.GetBooking(0)
            expect(booking.index).is.equal(0)
        });
    })
    describe("ClaimPayment", function () {
        it("should revert if booking does not exist", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            await expect(broker.connect(miner).ClaimPayment(1)).to.be.reverted
        });
        it("should revert if user is not the miner of this booking", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            await expect(broker.connect(user).ClaimPayment(0)).to.be.reverted
        });
        it("should transfer payment to miner and commision to community", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);
            const FEE = 2000
            const week = 86400

            // Contract settings
            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)
            await broker.SetCommunityFee(FEE)

            // Add User offers
            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.connect(user).AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            // Allow broker deposit User tokens
            const amt = ethers.utils.parseUnits('1', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            // Deposit & book
            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)
            const [freeUserBalance, lockedUserBalance] = await broker.GetCoinBalance(user.address)

            // Waiting week
            await time.increase(week);

            // Get booking info
            const [[index, pricePerSecond]] = await broker.FindBookingsByUser(user.address)
            const pricePerWeek = ethers.utils.formatUnits(pricePerSecond.mul(week), 'mwei')

            // Claim payment
            expect(await broker.connect(user).ClaimPayment(index)).not.to.be.reverted

            // Get balances & calc percents
            const [freeClaimedUserBalance, lockedClaimedUserBalance] = await broker.GetCoinBalance(user.address)
            const [free, locked] = await broker.GetCoinBalance(miner.address)

            const userClaimed = ethers.utils.formatUnits(freeUserBalance.sub(freeClaimedUserBalance), 'mwei')
            const communityFee = ethers.utils.formatUnits(free, 'mwei')
            const percents = (+pricePerWeek * FEE) / 10000

            expect(Number(userClaimed)).is.equal(Number(communityFee)).is.equal(percents)
        });
    })
})