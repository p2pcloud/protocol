import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { Fixture } from "./fixtures";
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
            { name: 'nonce', type: 'uint32' },
        ]
    }

    return provider._signTypedData(domain, types, offer)
}


export async function setUserCoinBalance(fixture: Fixture, amt: BigNumberish) {
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