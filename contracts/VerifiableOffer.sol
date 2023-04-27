// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./DelegatedSigner.sol";

abstract contract VerifiableOffer is Initializable, DelegatedSigner {
    using ECDSA for bytes32;

    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
    bytes32 public constant OFFER_TYPEHASH =
        keccak256("UnsignedOffer(address client,uint64 pricePerMinute,uint32 nonce,bytes32 specs,uint256 expiresAt)");

    bytes32 public DOMAIN_SEPARATOR;

    struct UnsignedOffer {
        address client; //20
        uint64 pricePerMinute; //8
        uint32 nonce; //4
        bytes32 specs; //32
        uint256 expiresAt; //32
    }

    mapping(address => uint32) internal nonce;
    

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
