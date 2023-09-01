// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";

abstract contract ProviderRegistryV2 is BalanceHolderV2 {
    uint64 public constant PROVIDER_REGISTRATION_FEE = 100 * 1000000;

    struct ProviderInfo {
        //slot 1
        bytes32 url;
        //slot 2
        bool isRegistered;
        uint128 feePaid;
    }

    mapping(address => ProviderInfo) providerInfo;
    address[] public providerList;

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

    function registerProvider() public virtual {
        _registerProvider(msg.sender, PROVIDER_REGISTRATION_FEE);
    }

    function _registerProvider(address provider, uint64 fee) internal {
        require(!providerInfo[provider].isRegistered, "Provider is already registered");

        require(getFreeBalance(provider) >= fee, "Not enough coin to register ");

        _spendWithComission(provider, owner(), fee);

        providerInfo[provider].isRegistered = true;
        providerInfo[provider].feePaid += fee;

        providerList.push(provider);
    }

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