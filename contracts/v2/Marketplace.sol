// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";
import "./Broker.sol";

contract MarketplaceV2 is BrokerV2 {
    function initialize(IERC20 _coin) public initializer {
        _transferOwnership(msg.sender);

        __VerifiableOffer_init();
        coin = _coin;
    }
}
