import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { DEFAULT_USER_BALANCE, deployCreditMarketplaceV3Fixture, deployMarketplaceV3Fixture } from './fixtures'
import { NZ_HEX, US_HEX, signKYC } from "./lib";
import { randomBytes12 } from "./lib";
import { ethers } from "ethers";

describe("BalanceHolder - stablecoin", function () {
    describe("depositCoin", function () {
        it("should incrase balance", async function () {
            const { marketplace, stablecoin, user } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            const [userBalance1] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance1.toString()).is.equal('0')

            await stablecoin.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123')

            const [userBalance2] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance2.toString()).is.equal('123')
        });

        it("should revert if transfer fails", async function () {
            const { marketplace, stablecoin, user } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            const [userBalance1] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance1.toString()).is.equal('0')

            await stablecoin.connect(user).approve(marketplace.address, '123456')

            await stablecoin.transferShouldFail(true)

            await expect(marketplace.connect(user).depositCoin('123'))
                .to.be.revertedWithCustomError(marketplace, "ERC20TransferFailed")
        });

        it("should revert if coin is empty", async function () {
            const { marketplace, admin, user } = await loadFixture(deployCreditMarketplaceV3Fixture);
            await expect(marketplace.connect(user).depositCoin('123'))
                .to.be.revertedWithCustomError(marketplace, "DirectMintBurnOnly")
        });

        it("should require user's KYC", async function () {
            const { marketplace, admin, stablecoin, anotherUser, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await stablecoin.connect(admin).transfer(anotherUser.address, '123456')
            await stablecoin.connect(anotherUser).approve(marketplace.address, '123456')

            await expect(marketplace.connect(anotherUser).depositCoin('123456'))
                .to.be.revertedWithCustomError(marketplace, "KYCProblem")
                .withArgs(anotherUser.address, "0x0000")

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, NZ_HEX,
                await signKYC(anotherUser.address, NZ_HEX, kycSigner)
            )

            await expect(marketplace.connect(anotherUser).depositCoin('123456'))
                .to.be.revertedWithCustomError(marketplace, "KYCProblem")
                .withArgs(anotherUser.address, NZ_HEX)

            await marketplace.connect(admin).allowUserCountry(NZ_HEX)

            await marketplace.connect(anotherUser).depositCoin('123456')
        });
    })
    describe("WithdrawCoin", function () {
        it("should not withdraw locked balance", async function () {
            const { marketplace, stablecoin, user } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const [defaultBalance] = await marketplace.connect(user).getBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123456 tokens
            await stablecoin.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '1')
            const [free1, locked1] = await marketplace.connect(user).getBalance(user.address)
            expect(locked1.toString()).is.equal('10080')
            expect(free1.toString()).is.equal(String(123456 - 10080))

            //fail to withdraw free+1
            await expect(marketplace.connect(user).withdrawCoin(123456 - 10080 + 1))
                .to.be.revertedWithCustomError(marketplace, "InsufficientBalance")
                .withArgs(123456 - 10080 + 1, 123456 - 10080)

            //withdraw excatly free
            await marketplace.connect(user).withdrawCoin(123456 - 10080)
        });

        it("should require user's KYC", async function () {
            const { marketplace, admin, stablecoin, anotherUser, kycSigner, user } = await loadFixture(deployMarketplaceV3Fixture);

            //change user's country to NZ
            await marketplace.connect(user).submitKYC(
                user.address, NZ_HEX,
                await signKYC(user.address, NZ_HEX, kycSigner)
            )

            await expect(marketplace.connect(anotherUser).withdrawCoin('123456'))
                .to.be.revertedWithCustomError(marketplace, "KYCProblem")
                .withArgs(anotherUser.address, "0x0000")

            //change user's country to US
            await marketplace.connect(user).submitKYC(
                user.address, US_HEX,
                await signKYC(user.address, US_HEX, kycSigner)
            )

            await marketplace.connect(user).withdrawCoin('123456')
        });

        it("should revert if coin is empty", async function () {
            const { marketplace, user } = await loadFixture(deployCreditMarketplaceV3Fixture);
            await expect(marketplace.connect(user).withdrawCoin('123'))
                .to.be.revertedWithCustomError(marketplace, "DirectMintBurnOnly")
        });

        it("should revert if transfer fails", async function () {
            //TODO: modify coin contract to make this test pass
            const { marketplace, stablecoin, user } = await loadFixture(deployMarketplaceV3Fixture);

            const [userBalance1] = await marketplace.connect(user).getBalance(user.address)

            await stablecoin.transferShouldFail(true)
            await expect(marketplace.connect(user).withdrawCoin(1))
                .to.be.revertedWithCustomError(marketplace, "ERC20TransferFailed")

            const [userBalance2] = await marketplace.connect(user).getBalance(user.address)
            expect(userBalance2.toString()).is.equal(userBalance1.toString())

        });
    })
    describe("getBalance", function () {
        it("should return locked and free balance", async function () {
            const { marketplace, stablecoin, user } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const defaultBalance = await marketplace.connect(user).getFreeBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123 tokens
            await stablecoin.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '1')

            const [free, locked] = await marketplace.connect(user).getBalance(user.address)
            expect(free.toString()).is.equal(String(123456 - 60 * 24 * 7))
            expect(locked.toString()).is.equal(String(60 * 24 * 7))
        });

        it("should not fail if locked is more than free", async function () {
            const { marketplace, stablecoin, user } = await loadFixture(deployMarketplaceV3Fixture);

            //remove default balance
            const defaultBalance = await marketplace.connect(user).getFreeBalance(user.address)
            await marketplace.connect(user).withdrawCoin(defaultBalance)

            //deposit 123 tokens
            await stablecoin.connect(user).approve(marketplace.address, '123456')
            await marketplace.connect(user).depositCoin('123456')

            //lock 1 * 60 * 24 * 7 tokens
            await marketplace.test__increaseSpendingPerMinute(user.address, '10000')

            const [free, locked] = await marketplace.connect(user).getBalance(user.address)
            expect(free).is.equal(0)
            expect(locked).is.equal(60 * 24 * 7 * 10000)
        })
    })
});

