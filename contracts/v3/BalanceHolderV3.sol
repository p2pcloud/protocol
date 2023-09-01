// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./VerifiableKYCV3.sol";

error DirectMintBurnNotAllowed();
error AlreadyMinted(bytes12 mintId);

abstract contract BalanceHolderV3 is VerifiableKYCV3 {
    function depositCoin(uint256 numTokens) public virtual {
        checkUserKYC(msg.sender);
        if (!coin.transferFrom(msg.sender, address(this), numTokens)) {
            revert ERC20TransferFailed();
        }

        _coinBalance[msg.sender] = _coinBalance[msg.sender] + numTokens;
    }

    function withdrawCoin(uint256 amt) public virtual {
        checkUserKYC(msg.sender);
        if (getFreeBalance(msg.sender) < amt) {
            revert InsufficientBalance(amt, getFreeBalance(msg.sender));
        }
        _coinBalance[msg.sender] -= amt;
        if (!coin.transfer(msg.sender, amt)) {
            revert ERC20TransferFailed();
        }
    }

    function _spendWithComission(address spender, address receiver, uint256 amt) internal returns (bool defaulted) {
        uint256 spentAmt = _min(_coinBalance[spender], amt);

        uint256 communityPayout = (spentAmt * COMMUNITY_FEE) / (100 * 100);
        uint256 providerPayout = spentAmt - communityPayout;

        _coinBalance[spender] -= spentAmt;
        _coinBalance[receiver] += providerPayout;
        _coinBalance[owner()] += communityPayout;

        return spentAmt != amt;
    }

    function getFreeBalance(address user) public view returns (uint256) {
        uint256 locked = getLockedBalance(user);
        if (_coinBalance[user] < locked) {
            return 0;
        } else {
            return _coinBalance[user] - locked;
        }
    }

    function getTotalBalance(address user) public view returns (uint256) {
        return _coinBalance[user];
    }

    function getLockedBalance(address user) public view virtual returns (uint256);

    function getBalance(address user) public view returns (uint256 free, uint256 locked) {
        return (getFreeBalance(user), getLockedBalance(user));
    }

    function _min(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a <= b) {
            return a;
        }
        return b;
    }

    function mint(address payable to, uint256 amount, bytes12 mintId) public payable onlyOwner {
        checkUserKYC(to);

        if (address(coin) != address(0)) {
            revert DirectMintBurnNotAllowed();
        }

        if (mintBurnIdempotency[mintId]) {
            revert AlreadyMinted(mintId);
        }

        mintBurnIdempotency[mintId] = true;

        _coinBalance[to] = _coinBalance[to] + amount;

        // Transfer any ether sent with the transaction to the 'to' address.
        if (msg.value > 0) {
            to.transfer(msg.value);
        }
    }

    function burn(address payable to, uint256 amount, bytes12 mintId) public onlyOwner {
        if (address(coin) != address(0)) {
            revert DirectMintBurnNotAllowed();
        }

        if (mintBurnIdempotency[mintId]) {
            revert AlreadyMinted(mintId);
        }

        mintBurnIdempotency[mintId] = true;

        _coinBalance[to] = _coinBalance[to] - amount;
    }
}
