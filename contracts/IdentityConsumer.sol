// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./TestableIdentityProvider.sol";

abstract contract IdentityConsumer is OwnableUpgradeable {
    IIdentityProvider identityProvider;

    bytes2[] allowedProviderCountries;

    function setProviderCountries(bytes2[] memory allowedCountries) public onlyOwner {
        allowedProviderCountries = allowedCountries;
    }

    function isProvidersCountryValid(address provider) internal view returns (bool) {
        (bool verified, bytes2 country) = identityProvider.getCountry(provider);
        return verified && _isAllowedCountry(country);
    }

    function _isAllowedCountry(bytes2 country) internal view returns (bool) {
        for (uint256 i = 0; i < allowedProviderCountries.length; i++) {
            if (allowedProviderCountries[i] == country) {
                return true;
            }
        }
        return false;
    }
}
