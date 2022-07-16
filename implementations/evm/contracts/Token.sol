// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "./ERC20/ERC20.sol";

contract Token is ERC20 {
    constructor(uint256 initialSupply) public ERC20("USD stablecoin", "USDC") {
        _mint(msg.sender, initialSupply);
    }

    function transferForTests(address from, address to, uint256 amount) public returns (bool) {
        _transfer(from, to, amount);
        return true;
    }
}