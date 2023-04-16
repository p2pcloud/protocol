// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./BalanceHolder.sol";

abstract contract ProviderRegistry is BalanceHolder {
    struct ProviderInfo {
        //slot 1
        bytes32 url;
        //slot 2
        bool isRegistered;
        uint128 feePaid;
    }

    mapping(address => ProviderInfo) providerInfo;

    function setProviderUrl(bytes32 url) public {
        require(
            providerInfo[msg.sender].isRegistered,
            "Provider must be registered to set url"
        );
        providerInfo[msg.sender].url = url;
    }

    function getProviderUrl(address _user) public view returns (bytes32) {
        return providerInfo[_user].url;
    }

    function isProviderRegistered(address _user) public view returns (bool) {
        return providerInfo[_user].isRegistered;
    }

    uint64 public constant PROVIDER_REGISTRATION_FEE = 100 * 1000000;

    function registerProvider() public {
        require(
            !providerInfo[msg.sender].isRegistered,
            "Provider is already registered"
        );

        require(
            _isSpendable(msg.sender, PROVIDER_REGISTRATION_FEE),
            "Not enough coin to register "
        );

        _spendWithComission(msg.sender, owner(), PROVIDER_REGISTRATION_FEE);

        providerInfo[msg.sender].isRegistered = true;
        providerInfo[msg.sender].feePaid += PROVIDER_REGISTRATION_FEE;
    }
}
