// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "./ERC20/ERC20.sol";

contract Token is ERC20 {
    constructor(uint256 initialSupply) public ERC20("USD stablecoin", "USDC") {
        _mint(msg.sender, initialSupply);
    }
}