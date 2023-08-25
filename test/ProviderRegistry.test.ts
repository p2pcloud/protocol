import { expect } from "chai";
import { ethers } from "hardhat";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { deployMarketplaceV3Fixture } from './fixtures'
import { NZ_HEX, US_HEX, signKYC } from "./lib";

describe("ProviderRegistry", function () {
    describe("getProviderUrl", function () {
        it("should set provider url", async function () {
            const { marketplace, anotherUser, admin, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            const sig = await signKYC(anotherUser.address, US_HEX, kycSigner)
            await marketplace.connect(anotherUser).submitKYC(anotherUser.address, US_HEX, sig)


            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)

            const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)
            const url = await marketplace.connect(anotherUser).getProviderUrl(anotherUser.address);
            expect(url).to.equal(urlBytes);
        });

        it("should not set provider url for non-registered provider", async function () {
            const { marketplace, anotherUser, admin, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            const sig = await signKYC(anotherUser.address, US_HEX, kycSigner)
            await marketplace.connect(anotherUser).submitKYC(anotherUser.address, US_HEX, sig)

            const urlBytes = ethers.utils.formatBytes32String("woop.woop/woop");
            expect(marketplace.connect(anotherUser).setProviderUrl(urlBytes)).to.be.revertedWith("Provider must be registered to set url")

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)

            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)
        });
    })
    describe("getAllProviderURLs", function () {
        it("should return all provider urls", async function () {
            const { marketplace, anotherUser, admin, provider, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            await marketplace.connect(admin).submitKYC(
                admin.address, US_HEX,
                await signKYC(admin.address, US_HEX, kycSigner)
            )

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)


            const urlBytes = ethers.utils.formatBytes32String("another.example.com");
            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)

            // Register admin as another provider
            await marketplace.connect(admin).registerProviderByCommunity(admin.address)

            const anotherUrlBytes = ethers.utils.formatBytes32String("admin.com");
            await marketplace.connect(admin).setProviderUrl(anotherUrlBytes)

            // Get all provider URLs
            const [addresses, urls] = await marketplace.connect(anotherUser).getAllProviderURLs();

            expect(addresses).to.deep.equal([provider.address, anotherUser.address, admin.address]);
            expect(urls.map(ethers.utils.parseBytes32String))
                .to.deep.equal(["", "another.example.com", "admin.com"]);
        });
        it("should ignore deleted providers", async function () {
            const { marketplace, anotherUser, admin, provider, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);


            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            await marketplace.connect(admin).submitKYC(
                admin.address, US_HEX,
                await signKYC(admin.address, US_HEX, kycSigner)
            )

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)


            const urlBytes = ethers.utils.formatBytes32String("another.example.com");
            await marketplace.connect(anotherUser).setProviderUrl(urlBytes)

            // Register admin as another provider
            await marketplace.connect(admin).registerProviderByCommunity(admin.address)

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
            const { marketplace, anotherUser, admin, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            const sig = await signKYC(anotherUser.address, US_HEX, kycSigner)
            await marketplace.connect(anotherUser).submitKYC(anotherUser.address, US_HEX, sig)

            const isregistered1 = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isregistered1).to.equal(false)

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)

            const isregistered2 = await marketplace.connect(anotherUser).isProviderRegistered(anotherUser.address)
            expect(isregistered2).to.equal(true)
        })
    })

    describe("registerProvider", function () {
        it("should require KYC", async function () {
            const { marketplace, anotherUser, admin, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await expect(
                marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)
            ).to.be.rejectedWith("No KYC or country is not allowed")

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)

        })

        it("should enforce provider's country", async function () {
            const { marketplace, anotherUser, admin, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);
            await marketplace.connect(admin).allowUserCountry(NZ_HEX)


            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, NZ_HEX,
                await signKYC(anotherUser.address, NZ_HEX, kycSigner)
            )

            await expect(
                marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)
            ).to.be.rejectedWith("No KYC or country is not allowed")

            await marketplace.connect(admin).allowProviderCountry(NZ_HEX)

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)
        })
        it("should be called only by admin", async function () {
            const { marketplace, provider, anotherUser, admin, kycSigner } = await loadFixture(deployMarketplaceV3Fixture);

            await marketplace.connect(anotherUser).submitKYC(
                anotherUser.address, US_HEX,
                await signKYC(anotherUser.address, US_HEX, kycSigner)
            )

            await expect(
                marketplace.connect(provider).registerProviderByCommunity(anotherUser.address)
            ).to.be.rejectedWith("Ownable: caller is not the owner")

            await expect(
                marketplace.connect(anotherUser).registerProviderByCommunity(anotherUser.address)
            ).to.be.rejectedWith("Ownable: caller is not the owner")

            await marketplace.connect(admin).registerProviderByCommunity(anotherUser.address)
        })
    });
    describe("setProviderRegistrationFee", function () {
        it("change set provider fee for new providers only")
    })

    describe("deleteProvider", function () {
        it("can be called by provdier", async () => {
            const { marketplace, provider } = await loadFixture(deployMarketplaceV3Fixture);

            let isRegistered
            isRegistered = await marketplace.isProviderRegistered(provider.address)
            expect(isRegistered).to.equal(true)

            await marketplace.connect(provider).deleteProvider(provider.address)
            isRegistered = await marketplace.isProviderRegistered(provider.address)
            expect(isRegistered).to.equal(false)
        })
        it("can be called by community", async () => {
            const { marketplace, provider, admin, user } = await loadFixture(deployMarketplaceV3Fixture);

            await expect(marketplace.connect(user).deleteProvider(provider.address)).to.be.revertedWith("Provider or community only")
            let isRegistered
            isRegistered = await marketplace.isProviderRegistered(provider.address)
            expect(isRegistered).to.equal(true)

            await marketplace.connect(admin).deleteProvider(provider.address)
            isRegistered = await marketplace.isProviderRegistered(provider.address)
            expect(isRegistered).to.equal(false)
        })
    })
});
