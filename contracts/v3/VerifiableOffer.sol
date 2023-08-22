// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./DelegatedSigner.sol";
import "./Storage.sol";

abstract contract VerifiableOfferV3 is DelegatedSignerV3 {
    using ECDSA for bytes32;

    function getNonce(address client) external view returns (uint32) {
        return nonce[client];
    }

    function __VerifiableOffer_init() internal onlyInitializing {
        uint256 chainId;
        assembly {
            chainId := chainid()
        }

        DOMAIN_SEPARATOR = keccak256(
            abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("p2pcloud.io")), keccak256(bytes("2")), chainId, address(this))
        );
    }

    function _getOfferProvider(UnsignedOffer calldata offer, bytes calldata signature) internal view returns (address) {
        require(offer.expiresAt > block.timestamp, "Offer expired");
        require(offer.client == msg.sender, "Invalid offer client");
        require(offer.nonce == nonce[offer.client], "Invalid offer nonce");

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
