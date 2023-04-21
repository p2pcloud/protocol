// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

import "./BalanceHolder.sol";
import "./Broker.sol";

contract Marketplace is BalanceHolder, Broker {
    function initialize(IERC20 _coin) public initializer {
        _transferOwnership(msg.sender);

        __VerifiableOffer_init();
        communityFee = 2000;
        coin = _coin;
    }
}
