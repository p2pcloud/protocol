// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

abstract contract CommunityOwnable is OwnableUpgradeable {
    uint16 private _communityFee;

    function communityFee() public view returns (uint16) {
        return _communityFee;
    }

    function setCommunityFee(uint16 fee) public onlyOwner returns (bool) {
        require(
            fee < 2500,
            "community fee should be in range of 0 (0%) to 2500 (25%)"
        );

        _communityFee = fee;
        return false;
    }
}
