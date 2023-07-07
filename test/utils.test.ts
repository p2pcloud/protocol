import { expect } from "chai";
import { utf8StringToFixedLengthHex } from "./lib";

describe("utf8StringToFixedLengthHex", () => {
    it("has to return 0x5553 for US with length 3", async () => {
        expect(await utf8StringToFixedLengthHex("US", 2)).to.equal("0x5553")
    })
    it("has to return 0x464c00 for FL with length 3", async () => {
        expect(await utf8StringToFixedLengthHex("FL", 3)).to.equal("0x464c00")
    })
    it("has to return 0x414243 for ABС with length 3", async () => {
        expect(await utf8StringToFixedLengthHex("ABC", 3)).to.equal("0x414243")
    })
})