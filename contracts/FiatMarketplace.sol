// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";
import "./Broker.sol";
import "./Marketplace.sol";

contract FiatMarketplace is Marketplace {
    using ECDSA for bytes32;

    //voucher logic
    bytes32 public constant VOUCHER_TYPEHASH = keccak256("UnsignedVoucher(uint256 amount,bytes32 paymentId)");

    struct UnsignedVoucher {
        uint256 amount;
        bytes32 paymentId;
    }

    event VoucherClaimed(address indexed client, uint256 amount, bytes32 paymentId);
    event CoinBurned(address indexed client, uint256 amount);

    mapping(bytes32 => bool) public usedVouchers;
    address public voucherSigner;

    function setVoucherSigner(address _voucherSigner) public onlyOwner {
        voucherSigner = _voucherSigner;
    }

    function isVoucherAlreadyClaimed(bytes32 paymentId) public view returns (bool) {
        return usedVouchers[paymentId];
    }

    function claimVoucher(UnsignedVoucher calldata voucher, bytes calldata signature, address receiver) public {
        require(usedVouchers[voucher.paymentId] == false, "Voucher already used");
        usedVouchers[voucher.paymentId] = true;

        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(VOUCHER_TYPEHASH, voucher.amount, voucher.paymentId))
            )
        );

        address recoveredSigner = digest.recover(signature);
        require(recoveredSigner == voucherSigner, "Invalid signature");

        _coinBalance[receiver] = _coinBalance[receiver] + voucher.amount;

        emit VoucherClaimed(receiver, voucher.amount, voucher.paymentId);
    }

    //only provider is allowed to withddraw
    function withdrawCoin(uint256 amt) public pure override {
        require(false, "Not supported");
    }

    //deposit is disabled
    function depositCoin(uint256 amt) public pure override {
        require(false, "Not supported");
    }

    //bun coins instead of withdrawals
    function burnCoin(uint256 amt, address client) public onlyOwner {
        require(_coinBalance[client] >= amt, "Not enough balance to burn");
        _coinBalance[client] -= amt; //no worries about overflow since only owner can burn
        emit CoinBurned(client, amt);
    }

    function registerProvider() public pure override {
        require(false, "Not supported");
    }

    function registerFiatProvider(address provider) public onlyOwner {
        _registerProvider(provider, 0);
    }
}
