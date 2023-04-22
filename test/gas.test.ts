import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceFixture } from './fixtures'
import { randomBytes } from 'crypto'
import { UnsignedOffer, setUserCoinBalance, signOffer } from "./lib";
import { ethers } from "ethers";

const GAS_TEST_RUNS = parseInt(process.env.GAS_TEST_RUNS || "100")


if (process.env.REPORT_GAS === "true") {
    describe("Gas", function () {
        it("books many VMs", async function () {
            const fixture = await loadFixture(deployMarketplaceFixture);
            const { marketplace, provider, user, providersSigner } = fixture;

            await setUserCoinBalance(fixture, '100000000000')

            for (let i = 0; i < GAS_TEST_RUNS; i++) {
                const offer: UnsignedOffer = {
                    specs: ethers.utils.formatBytes32String("hello world " + i),
                    pricePerMinute: i,
                    client: user.address,
                    expiresAt: await time.latest() + 60 * 60 * 24 * 7,
                    nonce: await marketplace.getNonce(user.address),
                };
                const signature = await signOffer(providersSigner, offer, marketplace.address);
                await marketplace.connect(user).bookResource(offer, signature);
                await time.increase(3600);
            }

            for (let i = 0; i < GAS_TEST_RUNS; i++) {
                await marketplace.connect(provider).claimPayment(user.address)
                await time.increase(3600);
            }

            for (let i = 0; i < GAS_TEST_RUNS; i++) {
                if (i % 2 === 0) {
                    await marketplace.connect(user).cancelBooking(i, true)
                } else {
                    await marketplace.connect(provider).cancelBooking(i, false)
                }
                await time.increase(3600);
            }

        });
    })
}