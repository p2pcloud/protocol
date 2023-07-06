import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceFixture } from './fixtures'

describe("ProviderRegistry", function () {
    describe("getProviderUrl", function () {
        it("should set provider url", async function () {
            const { marketplace, token, anotherUser, admin, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])

            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(anotherUser).depositCoin(fee)
            await marketplace.connect(anotherUser).registerProvider()

            const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)
            const url = await marketplace.connect(anotherUser).getProviderUrl(anotherUser.address);
            expect(url).to.equal(urlBytes);
        });

        it("should not set provider url for non-registered provider", async function () {
            const { marketplace, token, anotherUser, admin, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);


            const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
            expect(marketplace.connect(anotherUser).setProviderUrl(urlBytes)).to.be.revertedWith("Provider must be registered to set url")

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])

            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(anotherUser).depositCoin(fee)
            await marketplace.connect(anotherUser).registerProvider()

            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)
        });
    })
    describe("getAllProviderURLs", function () {
        it("should return all provider urls", async function () {
            const { marketplace, token, anotherUser, admin, provider, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])
            await testableIdentityProvider.connect(admin).test__injectVerification(admin.address, [1, 1], [0, 0, 0])

            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(anotherUser).depositCoin(fee)
            await marketplace.connect(anotherUser).registerProvider()

            const urlBytes = ethers.utils.formatBytes32String("another.example.com");
            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)

            // Register admin as another provider
            await token.connect(admin).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(admin).depositCoin(fee)
            await marketplace.connect(admin).registerProvider()

            const anotherUrlBytes = ethers.utils.formatBytes32String("admin.com");
            await marketplace.connect(admin).setProviderUrl(anotherUrlBytes)

            // Get all provider URLs
            const [addresses, urls] = await marketplace.connect(anotherUser).getAllProviderURLs();

            expect(addresses).to.deep.equal([provider.address, anotherUser.address, admin.address]);
            expect(urls.map(ethers.utils.parseBytes32String))
                .to.deep.equal(["", "another.example.com", "admin.com"]);
        });
        it("should ignore deleted providers", async function () {
            const { marketplace, token, anotherUser, admin, provider, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])
            await testableIdentityProvider.connect(admin).test__injectVerification(admin.address, [1, 1], [0, 0, 0])


            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(anotherUser).depositCoin(fee)
            await marketplace.connect(anotherUser).registerProvider()

            const urlBytes = ethers.utils.formatBytes32String("another.example.com");
            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)

            // Register admin as another provider
            await token.connect(admin).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(admin).depositCoin(fee)
            await marketplace.connect(admin).registerProvider()

            const anotherUrlBytes = ethers.utils.formatBytes32String("admin.com");
            await marketplace.connect(admin).setProviderUrl(anotherUrlBytes)

            // Get all provider URLs
            const [addresses, urls] = await marketplace.getAllProviderURLs();
            expect(addresses).to.deep.equal([provider.address, anotherUser.address, admin.address]);

            //delete one provider
            await marketplace.connect(provider).deleteProvider(provider.address)

            const [addresses2, urls2] = await marketplace.getAllProviderURLs();
            expect(addresses2).to.deep.equal([anotherUser.address, admin.address]);

            //delete one provider
            await marketplace.connect(admin).deleteProvider(admin.address)

            const [addresses3, urls3] = await marketplace.getAllProviderURLs();
            expect(addresses3).to.deep.equal([anotherUser.address]);

        });
    });


    describe("isProviderRegistered", function () {
        it("should return true for registered and false by default", async function () {
            const { marketplace, token, anotherUser, admin, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])

            const isregistered1 = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isregistered1).to.equal(false)

            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee)
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee)
            await marketplace.connect(anotherUser).depositCoin(fee)
            await marketplace.connect(anotherUser).registerProvider()

            const isregistered2 = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isregistered2).to.equal(true)
        })
    })


    describe("registerProvider", function () {
        it("should use exactly PROVIDER_REGISTRATION_FEE coins", async function () {
            const { marketplace, token, anotherUser, admin, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])

            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee.add(123))
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee.add(123))
            await marketplace.connect(anotherUser).depositCoin(fee.add(123))


            const balance1 = await marketplace.connect(anotherUser).getBalance(anotherUser.address)
            expect(balance1[0]).to.equal(fee.add(123))
            expect(balance1[1]).to.equal(0)

            await marketplace.connect(anotherUser).registerProvider()

            const balance2 = await marketplace.connect(anotherUser).getBalance(anotherUser.address)
            expect(balance2[0]).to.equal(123)
            expect(balance2[1]).to.equal(0)
        });

        it("should reject if not enough coins", async function () {
            const { marketplace, token, anotherUser, admin } = await loadFixture(deployMarketplaceFixture);

            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee.sub(1))
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee.sub(1))
            await marketplace.connect(anotherUser).depositCoin(fee.sub(1))

            await expect(marketplace.connect(anotherUser).registerProvider()).to.be.rejected
        });
    });
    describe("setProviderRegistrationFee", function () {
        it("change set provider fee for new providers only")
    })

    describe("deleteProvider", function () {
        it("can be called by provdier", async () => {
            const { marketplace, token, anotherUser, admin, user, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])


            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee.add(123))
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee.add(123))
            await marketplace.connect(anotherUser).depositCoin(fee.add(123))

            await marketplace.connect(anotherUser).registerProvider()

            await expect(marketplace.connect(user).deleteProvider(anotherUser.address)).to.be.revertedWith("Provider or community only")
            let isRegistered
            isRegistered = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isRegistered).to.equal(true)


            await marketplace.connect(anotherUser).deleteProvider(anotherUser.address)
            isRegistered = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isRegistered).to.equal(false)

        })
        it("can be called by community", async () => {
            const { marketplace, token, anotherUser, admin, user, testableIdentityProvider } = await loadFixture(deployMarketplaceFixture);

            await testableIdentityProvider.connect(admin).test__injectVerification(anotherUser.address, [1, 1], [0, 0, 0])


            const fee = await marketplace.PROVIDER_REGISTRATION_FEE()
            await token.connect(admin).transfer(anotherUser.address, fee.add(123))
            await token.connect(anotherUser).increaseAllowance(marketplace.address, fee.add(123))
            await marketplace.connect(anotherUser).depositCoin(fee.add(123))

            await marketplace.connect(anotherUser).registerProvider()

            await expect(marketplace.connect(user).deleteProvider(anotherUser.address)).to.be.revertedWith("Provider or community only")
            let isRegistered
            isRegistered = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isRegistered).to.equal(true)


            await marketplace.connect(admin).deleteProvider(anotherUser.address)
            isRegistered = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isRegistered).to.equal(false)
        })
    })
});
