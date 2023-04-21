// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./Marketplace.sol";

contract TestableMarketplace is Marketplace {
    function test__getTime() public view returns (uint256) {
        return block.timestamp;
    }

    function test__increaseSpendingPerMinute(address user, uint256 amt) public {
        _totalSpendingPerMinute[user] += amt;
    }
}
