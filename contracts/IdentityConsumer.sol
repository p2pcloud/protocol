// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./TestableIdentityProvider.sol";

abstract contract IdentityConsumer is OwnableUpgradeable {
    IIdentityProvider identityProvider;

    bytes2[] allowedProviderCountries;
    mapping(address => bytes2[]) allowedClientCountries;

    function setProviderCountries(bytes2[] memory allowedCountries) public onlyOwner {
        allowedProviderCountries = allowedCountries;
    }

    function checkProvidersIdentity(address provider) internal view returns (bool) {
        (bool verified, bytes2 country) = identityProvider.getCountry(provider);
        return verified && _isAllowedProviderCountry(country);
    }

    function checkClientsIdentity(address client, address provider) internal view returns (bool) {
        (bool verified, bytes2 country) = identityProvider.getCountry(client);
        return verified && _isAllowedClientCountry(country, provider);
    }

    function _isAllowedProviderCountry(bytes2 country) internal view returns (bool) {
        return _isAllowedCountry(country, allowedProviderCountries);
    }

    function _isAllowedClientCountry(bytes2 country, address provider) internal view returns (bool) {
        return _isAllowedCountry(country, allowedClientCountries[provider]);
    }

    function _isAllowedCountry(bytes2 country, bytes2[] memory allowList) internal pure returns (bool) {
        for (uint256 i = 0; i < allowList.length; i++) {
            if (allowList[i] == country) {
                return true;
            }
        }
        return false;
    }

    function setClientCountries(bytes2[] memory allowedCountries) public {
        allowedClientCountries[msg.sender] = allowedCountries;
    }
}
