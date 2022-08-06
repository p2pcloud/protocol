// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "./ERC20/ERC20.sol";

contract Token is ERC20 {
    constructor() ERC20("DEBUG stablecoin", "DEBUG") {
        _mint(msg.sender, 1000000 * 9000); //9000 coins with 6 percision
    }

    function decimals() public view virtual override returns (uint8) {
        return 6;
    }
}
