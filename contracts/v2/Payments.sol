// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";

abstract contract PaymentsV2 is BalanceHolderV2 {
    uint256 public constant MONEY_LOCK_MINUTES = 60 * 24 * 7; // 7 days

    struct UserProviderAccounting {
        uint128 lastPaymentTs;
        uint128 pricePerMinute;
    }

    mapping(address => mapping(address => UserProviderAccounting)) public userProviderAccounting;
    mapping(address => uint256) internal _totalSpendingPerMinute;

    function getLockedBalance(address user) public view override returns (uint256) {
        return _totalSpendingPerMinute[user] * MONEY_LOCK_MINUTES;
    }

    function _executeClaimPayment(address provider, address client) internal {
        uint256 pricePerMinute = userProviderAccounting[provider][client].pricePerMinute;

        if (pricePerMinute == 0) {
            //initialize account
            userProviderAccounting[provider][client].lastPaymentTs = uint128(block.timestamp);
            return;
        }

        uint256 usedTime = (block.timestamp - userProviderAccounting[provider][client].lastPaymentTs);
        uint256 minutesPassed;
        assembly {
            minutesPassed := div(usedTime, 60)
        }
        uint256 amount = minutesPassed * pricePerMinute;

        _spendWithComission(client, provider, amount);
        userProviderAccounting[provider][client].lastPaymentTs += uint128(minutesPassed * 60);
    }

    function claimPayment(address client) external {
        _executeClaimPayment(msg.sender, client);
    }
}
