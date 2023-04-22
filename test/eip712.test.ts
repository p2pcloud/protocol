import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers"
import { ethers, upgrades } from "hardhat";
import { signOffer } from "./lib";
import { expect } from "chai";

describe("EIP712", () => {
    it("should sign an offer", async () => {
        //not really a test, just a way to track changes in the signature
        const privateKey = "0x4db495c188a75e74f70235e318b458d4e5fc16090d1f9583c226e5365508a724"
        const brokerAddress = "0x77761266ac6d623931bB8aB2BD31B2375A0d4c55"
        const userAddress = "0x6663435d35bdDA6dd7e196d3F92dEeb9B7fC10Fe"
        const specs = "0x68656c6c6f20776f726c64000000000000000000000000000000000000000000"

        //create jsonrpcsigner from private key
        const signer = new ethers.Wallet(privateKey, ethers.provider)
        //@ts-ignore
        const signerWithAddress = await SignerWithAddress.create(signer)

        const signature = await signOffer(signerWithAddress, {
            specs: specs,
            pricePerMinute: 100,
            client: userAddress,
            expiresAt: 10000000,
            nonce: 7,
        }, brokerAddress)

        expect(signature).to.equal("0x3ed4c93bf2b13d37208a29cc79ef6580ae8563ec2ba16e82f311219577cfb7d65c52f9e28508c928857c62718a0347ecdd1b04cb3f98662a476cf466eede64d51b")
    })
})