describe("BalanceHolder - credits", function () {
    describe("mint", () => {
        it("should increase balance", async function () {
            const { marketplace, anotherUser, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            const [balance1] = await marketplace.getBalance(anotherUser.address)
            expect(balance1).is.equal(0)

            const mintId1 = randomBytes12()
            await marketplace.connect(kycSigner).mint(anotherUser.address, '123456', mintId1)

            const [balance2] = await marketplace.getBalance(anotherUser.address)
            expect(balance2).is.equal(123456)

            const mintId2 = randomBytes12()
            await marketplace.connect(kycSigner).mint(anotherUser.address, '111111', mintId2)

            const [balance3] = await marketplace.getBalance(anotherUser.address)
            expect(balance3).is.equal(123456 + 111111)
        })
        it("should prevent double transaction", async function () {
            const { marketplace, anotherUser, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            const [balance1] = await marketplace.getBalance(anotherUser.address)
            expect(balance1).is.equal(0)

            const mintId1 = randomBytes12()
            await marketplace.connect(kycSigner).mint(anotherUser.address, '123456', mintId1)

            const [balance2] = await marketplace.getBalance(anotherUser.address)
            expect(balance2).is.equal(123456)

            await expect(marketplace.connect(kycSigner).mint(anotherUser.address, '111111', mintId1))
                .to.be.revertedWithCustomError(marketplace, "MintBurnIdReused")
                .withArgs(mintId1)
        })
        it("should check user KYC", async function () {
            const { marketplace, anotherUser, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            const mintId1 = randomBytes12()
            await expect(marketplace.connect(kycSigner).mint(anotherUser.address, '111111', mintId1))
                .to.be.revertedWithCustomError(marketplace, "KYCProblem")
                .withArgs(anotherUser.address, "0x0000")

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            await marketplace.connect(kycSigner).mint(anotherUser.address, '111111', mintId1)
        })
        it("should reject if not a KYCSigner", async function () {
            const { marketplace, admin, anotherUser, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            const mintId1 = randomBytes12()
            await expect(marketplace.connect(admin).mint(anotherUser.address, '111111', mintId1))
                .to.be.revertedWithCustomError(marketplace, "NotAuthorized")

            await marketplace.connect(kycSigner).mint(anotherUser.address, '111111', mintId1)
        })
        it("should transfer value to recipient", async function () {
            const { marketplace, admin, anotherUser, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            await marketplace.connect(admin).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            //check creditToken balance 
            expect(await anotherUser.getBalance()).to.equal(ethers.utils.parseEther("10000"))
            const initialMinerBalance = await kycSigner.getBalance()

            const mintId = randomBytes12()
            const tx = await marketplace.connect(kycSigner).mint(anotherUser.address, 7777, mintId, {
                value: ethers.utils.parseEther("123")
            })

            const gasUsed = await tx.wait().then((receipt) => receipt.gasUsed)
            const gasPrice = await tx.wait().then((receipt) => receipt.effectiveGasPrice)

            const expectedGasCost = gasPrice.mul(gasUsed)

            expect(await marketplace.getFreeBalance(anotherUser.address)).to.equal(7777)
            expect(await anotherUser.getBalance()).to.equal(ethers.utils.parseEther("10123"))

            const finalMinerBalance = await kycSigner.getBalance()
            expect(initialMinerBalance.sub(finalMinerBalance).sub(expectedGasCost)).to.equal(ethers.utils.parseEther("123"))
        })
        it("should revert if coin is NOT empty", async function () {
            const { marketplace, anotherUser, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            await expect(marketplace.connect(kycSigner).mint(anotherUser.address, '111111', randomBytes12()))
                .to.be.revertedWithCustomError(marketplace, "DirectMintBurnNotAllowed")
        })
    })
    describe("burn", () => {
        it("should decrease balance", async function () {
            const { marketplace, user, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            const [balance1] = await marketplace.getBalance(user.address)
            expect(balance1).is.equal(DEFAULT_USER_BALANCE)

            const burnId1 = randomBytes12()
            await marketplace.connect(kycSigner).burn(user.address, '123456', burnId1)

            const [balance2] = await marketplace.getBalance(user.address)
            expect(balance2).is.equal(DEFAULT_USER_BALANCE - 123456)
        })
        it("should prevent re-burning", async function () {
            const { marketplace, user, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            const mintId1 = randomBytes12()
            await marketplace.connect(kycSigner).burn(user.address, '123456', mintId1)

            await expect(marketplace.connect(kycSigner).burn(user.address, '111111', mintId1))
                .to.be.revertedWithCustomError(marketplace, "MintBurnIdReused")
                .withArgs(mintId1)
        })
        it("should reject if not a KYCSigner", async function () {
            const { marketplace, user, kycSigner } = await loadFixture(deployCreditMarketplaceV3Fixture);

            const mintId1 = randomBytes12()
            await expect(marketplace.connect(user).burn(user.address, '111111', mintId1))
                .to.be.revertedWithCustomError(marketplace, "NotAuthorized")

            await marketplace.connect(kycSigner).burn(user.address, '111111', mintId1)
        })
        it("should revert if coin is NOT empty", async function () {
            const { marketplace, anotherUser, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await expect(marketplace.connect(kycSigner).burn(anotherUser.address, '111111', randomBytes12()))
                .to.be.revertedWithCustomError(marketplace, "DirectMintBurnNotAllowed")
        })
    });
});
