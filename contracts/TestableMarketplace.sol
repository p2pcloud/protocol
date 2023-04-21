// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./Marketplace.sol";

contract TestableMarketplace is Marketplace {
    function test__getTime() public view returns (uint256) {
        return block.timestamp;
    }

    function test__increaseSpendingPerMinute(address user, uint256 amt) public {
        _totalSpendingPerMinute[user] += amt;
    }
}
