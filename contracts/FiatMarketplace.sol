// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";
import "./Broker.sol";
import "./Marketplace.sol";

contract FiatMarketplace is Marketplace {
    using ECDSA for bytes32;

    //voucher logic
    bytes32 public constant VOUCHER_TYPEHASH = keccak256("Voucher(uint64 amount,bytes32 paymentId)");

    struct Voucher {
        address client;
        uint64 amount;
        bytes32 paymentId;
    }

    mapping(bytes32 => bool) public usedVouchers;
    address public voucherSigner;

    function setVoucherSigner(address _voucherSigner) public onlyOwner {
        voucherSigner = _voucherSigner;
    }

    function claimVoucher(Voucher calldata voucher, bytes calldata signature) public {
        require(usedVouchers[voucher.paymentId] == false, "Voucher already used");

        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(VOUCHER_TYPEHASH, voucher.client, voucher.amount, voucher.paymentId))
            )
        );

        address recoveredSigner = digest.recover(signature);
        require(recoveredSigner == voucherSigner, "Invalid signature");

        usedVouchers[voucher.paymentId] = true;
        _coinBalance[msg.sender] = _coinBalance[msg.sender] + voucher.amount;
    }

    //only provider is allowed to withddraw
    function withdrawCoin(uint256 amt) public override {
        require(isProviderRegistered(msg.sender), "Only provider can withdraw");
        require(getFreeBalance(msg.sender) >= amt, "Not enough balance to withdraw");
        _coinBalance[msg.sender] -= amt;
        require(coin.transfer(msg.sender, amt), "ERC20 transfer failed");
    }

    //allow clients balance burn for refunds
    function burnCoin(uint256 amt, address client) public onlyOwner {
        _coinBalance[client] -= amt; //no worries about overflow since only owner can burn
    }

    function registerProvider() public override {
        require(false, "Not supported");
    }

    function registerFiatProvider(address provider) public onlyOwner {
        _registerProvider(provider);
    }
}
