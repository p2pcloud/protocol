// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract CommunityOwnable is Ownable {
    uint16 public communityFee;

    function SetCommunityFee(uint16 fee) public onlyOwner returns (bool) {
        require(
            fee < 2500,
            "community fee should be in range of 0 (0%) to 2500 (25%)"
        );

        communityFee = fee;
        return false;
    }
}
