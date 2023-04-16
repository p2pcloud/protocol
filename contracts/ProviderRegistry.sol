// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./BalanceHolder.sol";

contract ProviderRegistry is BalanceHolder{
    struct ProviderInfo {
        //slot 1
        bytes32 url;
        //slot 2
        bool isRegistered;
        uint128 feePaid;
    }

    mapping(address => ProviderInfo) providerInfo;


    function SetProviderUrl(bytes32 url) public {
        require(
            providerInfo[msg.sender].isRegistered,
            "Provider must be registered to set url"
        );
        providerInfo[msg.sender].url = url;
    }

    function GetProviderUrl(address _user) public view returns (bytes32) {
        return providerInfo[_user].url;
    }

    function IsProviderRegistered(address _user) public view returns (bool) {
        return providerInfo[_user].isRegistered;
    }


    uint64 public constant PROVIDER_REGISTRATION_FEE = 100 * 1000000;

    function RegisterProvider() public {
        uint256 freeBalance = coinBalance[msg.sender] -
            lockedBalance[msg.sender];
        require(
            freeBalance >= PROVIDER_REGISTRATION_FEE,
            "Not enough coin to register "
        );
        require(
            !providerInfo[msg.sender].isRegistered,
            "Provider is already registered"
        );

        coinBalance[msg.sender] -= PROVIDER_REGISTRATION_FEE;
        coinBalance[owner()] += PROVIDER_REGISTRATION_FEE;

        providerInfo[msg.sender].isRegistered = true;
        providerInfo[msg.sender].feePaid += PROVIDER_REGISTRATION_FEE;
    }
}