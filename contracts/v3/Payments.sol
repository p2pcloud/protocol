// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";

abstract contract PaymentsV3 is BalanceHolderV3 {
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
