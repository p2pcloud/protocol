// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./BalanceHolder.sol";
import "./ProviderRegistry.sol";
import "./VerifiableOffer.sol";

abstract contract Broker is VerifiableOffer, BalanceHolder {
    struct UserProviderAccounting {
        uint256 lastPaymentTs;
        uint256 pricePerMinute;
    }

    mapping(address => mapping(address => UserProviderAccounting)) public userProviderAccounting;

    struct Booking {
        bytes32 specs;
        uint256 pricePerMinute;
        address client;
        address provider;
    }

    uint256 public bookingCount;

    mapping(uint256 => Booking) public bookings;
    mapping(address => uint256) internal _totalSpendingPerMinute;

    //TODO: solve a client zero balance problem

    uint256 public constant MONEY_LOCK_MINUTES = 60 * 24 * 7; // 7 days

    event BookingCreated(uint256 bookingId, uint256 pricePerMinute, address client, address provider);

    function bookResource(UnsignedOffer calldata offer, bytes calldata signature) external {
        address provider = _getOfferProvider(offer, signature);

        //TODO: execute payment claim
        //TODO: check enough non-locked balance
        //TODO: check provider is registered

        bookings[bookingCount] = Booking({
            specs: offer.specs,
            pricePerMinute: offer.pricePerMinute,
            client: offer.client,
            provider: provider
        });
        emit BookingCreated(bookingCount, offer.pricePerMinute, offer.client, provider);

        userProviderAccounting[provider][msg.sender].pricePerMinute += offer.pricePerMinute;

        bookingCount++;
        nonce[msg.sender]++;
        _totalSpendingPerMinute[msg.sender] += offer.pricePerMinute;
    }

    function _getLockedBalance(address user) internal view override returns (uint256) {
        return _totalSpendingPerMinute[user] * MONEY_LOCK_MINUTES;
    }

    function listClientBookings(address client) external view returns (Booking[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i <= bookingCount; i++) {
            if (bookings[i].client == client) {
                count++;
            }
        }

        Booking[] memory result = new Booking[](count);
        uint256 index = 0;
        for (uint256 i = 1; i <= bookingCount; i++) {
            if (bookings[i].client == client) {
                result[index] = bookings[i];
                index++;
            }
        }

        return result;
    }

    function listProvidersBookings(address provider) external view returns (Booking[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i <= bookingCount; i++) {
            if (bookings[i].provider == provider) {
                count++;
            }
        }

        Booking[] memory result = new Booking[](count);
        uint256 index = 0;
        for (uint256 i = 1; i <= bookingCount; i++) {
            if (bookings[i].provider == provider) {
                result[index] = bookings[i];
                index++;
            }
        }

        return result;
    }

    function getBooking(uint256 id) external view returns (Booking memory) {
        return bookings[id];
    }
}
