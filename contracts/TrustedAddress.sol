// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

contract TrustedAddress {
    mapping(address => address) trustedAddresses;

    function setTrustedAddress(address trusted) public {
        trustedAddresses[msg.sender] = trusted;
    }

    function getTrustedAddress(address user) public view returns (address) {
        return trustedAddresses[user];
    }
}
