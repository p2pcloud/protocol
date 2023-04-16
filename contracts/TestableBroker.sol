// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./Broker.sol";

contract TestableBroker is Broker {
    function test__getTime() public view returns (uint256) {
        return block.timestamp;
    }

    function test__increaseLockedBalance(address user, uint256 amt) public {
        _lockedBalance[user] += uint32(amt);
    }
}
