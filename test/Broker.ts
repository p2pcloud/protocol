import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployMarketplaceFixture } from './fixtures'
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

const HARDHAT_NETWORK_ID = 31337;

type UnsignedOffer = {
    specs: string;
    pricePerMinute: number;
    client: string;
    expiresAt: number;
}

async function signOffer(
    provider: SignerWithAddress,
    offer: UnsignedOffer,
    brokerAddress: string,
): Promise<string> {
    const domain = {
        chainId: HARDHAT_NETWORK_ID,
        name: 'p2pcloud.io',
        verifyingContract: brokerAddress,
        version: '2',
    }

    const types = {
        UnsignedOffer: [
            { name: 'specs', type: 'bytes32' },
            { name: 'pricePerMinute', type: 'uint256' },
            { name: 'client', type: 'address' },
            { name: 'expiresAt', type: 'uint256' },
        ]
    }

    return provider._signTypedData(domain, types, offer)
}

describe("Broker", function () {
    describe("BookVM", function () {
        it.only("should create a new Booking", async function () {
            const { marketplace, user, provider } = await loadFixture(deployMarketplaceFixture);

            const offer: UnsignedOffer = {
                specs: "0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9",
                pricePerMinute: 100,
                client: user.address,
                expiresAt: Math.floor(Date.now() / 1000) + 60 * 60 * 24 * 7,
            };

            const signature = await signOffer(provider, offer, marketplace.address);

            const tx = await marketplace.connect(user).Book(offer, signature);
            const rc = await tx.wait();

            const event = rc.events?.find(event => event.event === 'BookingCreated');
            const newId = event?.args?.bookingId;

            const bookingFromChain = await marketplace.getBooking(newId);

            expect(bookingFromChain.specs).to.equal(offer.specs);
            expect(bookingFromChain.pricePerMinute).to.equal(offer.pricePerMinute);
            expect(bookingFromChain.client).to.equal(user.address);
            expect(bookingFromChain.provider).to.equal(provider.address);
        })
    })
})
