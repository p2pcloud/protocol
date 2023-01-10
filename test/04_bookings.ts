import { expect } from "chai";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { brokerWithFiveOffers, brokerWithOfferAndUserBalance, } from './fixtures'
import { BN } from "bn.js";
import { ethers } from "ethers";

describe("Broker_bookings", function () {
    describe("Book", function () {
        it("should create a new booking", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(brokerWithOfferAndUserBalance);


            let booked = await broker.FindBookingsByUser(user.address)
            expect(booked.length).is.equal(0)


            await broker.connect(user).Book(0)

            booked = await broker.connect(user).FindBookingsByUser(user.address)
            expect(booked.length).is.equal(1)
        });

        it("should emit NewBooking event", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(brokerWithOfferAndUserBalance);
            const bookingId = 3

            const pps = (await broker.GetOffer(bookingId)).pricePerSecond

            await expect(broker.connect(user).Book(bookingId))
                .to.emit(broker, 'NewBooking')
                .withArgs(user.address, miner.address, 0, pps)
        });

        it("should revert if not enough free coin balance", async function () {
            const { broker, token, miner, user, admin } = await loadFixture(brokerWithFiveOffers);

            //add some money to user
            await token.connect(admin).transfer(user.address, '1000000000')
            await token.connect(user).approve(broker.address, '1000000000')

            const pps = 10
            await broker.connect(miner).AddOffer(pps, 1, Array(32).fill(0))
            const offers = (await broker.GetAvailableOffers())
            const lastOffer = offers[offers.length - 1]

            const moneyNeeded = pps * 3600 * 24 * 7

            //less than needed
            await broker.connect(user).DepositCoin(moneyNeeded - 1)
            await expect(broker.connect(user).Book(lastOffer.index)).to.be.reverted

            //more than needed
            await broker.connect(user).DepositCoin(2)
            await expect(broker.connect(user).Book(lastOffer.index)).to.not.be.reverted

        });
        it("should increase locked coin balance", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            const pps = 10
            await broker.connect(miner).AddOffer(pps, 1, Array(32).fill(0))
            const offers = (await broker.GetAvailableOffers())
            const lastOffer = offers[offers.length - 1]

            const [free, locked] = await broker.GetCoinBalance(user.address)
            expect(locked.toString()).to.equal('0')

            await broker.connect(user).Book(lastOffer.index)

            const [free2, locked2] = await broker.GetCoinBalance(user.address)
            expect(locked2.toString()).to.equal((pps * 3600 * 24 * 7).toString())
        });
        it("should increase machines booked", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(miner).AddOffer(1, 10, Array(32).fill(0))

            let offers = (await broker.GetAvailableOffers())
            let lastOfferId = offers[offers.length - 1].index

            let offer = await broker.GetOffer(lastOfferId)
            expect(offer.machinesBooked).to.equal(0)

            await broker.connect(user).Book(offer.index)

            offer = await broker.GetOffer(lastOfferId)
            expect(offer.machinesBooked).to.equal(1)
        });
        it("should revert if no machines available", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(miner).AddOffer(1, 1, Array(32).fill(0))

            let offers = (await broker.GetAvailableOffers())
            let lastOfferId = offers[offers.length - 1].index

            //first ok
            await expect(broker.connect(user).Book(lastOfferId)).to.not.be.reverted
            //second reverted
            await expect(broker.connect(user).Book(lastOfferId)).to.be.reverted
        });
    })
    describe("Terminate", function () {
        it("should revert if booking does not exist", async function () {
            const { broker, token, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const REASON = 0

            await expect(broker.connect(user).Terminate(999, REASON)).to.be.reverted
            await expect(broker.connect(user).Terminate(0, REASON)).to.not.be.reverted
        });
        it("should revert if user is not the owner", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const REASON = 0

            await expect(broker.connect(anotherUser).Terminate(0, REASON)).to.be.reverted
            await expect(broker.connect(user).Terminate(0, REASON)).to.not.be.reverted
        });
        it("should delete booking", async function () {
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
        it("should emit Terminate event", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)

            const REASON = 123
            await expect(broker.connect(user).Terminate(0, REASON))
                .to.emit(broker, 'Termination')
                .withArgs(0, REASON);
        });
        it("should decrease machines booked", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const REASON = 0

            await broker.connect(user).Book(0)

            let offer = await broker.GetOffer(0)
            const initialMachinesBooked = offer.machinesBooked

            await broker.connect(user).Terminate(0, REASON)
            offer = await broker.GetOffer(0)

            expect(offer.machinesBooked).is.equal(initialMachinesBooked - 1)
        });
        it("should execute payment", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const SECONDS = 3600 * 24
            const OFFER_ID = 3
            const PPS = (await broker.GetOffer(OFFER_ID)).pricePerSecond

            const [minerBalanceFree,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree.toString()).to.equal('0')

            await broker.connect(user).Book(OFFER_ID)
            await time.increase(3600 * 24);
            await broker.connect(user).Terminate(0, 0)
            const CORRECTION = 1//TODO: right now correction is 1 second

            const [minerBalanceFree2,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree2.toString()).to.equal((PPS * (SECONDS + CORRECTION)).toString())
        });

        it("should work for miner too", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const SECONDS = 3600 * 24
            const OFFER_ID = 3
            const PPS = (await broker.GetOffer(OFFER_ID)).pricePerSecond

            const [minerBalanceFree,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree.toString()).to.equal('0')

            await broker.connect(user).Book(OFFER_ID)
            await time.increase(3600 * 24);
            await broker.connect(miner).Terminate(0, 2)
            const CORRECTION = 1//TODO: right now correction is 1 second

            const [minerBalanceFree2,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree2.toString()).to.equal((PPS * (SECONDS + CORRECTION)).toString())
        });

        it("should allow communityContract to terminate anything", async function () {
            const { broker, admin, miner, user } = await loadFixture(brokerWithOfferAndUserBalance);

            const SECONDS = 3600 * 24
            const OFFER_ID = 3
            const PPS = (await broker.GetOffer(OFFER_ID)).pricePerSecond

            const [minerBalanceFree,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree.toString()).to.equal('0')

            await broker.connect(user).Book(OFFER_ID)
            await time.increase(3600 * 24);

            await broker.connect(admin).Terminate(0, await broker.REASON_COMMUNITY_TERMINATED())

            const CORRECTION = 1//TODO: right now correction is 1 second

            const [minerBalanceFree2,] = await broker.GetCoinBalance(miner.address)
            expect(minerBalanceFree2.toString()).to.equal((PPS * (SECONDS + CORRECTION)).toString())
        });

        it("should throw exception for miner if code not equal 2", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const OFFER_ID = 3

            await broker.connect(user).Book(OFFER_ID)
            await time.increase(3600 * 24);

            await expect(broker.connect(miner).Terminate(0, 0)).to.be.reverted
            await expect(broker.connect(miner).Terminate(0, 1)).to.be.reverted
            await expect(broker.connect(miner).Terminate(0, 3)).to.be.reverted

            await broker.connect(miner).Terminate(0, 2)
        });
    })
    describe("FindBookingsByUser", function () {
        it("should return array of bookings", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)
            await broker.connect(anotherUser).Book(1)

            const booked = (await broker.FindBookingsByUser(user.address))
            expect(booked.length).is.equal(1)
            expect(booked[0].offerIndex).is.equal(0)
            expect(booked[0].user).is.equal(user.address)

        });
    })
    describe("FindBookingsByMiner", function () {
        it("should return array of bookings", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)
            await broker.connect(anotherUser).Book(1)

            let booked = (await broker.FindBookingsByMiner(anotherUser.address))
            expect(booked.length).is.equal(0)

            booked = (await broker.FindBookingsByMiner(miner.address))
            expect(booked.length).is.equal(2)
            expect(booked[0].offerIndex).is.equal(0)
            expect(booked[0].user).is.equal(user.address)
        });
    })
    describe("GetBooking", function () {
        it("should return booking", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)
            await broker.connect(anotherUser).Book(1)

            const booking = await broker.GetBooking(0)
            expect(booking.offerIndex).is.equal(0)
            expect(booking.user).is.equal(user.address)
        });
    })
    describe("ClaimPayment", function () {
        it("should revert if booking does not exist", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);
            await time.increase(3600 * 24);
            expect(broker.connect(miner).ClaimPayment(0)).to.be.reverted
        });
        it('should terminate booking if not enough deposit to pay for it', async function () {
            const { broker, token, miner, admin, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const FEE = 500//5%
            const OFFER_ID = 2

            // Contract settings
            await broker.SetCommunityFee(FEE)
            const pps = (await broker.GetOffer(OFFER_ID)).pricePerSecond
            const [free, locked] = await broker.GetCoinBalance(user.address)
            const seconds = free.div(pps).toNumber()

            await broker.connect(user).Book(OFFER_ID)

            //enough money, not terminated
            await time.increase(seconds - 115);

            await expect(broker.connect(miner).ClaimPayment(0))
                .to.not.emit(broker, 'Termination')

            const booking1 = await broker.GetBooking(0)
            expect(booking1.miner).to.not.equal(ethers.constants.AddressZero)

            //120 seconds later - enough
            await time.increase(120);
            await expect(broker.connect(miner).ClaimPayment(0))
                .to.emit(broker, 'Termination')
                .withArgs(0, await broker.REASON_NON_PAYMENT);

            const booking2 = await broker.GetBooking(0)
            expect(booking2.miner).to.equal(ethers.constants.AddressZero)
        })
        it("should revert if user is not the miner of this booking", async function () {
            const { broker, token, miner, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            await broker.connect(user).Book(0)
            await time.increase(3600 * 24);

            expect(broker.connect(anotherUser).ClaimPayment(0)).to.be.reverted
            expect(broker.connect(miner).ClaimPayment(0)).to.not.be.reverted
        });
        it("should transfer payment to miner and commision to community", async function () {
            const { broker, token, miner, admin, user, anotherUser } = await loadFixture(brokerWithOfferAndUserBalance);

            const FEE = 500//5%
            const SECONDS = 3600 * 24
            const OFFER_ID = 2
            const PPS = (await broker.GetOffer(OFFER_ID)).pricePerSecond

            // Contract settings
            await broker.SetCommunityFee(FEE)

            //initial balances
            await broker.connect(user).Book(OFFER_ID)

            const [initialUserBalance] = await broker.GetCoinBalance(user.address)
            const [initialMinerBalance] = await broker.GetCoinBalance(miner.address)
            const [initialCommunityBalance] = await broker.GetCoinBalance(admin.address)

            const registrationFee = await broker.MINER_REGISTRATION_FEE()

            expect(initialMinerBalance.toString()).to.equal('0')
            expect(initialCommunityBalance.toString()).to.equal(registrationFee)

            await time.increase(SECONDS);
            await broker.connect(miner).ClaimPayment(0)

            const [userBalance] = await broker.GetCoinBalance(user.address)
            const [minerBalance] = await broker.GetCoinBalance(miner.address)
            const [communityBalance] = await broker.GetCoinBalance(admin.address)

            const totalCost = PPS * (SECONDS + 1)//TODO: fix this correction

            expect(userBalance.toString()).to.equal(initialUserBalance.sub(totalCost).toString())

            const communityPayment = new BN(totalCost).mul(new BN(FEE)).div(new BN(10000))
            expect(communityBalance).to.equal(registrationFee.add(communityPayment.toString()))

            const minerPayment = new BN(totalCost).sub(communityPayment)
            expect(minerBalance.toString()).to.equal(minerPayment.toString())
        });
    })
})