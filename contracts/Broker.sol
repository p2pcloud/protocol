// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./CommunityOwnable.sol";
import "./BalanceHolder.sol";
import "./ProviderRegistry.sol";
import "./utils/Testable.sol";

contract Broker is BalanceHolder, ProviderRegistry, Testable {
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256(
            "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
        );
    bytes32 public constant AMENDMENT_TYPEHASH =
        keccak256(
            "Amendment(uint256 agreementId,bytes32 ipfsHash,uint256 pricePerMinute)"
        );

    bytes32 public DOMAIN_SEPARATOR;

    bool private initialized;

    function initialize() public {
        require(initialized != true, "Already initialized");
        initialized = true;

        uint256 chainId;
        assembly {
            chainId := chainid()
        }

        DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                DOMAIN_TYPEHASH,
                keccak256(bytes("p2pcloud.io")),
                keccak256(bytes("2")),
                chainId,
                address(this)
            )
        );

        _transferOwnership(msg.sender);
    }
}

// contract Broker is BalanceHolder, ProviderRegistry, Testable {
// uint256 public constant MONEY_LOCK_MINUTES = 60 * 24 * 7; // 7 days

// using ECDSA for bytes32;

// struct Agreement {
//     bytes32 ipfsHash;
//     uint256 pricePerMinute;
//     address client;
//     address provider;
//     uint256 lastPayment;
//     uint16 nonce;
// }

// struct Amendment {
//     bytes32 ipfsHash;
//     uint256 pricePerMinute;
// }

// mapping(uint256 => Agreement) public agreements;
// uint256 public agreementCount;

// event AgreementAmended(
//     uint256 agreementId,
//     bytes32 ipfsHash,
//     uint256 pricePerMinute
// );
// event AgreementBroken(uint256 agreementId);
// event PaymentClaimed(uint256 agreementId, uint256 amount);

// function amendAgreement(
//     uint256 agreementId,
//     Amendment calldata amendment,
//     bytes calldata signature
// ) external {
//     Agreement storage agreement = agreements[agreementId];
//     require(
//         agreement.client == msg.sender,
//         "Only client can submit amendment"
//     );
//     require(agreement.provider != address(0), "Agreement not found");

//     bytes32 digest = keccak256(
//         abi.encodePacked(
//             "\x19\x01",
//             DOMAIN_SEPARATOR,
//             keccak256(
//                 abi.encode(
//                     AMENDMENT_TYPEHASH,
//                     agreementId,
//                     amendment.ipfsHash,
//                     amendment.pricePerMinute
//                 )
//             )
//         )
//     );

//     require(
//         digest.recover(signature) == agreement.provider,
//         "Invalid provider signature"
//     );

//     agreement.ipfsHash = amendment.ipfsHash;
//     agreement.pricePerMinute = amendment.pricePerMinute;

//     emit AgreementAmended(
//         agreementId,
//         amendment.ipfsHash,
//         amendment.pricePerMinute
//     );
// }

// function listClientsAgreements(
//     address client
// ) external view returns (Agreement[] memory) {
//     uint256 count = 0;
//     for (uint256 i = 1; i <= agreementCount; i++) {
//         if (agreements[i].client == client) {
//             count++;
//         }
//     }

//     Agreement[] memory clientAgreements = new Agreement[](count);
//     uint256 index = 0;
//     for (uint256 i = 1; i <= agreementCount; i++) {
//         if (agreements[i].client == client) {
//             clientAgreements[index] = agreements[i];
//             index++;
//         }
//     }

//     return clientAgreements;
// }

// function listProvidersAgreements(
//     address provider
// ) external view returns (Agreement[] memory) {
//     uint256 count = 0;
//     for (uint256 i = 1; i <= agreementCount; i++) {
//         if (agreements[i].provider == provider) {
//             count++;
//         }
//     }

//     Agreement[] memory providerAgreements = new Agreement[](count);
//     uint256 index = 0;
//     for (uint256 i = 1; i <= agreementCount; i++) {
//         if (agreements[i].provider == provider) {
//             providerAgreements[index] = agreements[i];
//             index++;
//         }
//     }

//     return providerAgreements;
// }

// function getAgreement(
//     uint256 agreementId
// ) external view returns (Agreement memory) {
//     return agreements[agreementId];
// }
// }
