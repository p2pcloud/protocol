import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { MarketplaceFixture } from "./fixtures";
import { BigNumber, BigNumberish } from "ethers";

export async function getEvent<T>(tx: Promise<any>, eventName: string): Promise<T> {
    const receipt = await tx;
    const event = receipt.events.find((e: any) => e.event === eventName);
    return event.args;
}

const HARDHAT_NETWORK_ID = 31337;

export type UnsignedOffer = {
    specs: string;
    pricePerMinute: number;
    client: string;
    expiresAt: number;
    nonce: number;
}

export async function signOffer(
    signer: SignerWithAddress,
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
            { name: 'client', type: 'address' },
            { name: 'pricePerMinute', type: 'uint64' },
            { name: 'nonce', type: 'uint32' },
            { name: 'specs', type: 'bytes32' },
            { name: 'expiresAt', type: 'uint256' },
        ]
    }

    return signer._signTypedData(domain, types, offer)
}

export type UnsignedVoucher = {
    amount: number;
    paymentId: string;
}

export async function signVoucher(
    signer: SignerWithAddress,
    voucher: UnsignedVoucher,
    brokerAddress: string,
): Promise<string> {
    const domain = {
        chainId: HARDHAT_NETWORK_ID,
        name: 'p2pcloud.io',
        verifyingContract: brokerAddress,
        version: '2',
    }

    const types = {
        UnsignedVoucher: [
            { name: 'amount', type: 'uint256' },
            { name: 'paymentId', type: 'bytes32' },
        ]
    }

    return signer._signTypedData(domain, types, voucher)
}

export async function setUserCoinBalance(fixture: MarketplaceFixture, amt: BigNumberish) {
    const { marketplace, token, user, admin } = fixture;

    const total = await marketplace.connect(user).getTotalBalance(user.address)

    const diff = BigNumber.from(amt).sub(total)

    if (diff.gt(0)) {
        await token.connect(admin).transfer(user.address, diff)
        await token.connect(user).approve(marketplace.address, diff)
        await marketplace.connect(user).depositCoin(diff)
    } else {
        await marketplace.connect(user).withdrawCoin(diff.mul(-1))
    }
}