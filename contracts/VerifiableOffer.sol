// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract VerifiableOffer is Initializable {
    using ECDSA for bytes32;

    mapping(address => uint32) internal nonce;

    struct UnsignedOffer {
        bytes32 specs;
        uint256 pricePerMinute;
        address client;
        uint256 expiresAt;
        uint32 nonce;
    }

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

    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
    bytes32 public constant OFFER_TYPEHASH =
        keccak256("UnsignedOffer(bytes32 specs,uint256 pricePerMinute,address client,uint256 expiresAt,uint32 nonce)");

    bytes32 public DOMAIN_SEPARATOR;

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
                        offer.specs,
                        offer.pricePerMinute,
                        offer.client,
                        offer.expiresAt,
                        offer.nonce
                    )
                )
            )
        );

        return digest.recover(signature);
    }
}