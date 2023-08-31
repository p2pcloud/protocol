// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockERC20V2 is ERC20 {
    bool public transferWillFail = false;

    constructor(uint256 initialSupply) ERC20("Gold", "GLD") {
        _mint(msg.sender, initialSupply);
    }

    function transferFrom(address from, address to, uint256 amount) public virtual override returns (bool) {
        require(!transferWillFail, "Transfer from failed becuase test wanted it");

        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        _transfer(from, to, amount);
        return true;
    }

    function transfer(address to, uint256 amount) public virtual override returns (bool) {
        require(!transferWillFail, "Transfer from failed becuase test wanted it");

        address owner = _msgSender();
        _transfer(owner, to, amount);
        return true;
    }

    function transferShouldFail(bool _shouldIt) public virtual {
        transferWillFail = _shouldIt;
    }
}
