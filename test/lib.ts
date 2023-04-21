import { Fixture } from "./fixtures";
import { BigNumber, BigNumberish } from "ethers";

export async function getEvent<T>(tx: Promise<any>, eventName: string): Promise<T> {
    const receipt = await tx;
    const event = receipt.events.find((e: any) => e.event === eventName);
    return event.args;
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