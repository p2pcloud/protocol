// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

abstract contract BalanceHolder is OwnableUpgradeable {
    uint16 public communityFee;

    mapping(address => uint256) internal coinBalance;
    IERC20 public coin;

    function depositCoin(uint256 numTokens) public {
        require(coin.transferFrom(msg.sender, address(this), numTokens), "Failed to transfer tokens");

        coinBalance[msg.sender] = coinBalance[msg.sender] + numTokens;
    }

    function withdrawCoin(uint256 amt) public {
        require(getFreeBalance(msg.sender) >= amt, "Not enough balance to withdraw");
        coinBalance[msg.sender] -= amt;
        require(coin.transfer(msg.sender, amt), "ERC20 transfer failed");
    }

    function _spendWithComission(address spender, address receiver, uint256 amt) internal returns (bool defaulted) {
        uint256 spentAmt = _min(coinBalance[spender], amt);

        uint256 communityPayout = (spentAmt * communityFee) / (100 * 100);
        uint256 providerPayout = spentAmt - communityPayout;

        coinBalance[spender] -= spentAmt;
        coinBalance[receiver] += providerPayout;
        coinBalance[owner()] += communityPayout;

        return spentAmt != amt;
    }

    function getFreeBalance(address user) public view returns (uint256) {
        return coinBalance[user] - getLockedBalance(user);
    }

    function getLockedBalance(address user) public view virtual returns (uint256);

    function getBalance(address user) public view returns (uint256 free, uint256 locked) {
        return (getFreeBalance(user), getLockedBalance(user));
    }

    function _min(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a < b) {
            return a;
        }
        return b;
    }
}
