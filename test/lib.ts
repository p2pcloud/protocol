import { Fixture } from "./fixtures";
import { BigNumber, BigNumberish } from "ethers";

export async function setUserCoinBalance(fixture: Fixture, amt: BigNumberish) {
    const { marketplace, token, user, admin } = fixture;

    const [free, locked] = await marketplace.connect(user).getBalance(user.address)
    const total = free.add(locked)

    const diff = BigNumber.from(amt).sub(total)

    if (diff.gt(0)) {
        await token.connect(admin).transfer(user.address, diff)
        await token.connect(user).approve(marketplace.address, diff)
        await marketplace.connect(user).depositCoin(diff)
    } else {
        await marketplace.connect(user).withdrawCoin(diff.mul(-1))
    }
}