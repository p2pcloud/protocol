import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { deployBrokerFixture } from './fixtures'
import { randomBytes } from 'crypto'

if (process.env.REPORT_GAS === "true") {
    describe("Gas", function () {
        it("books many VMs", async function () {
            const TOTAL_ROUNDS = 200
            const TOTAL_OFFERS = 10
            const BOOKING_PER_OFFER = Math.floor(TOTAL_ROUNDS / TOTAL_OFFERS)


            const { broker, token, miner, user, admin } = await loadFixture(deployBrokerFixture);

            //balances for user and miner
            for (let person of [user, miner]) {
                const amt = '100000000000'
                await token.connect(admin).transfer(person.address, amt)
                await token.connect(person).approve(broker.address, amt)
                await broker.connect(person).DepositCoin(amt)
            }

            for (let offerIndex = 0; offerIndex < TOTAL_OFFERS; offerIndex++) {
                //add offer 
                const exampleSpecBytes = "0x" + randomBytes(32).toString('hex')
                await broker.connect(miner).AddOffer(1, 10000, exampleSpecBytes)

                //book
                for (let bookingIndex = 0; bookingIndex < BOOKING_PER_OFFER; bookingIndex++) {
                    await broker.connect(user).Book(offerIndex)
                }
            }

            await time.increase(3600 * 24 * 7);

            for (let bookingId = 0; bookingId < TOTAL_OFFERS * BOOKING_PER_OFFER; bookingId++) {
                await broker.connect(miner).ClaimPayment(bookingId)
            }

            await time.increase(120);

            for (let bookingId = 0; bookingId < TOTAL_OFFERS * BOOKING_PER_OFFER; bookingId++) {
                await broker.connect(user).Terminate(bookingId, 0)
            }

            for (let offerIndex = 0; offerIndex < TOTAL_OFFERS; offerIndex++) {
                await broker.connect(miner).RemoveOffer(offerIndex)
            }

        });
    })
}

