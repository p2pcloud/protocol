import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { deployBrokerFixture } from './fixtures'

const HARDHAT_NETWORK_ID = 31337;


describe("BalanceHolder", function () {
    describe("createAgreement", function () {
        it.only("should create a new agreement", async function () {
            const { broker, user, provider } = await loadFixture(deployBrokerFixture);

            const agreement = {
                agreementId: 0,
                ipfsHash: "0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9",
                pricePerMinute: 100,
            };


            const domain = {
                chainId: HARDHAT_NETWORK_ID,
                name: 'p2pcloud.io',
                verifyingContract: broker.address,
                version: '2',
            }

            const types = {
                Amendment: [
                    { name: 'agreementId', type: 'uint256' },
                    { name: 'ipfsHash', type: 'bytes32' },
                    { name: 'pricePerMinute', type: 'uint256' },
                ]
            }

            const signature = provider._signTypedData(domain, types, agreement)
            
            await broker.connect(user).createAgreement(agreement, signature);

            const agreementFromChain = await broker.getAgreement(1);

            expect(agreementFromChain.ipfsHash).to.equal(agreement.ipfsHash);
            expect(agreementFromChain.pricePerMinute).to.equal(agreement.pricePerMinute);
            expect(agreementFromChain.client).to.equal(user.address);
            expect(agreementFromChain.provider).to.equal(provider.address);
        })
    })
    describe("amendAgreement", function () {
        it("should create a new agreement if agreementId = 0 ??")
        it("should amend an existing agreement")
        it("should not amend an existing agreement if agreementId is not correct")
        it("should not amend an existing agreement if the new agreement is not signed by the cprovider")
        it("should not amend an existing agreement owned by another client")
    })
    describe("listClientsAgreements", function () {
        it("should list all agreements for a client")
    })
})
