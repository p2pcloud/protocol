import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture, deployOffersFixture, offerFromRaw, offers, OffersItem } from './fixtures'

describe("BrokerV1_bookings", function () {
    describe("Book", function () {
        it("should create a new booking", async function () {
            const { broker, token, miner, user } = await loadFixture(deployBrokerFixture);

            await broker.SetCommunityContract(miner.address)
            await broker.SetCoinAddress(token.address)

            const amt = ethers.utils.parseUnits('5', 'mwei')
            await token.transfer(user.address, amt)
            await token.connect(user).approve(broker.address, amt)
            const allowance = await token.allowance(user.address, broker.address)

            await Promise.all(offers.map(async (offer) => {
                const [pricePerSecond, vmTypeId, machinesAvailable]: OffersItem = offer
                return await broker.AddOffer(pricePerSecond, vmTypeId, machinesAvailable);
            }))

            await broker.connect(user).DepositCoin(allowance)
            await broker.connect(user).Book(0)

            const booked = await broker.connect(user).FindBookingsByUser(user.address)
            expect(booked.length).is.equal(1)
        });
        it("should revert if not enough free coin balance", async function () {
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
            await expect(broker.connect(user).Book(1)).to.be.reverted
        });
        it("should increase locked coin balance", async function () {
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
            const [free, locked] = await broker.GetCoinBalance(user.address)
            expect(locked.isZero()).false
        });
        it("should decrease machines available", async function () {
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

            const offersRaw = await broker.GetMinersOffers(miner.address)
            const offersObj = offersRaw.map(offerFromRaw)
            const unavailable = offersObj.find(({ Availablility }) => !Availablility)
            expect(unavailable?.Index).is.equal(0)
        });
        it("should revert if no machines available", async function () {
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

            await expect(broker.connect(user).Book(0)).to.be.reverted
        });
    })
    describe("Terminate", function () {
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
            await expect(broker.connect(user).Terminate(1, 0)).to.be.reverted
        });
        it("should revert if user is not the owner", async function () {
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
            await expect(broker.connect(miner).Terminate(0, 0)).to.be.reverted
        });
        it("should delete booking", async function () {
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

            await expect(broker.connect(user).Terminate(0, 0)).not.to.be.reverted

            const booked = await broker.connect(user).FindBookingsByUser(user.address)
            expect(booked.length).is.equal(0)
        });
        it("should increase machines available", async function () {
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

            await expect(broker.connect(user).Terminate(0, 0)).not.to.be.reverted

            const offersRaw = await broker.GetMinersOffers(miner.address)
            const offersObj = offersRaw.map(offerFromRaw)
            const available = offersObj.every(({ Availablility }) => Availablility !== 0)
            expect(available).true
        });
        it("should execute payment", async function () {
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

            await expect(broker.connect(user).Terminate(0, 0)).not.to.be.reverted

            const [free, locked] = await broker.GetCoinBalance(user.address)
            expect(locked.isZero()).true
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