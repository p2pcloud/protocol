// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "./IIdentityProvider.sol";

contract TestableIdentityProvider is IIdentityProvider {
    struct TestPayload {
        bytes2 country;
        bytes3 subdivision;
        bool verified;
    }

    mapping(address => TestPayload) testData;

    function test__injectVerification(address user, bytes2 country, bytes3 subdivision) external {
        testData[user] = TestPayload(country, subdivision, true);
    }

    function getRegion(address _customer) external view override returns (bool, bytes2, bytes3) {
        return (testData[_customer].verified, testData[_customer].country, testData[_customer].subdivision);
    }

    function getVerificationUrl() external pure override returns (string memory) {
        return "whatever";
    }

    function getCountry(address _customer) external view override returns (bool, bytes2) {
        return (testData[_customer].verified, testData[_customer].country);
    }
}
