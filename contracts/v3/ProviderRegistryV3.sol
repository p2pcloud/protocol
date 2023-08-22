// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./VerifiableKYCV3.sol";

abstract contract ProviderRegistryV3 is VerifiableKYCV3 {
    function setProviderUrl(bytes32 url) public {
        require(providerInfo[msg.sender].isRegistered, "Provider must be registered to set url");
        providerInfo[msg.sender].url = url;
    }

    function getProviderUrl(address _user) public view returns (bytes32) {
        return providerInfo[_user].url;
    }

    function isProviderRegistered(address _user) public view returns (bool) {
        return providerInfo[_user].isRegistered;
    }

    // function registerProvider() public virtual {
    //     require(providerRegistrationEnabled, "Provider registration is disabled");
    //     _registerProvider(msg.sender, PROVIDER_REGISTRATION_FEE);
    // }

    function registerProviderByCommunity(address _provider) public virtual onlyOwner {
        require(checkProviderKYC(_provider), "No KYC or country is not allowed");
        providerInfo[_provider].isRegistered = true;
        providerList.push(_provider);
    }

    // function _registerProvider(address provider, uint64 fee) internal {
    //     require(checkProviderKYC(provider), "No KYC or country is not allowed");
    //     require(!providerInfo[provider].isRegistered, "Provider is already registered");

    //     require(getFreeBalance(provider) >= fee, "Not enough coin to register ");

    //     _spendWithComission(provider, owner(), fee);

    //     providerInfo[provider].isRegistered = true;
    //     providerInfo[provider].feePaid += fee;

    //     providerList.push(provider);
    // }

    function getAllProviderURLs() public view returns (address[] memory, bytes32[] memory) {
        uint256 providerCount = 0;
        for (uint i = 0; i < providerList.length; i++) {
            if (providerInfo[providerList[i]].isRegistered) {
                providerCount++;
            }
        }
        address[] memory addresses = new address[](providerCount);
        bytes32[] memory urls = new bytes32[](providerCount);
        uint j = 0;
        for (uint256 i = 0; i < providerList.length; i++) {
            if (!providerInfo[providerList[i]].isRegistered) {
                continue;
            }
            addresses[j] = providerList[i];
            urls[j] = providerInfo[providerList[i]].url;
            j++;
        }
        return (addresses, urls);
    }

    function deleteProvider(address _user) public {
        require(_user == msg.sender || owner() == msg.sender, "Provider or community only");
        providerInfo[_user].isRegistered = false;
    }
}
