// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./BalanceHolder.sol";
import "./ProviderRegistry.sol";
import "./Broker.sol";

contract Marketplace is BalanceHolder, ProviderRegistry, Broker {
    function initialize() public initializer {
        _transferOwnership(msg.sender);

        __VerifiableOffer_init();
        communityFee = 2000;
    }
}
