// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

error ProviderIsNotRegistered();
error KYCProblem(address addr, bytes2 country);
error OfferExpired(uint256 expiresAt, uint256 now);
error OfferUserInvalid(address expected, address actual);
error OfferWrongNonce(uint32 currentNonce, uint32 offerNonce);
error ERC20TransferFailed();
error InsufficientBalance(uint256 requested, uint256 available);
error InvalidKYCSigner(address expected, address actual);
error NotAuthorized();
error MigrationComplete(uint8);

/*
    We have to keep all variables in one place.
    This way we can easily add variables or change order of inheritance without __gap problems.
*/

abstract contract StorageV3 is OwnableUpgradeable {
    // ----------------------
    // addressBook
    // ----------------------
    mapping(address => uint32) internal addressToID;
    mapping(uint32 => address) internal idToAddress;
    uint32 internal nextID;

    address public KYCSigner;

    // ----------------------
    // VerifiableKYC
    // ----------------------
    mapping(address => bytes2) public KYCStatus;
    mapping(bytes2 => bool) public AllowedUserCountries;
    mapping(bytes2 => bool) public AllowedProviderCountries;

    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
    bytes32 public constant OFFER_TYPEHASH =
        keccak256("UnsignedOffer(address client,uint64 pricePerMinute,uint32 nonce,bytes32 specs,uint256 expiresAt)");

    bytes32 public DOMAIN_SEPARATOR;

    // ----------------------
    // BalanceHolder
    // ----------------------
    uint16 public constant COMMUNITY_FEE = 2000;

    IERC20 public coin;

    mapping(address => uint256) internal _coinBalance;
    mapping(bytes12 => bool) public mintBurnIdempotency;

    // ----------------------
    // VerifiableOffer
    // ----------------------

    struct UnsignedOffer {
        address client; //20
        uint64 pricePerMinute; //8
        uint32 nonce; //4
        bytes32 specs; //32
        uint256 expiresAt; //32
    }

    mapping(address => uint32) internal nonce;

    // ----------------------
    // ProviderRegistry
    // ----------------------

    struct ProviderInfo {
        //slot 1
        bytes32 url;
        //slot 2
        bool isRegistered;
        uint128 feePaid;
    }

    mapping(address => ProviderInfo) providerInfo;
    address[] public providerList;

    // ----------------------
    // Broker
    // ----------------------

    uint8 public constant CANCEL_REASON_NOT_NEEDED = 0;
    uint8 public constant CANCEL_REASON_NOT_SATISFIED = 1;
    uint8 public constant CANCEL_REASON_PROVIDER = 2;
    uint8 public constant CANCEL_REASON_NON_PAYMENT = 3;

    uint32 public bookingCount;

    struct Booking {
        bytes32 specs;
        uint64 pricePerMinute;
        uint32 clientId;
        uint32 providerId;
        uint32 startTime;
    }

    struct BookingFull {
        uint32 id;
        bytes32 specs;
        uint64 pricePerMinute;
        address client;
        address provider;
        uint32 startTime;
    }

    mapping(uint32 => Booking) public bookings;

    event BookingCreated(uint256 bookingId, uint64 pricePerMinute, address client, address provider);
    event BookingCancelled(uint256 bookingId, uint8 reason);

    // ----------------------
    // DelegatedSigner
    // ----------------------
    mapping(address => address) internal _walletToSigner;
    mapping(address => address) internal _signerToWallet;

    // ----------------------
    // Payments
    // ----------------------
    uint256 public constant MONEY_LOCK_MINUTES = 60 * 24 * 7; // 7 days

    struct UserProviderAccounting {
        uint128 lastPaymentTs;
        uint128 pricePerMinute;
    }

    mapping(address => mapping(address => UserProviderAccounting)) public userProviderAccounting;
    mapping(address => uint256) internal _totalSpendingPerMinute;

    // ----------------------
    // Marketplace V2->V3 migration
    // ----------------------
    bool public v2MigrationComplete;
    mapping(address => bool) V2balanceMigrationComplete;
}
