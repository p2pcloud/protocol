// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./Marketplace.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract TestableMarketplace is Marketplace {
    using ECDSA for bytes32;

    function test__getTime() public view returns (uint256) {
        return block.timestamp;
    }

    function test__increaseSpendingPerMinute(address user, uint256 amt) public {
        _totalSpendingPerMinute[user] += amt;
    }

    function test__GetOfferSigner(
        UnsignedOffer calldata offer,
        bytes calldata signature
    ) public view returns (address) {
        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                keccak256(
                    abi.encode(
                        OFFER_TYPEHASH,
                        offer.client,
                        offer.pricePerMinute,
                        offer.nonce,
                        offer.specs,
                        offer.expiresAt
                    )
                )
            )
        );

        address recoveredSigner = digest.recover(signature);
        return resolveSigner(recoveredSigner);
    }
}
