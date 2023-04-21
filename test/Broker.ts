import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployMarketplaceFixture } from './fixtures'
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "ethers";
import { UnsignedOffer, setUserCoinBalance, signOffer } from "./lib";

describe("Broker", function () {
    describe("bookResource", function () {
        it("should create a new Booking", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);

            const tx = await marketplace.connect(user).bookResource(offer, signature);
            const rc = await tx.wait();
            const event = rc.events?.find(event => event.event === 'BookingCreated');
            const newId = event?.args?.bookingId;

            const bookingFromChain = await marketplace.getBooking(newId);

            expect(bookingFromChain.specs).to.equal(offer.specs);
            expect(bookingFromChain.pricePerMinute).to.equal(offer.pricePerMinute);
            expect(bookingFromChain.client).to.equal(user.address);
            expect(bookingFromChain.provider).to.equal(provider.address);
        })

        it("should increase locked balance", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };

            const locked1 = await marketplace.connect(user).getLockedBalance(user.address)
            expect(locked1).to.equal(0);

            const signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            const locked2 = await marketplace.connect(user).getLockedBalance(user.address)
            expect(locked2).to.equal(100 * 60 * 24 * 7);
        })
        it("should execute claim payment", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 2,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            //wait for 1 day
            const initialTime = await time.latest();
            await time.increase(3600 * 24);

            //balance is still 0
            const balance1 = await marketplace.connect(provider).getFreeBalance(provider.address);
            expect(balance1).to.equal(0);

            //booking another resource will trigger claim payment
            offer.nonce = await marketplace.getNonce(user.address);
            const signature2 = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature2);

            //check balance
            const wholeMinutesPassed = Math.floor((await time.latest() - initialTime) / 60);
            const balance2 = await marketplace.connect(provider).getFreeBalance(provider.address);
            const payoutPercent = (10000 - await marketplace.communityFee()) / 10000
            expect(balance2).to.equal(offer.pricePerMinute * wholeMinutesPassed * payoutPercent);//pricePerMinute * wholeMinutesPassed
        })
        it("should fail if user tries to reuse the same offer", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };

            const locked1 = await marketplace.connect(user).getLockedBalance(user.address)
            expect(locked1).to.equal(0);

            const signature = await signOffer(provider, offer, marketplace.address);
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.not.be.reverted;
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.be.revertedWith("Invalid offer nonce");
        })
        it("should fail if not enough balance to cover a new VM booking", async function () {
            const fixture = await loadFixture(deployMarketplaceFixture);
            const { marketplace, user, provider, token } = fixture

            //remove default balance and deposit some tokens
            const defaultBalance = await marketplace.connect(user).getFreeBalance(user.address)

            //add an offer
            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 1,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            let signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            //wait for 1 day
            const initialTime = await time.latest();
            await time.increase(3600 * 24);
            const wholeMinutesPassed = Math.floor((await time.latest() - initialTime) / 60);
            const moneyToBePaid = offer.pricePerMinute * wholeMinutesPassed;
            const lockedBalance = await marketplace.connect(user).getLockedBalance(user.address)

            const minBalanceToBook = lockedBalance
                .add(moneyToBePaid)
                .add(offer.pricePerMinute * 60 * 24 * 7);

            //prepeare a new offer
            offer.nonce++
            signature = await signOffer(provider, offer, marketplace.address);

            await setUserCoinBalance(fixture, minBalanceToBook.sub(1));
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.be.revertedWith("Not enough balance to add a new the booking");

            await setUserCoinBalance(fixture, minBalanceToBook);
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.not.be.reverted;
        })
        it("should fail if provider is not registered", async function () {
            const { marketplace, user, anotherUser } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };

            const locked1 = await marketplace.connect(user).getLockedBalance(user.address)
            expect(locked1).to.equal(0);

            const signature = await signOffer(anotherUser, offer, marketplace.address);
            await expect(marketplace.connect(user).bookResource(offer, signature)).to.be.revertedWith("Provider is not registered");
        })
    })
    describe("cancelBooking", function () {
        it("should remove a booking and decrease locked balance", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);

            await marketplace.connect(user).bookResource(offer, signature);

            //check booking
            let bookingFromChain = await marketplace.getBooking(0);
            expect(bookingFromChain.specs).to.equal(offer.specs);

            //check locked balance
            const locked1 = await marketplace.connect(user).getLockedBalance(user.address)
            expect(locked1).to.equal(offer.pricePerMinute * 60 * 24 * 7);

            //delete
            await marketplace.connect(user).cancelBooking(0, true);

            //check deleted
            bookingFromChain = await marketplace.getBooking(0);
            expect(bookingFromChain.specs).to.equal(ethers.utils.formatBytes32String(""));

            //check locked balance
            const locked2 = await marketplace.connect(user).getLockedBalance(user.address)
            expect(locked2).to.equal(0);
        })
        it("should fail if booking is already cancelled", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);

            await marketplace.connect(user).bookResource(offer, signature);

            //delete
            await marketplace.connect(user).cancelBooking(0, true);
            //once more
            await expect(marketplace.connect(user).cancelBooking(0, true)).to.be.revertedWith("Only client or provider can cancel the booking");
        })
        it("should create an event with reason = 0 or 1 depending on client's selection", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            //book vm 0
            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            //book vm 1
            offer.nonce++
            const signature2 = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature2);

            //delete CANCEL_REASON_NOT_SATISFIED
            const tx = await marketplace.connect(user).cancelBooking(0, false);
            await expect(tx).to.emit(marketplace, "BookingCancelled").withArgs(0, await marketplace.CANCEL_REASON_NOT_SATISFIED());

            //delete CANCEL_REASON_NOT_NEEDED
            const tx2 = await marketplace.connect(user).cancelBooking(1, true);
            await expect(tx2).to.emit(marketplace, "BookingCancelled").withArgs(1, await marketplace.CANCEL_REASON_NOT_NEEDED());
        })
        it("should create an event with reason = 2 if provider cancelled or reason = 3 if client defauled", async function () {
            const fixture = await loadFixture(deployMarketplaceFixture);
            const { marketplace, user, provider } = fixture;

            //book vm 0
            const offer: UnsignedOffer = {
                specs: ethers.utils.formatBytes32String("hello world"),
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
                nonce: await marketplace.getNonce(user.address),
            };
            const signature = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature);

            //set user balance to have enough to book 2 vms
            await setUserCoinBalance(fixture, offer.pricePerMinute * await marketplace.MONEY_LOCK_MINUTES() * 2);

            //book vm 1
            offer.nonce++
            const signature2 = await signOffer(provider, offer, marketplace.address);
            await marketplace.connect(user).bookResource(offer, signature2);

            //delete CANCEL_REASON_PROVIDER
            const tx = await marketplace.connect(provider).cancelBooking(0, true);
            await expect(tx).to.emit(marketplace, "BookingCancelled").withArgs(0, await marketplace.CANCEL_REASON_PROVIDER());

            //skip 10 weeks
            await time.increase(60 * await marketplace.MONEY_LOCK_MINUTES() * 2);

            //delete CANCEL_REASON_NON_PAYMENT
            const tx2 = await marketplace.connect(user).cancelBooking(1, false);//doesn't matter what client selects
            expect(await marketplace.connect(user).getTotalBalance(user.address)).to.equal(0);
            await expect(tx2).to.emit(marketplace, "BookingCancelled").withArgs(1, await marketplace.CANCEL_REASON_NON_PAYMENT());
        })
    })
    describe("claimPayment", function () {
        it("should increase provider's balance")
        it("should transfer not more than user's balance")
    });

});
