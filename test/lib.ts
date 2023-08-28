import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { MarketplaceFixture } from "./fixtures";
import { BigNumber, BigNumberish, ethers } from "ethers";

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
    const { marketplace, stablecoin, user, admin } = fixture;

    const total = await marketplace.connect(user).getTotalBalance(user.address)

    const diff = BigNumber.from(amt).sub(total)

    if (diff.gt(0)) {
        await stablecoin.connect(admin).transfer(user.address, diff)
        await stablecoin.connect(user).approve(marketplace.address, diff)
        await marketplace.connect(user).depositCoin(diff)
    } else {
        await marketplace.connect(user).withdrawCoin(diff.mul(-1))
    }
}

export function utf8StringToFixedLengthHex(s: string, length: number) {
    if (s.length > length) throw new Error("String too long")
    return ethers.utils.formatBytes32String(s).slice(0, 2 + length * 2)
}

export const US_HEX = utf8StringToFixedLengthHex("US", 2)
export const CA_HEX = utf8StringToFixedLengthHex("CA", 2)
export const NZ_HEX = utf8StringToFixedLengthHex("NZ", 2)

export async function signKYC(userAddress: string, countryHex: string, wallet: SignerWithAddress) {
    const message = ethers.utils.solidityPack(['address', 'bytes2'], [userAddress, countryHex]);

    // hash the message
    const messageHash = ethers.utils.keccak256(message);

    // sign the message
    const signature = await wallet.signMessage(ethers.utils.arrayify(messageHash));

    return signature;
}
export function randomBytes12() {
    return ethers.utils.hexlify(ethers.utils.randomBytes(12))
}