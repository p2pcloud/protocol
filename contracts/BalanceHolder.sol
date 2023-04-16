// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./CommunityOwnable.sol";

abstract contract BalanceHolder is CommunityOwnable {
    mapping(address => uint256) coinBalance;
    mapping(address => uint32) lockedBalance;
    mapping(address => uint256) nonWithdrawableBalance;
    IERC20 public coin;

    function SetCoinAddress(IERC20 newCoinAddress) public onlyOwner {
        coin = newCoinAddress;
    }

    function DepositCoin(uint256 numTokens) public {
        require(
            coin.transferFrom(msg.sender, address(this), numTokens),
            "Failed to transfer tokens"
        );

        coinBalance[msg.sender] = coinBalance[msg.sender] + numTokens;
    }

    function WithdrawCoin(uint256 amt) public {
        uint256 freeBalance = coinBalance[msg.sender] -
            lockedBalance[msg.sender] -
            nonWithdrawableBalance[msg.sender];

        require(freeBalance >= amt, "Not enough balance to withdraw");

        coinBalance[msg.sender] -= amt;

        require(coin.transfer(msg.sender, amt), "ERC20 transfer failed");
    }

    function GetCoinBalance(
        address user
    )
        public
        view
        returns (uint256 free, uint256 locked, uint256 nonWithdrawable)
    {
        locked = lockedBalance[user];
        free = coinBalance[user] - locked;
        nonWithdrawable = nonWithdrawableBalance[user];
        return (free, locked, nonWithdrawable);
    }

    function NonWithdrawableTransfer(uint256 amt, address to) public {
        uint256 freeBalance = coinBalance[msg.sender] -
            lockedBalance[msg.sender] -
            nonWithdrawableBalance[msg.sender];

        require(freeBalance >= amt, "Not enough balance to transfer");

        coinBalance[msg.sender] -= amt;

        coinBalance[to] += amt;
        nonWithdrawableBalance[to] += amt;
    }

    function _spend(
        address user,
        uint256 amt
    ) internal returns (uint256, bool) {
        uint256 spentAmt = _min(coinBalance[user], amt);

        coinBalance[user] -= spentAmt;
        nonWithdrawableBalance[user] -= _min(nonWithdrawableBalance[user], amt);
        return (spentAmt, spentAmt == amt);
    }

    function _isSpendable(
        address user,
        uint256 amt
    ) internal view returns (bool) {
        uint256 freeBalance = coinBalance[user] - lockedBalance[user];
        return freeBalance >= amt;
    }

    function _min(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a < b) {
            return a;
        }
        return b;
    }
}
