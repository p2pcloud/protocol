// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

abstract contract Testable {
    function GetTime() public view returns (uint256) {
        //TODO: remove. Test function only
        return block.timestamp;
    }
}
