// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

import "./BalanceHolder.sol";
import "./ProviderRegistry.sol";
import "./VerifiableOffer.sol";
import "./Payments.sol";

abstract contract Broker is VerifiableOffer, BalanceHolder, ProviderRegistry, Payments {
    uint8 public constant CANCEL_REASON_NOT_NEEDED = 0;
    uint8 public constant CANCEL_REASON_NOT_SATISFIED = 1;
    uint8 public constant CANCEL_REASON_PROVIDER = 2;
    uint8 public constant CANCEL_REASON_NON_PAYMENT = 3;

    event BookingCreated(uint256 bookingId, uint256 pricePerMinute, address client, address provider);
    event BookingCancelled(uint256 bookingId, uint8 reason);

    struct Booking {
        bytes32 specs;
        uint256 pricePerMinute;
        address client;
        address provider;
    }

    uint256 public bookingCount;
    mapping(uint256 => Booking) public bookings;

    function bookResource(UnsignedOffer calldata offer, bytes calldata signature) external {
        address provider = _getOfferProvider(offer, signature);

        _executeClaimPayment(provider, offer.client);

        require(isProviderRegistered(provider), "Provider is not registered");

        bookings[bookingCount] = Booking({
            specs: offer.specs,
            pricePerMinute: offer.pricePerMinute,
            client: offer.client,
            provider: provider
        });

        require(
            getFreeBalance(msg.sender) >= offer.pricePerMinute * MONEY_LOCK_MINUTES,
            "Not enough balance to add a new the booking"
        );

        userProviderAccounting[provider][msg.sender].pricePerMinute += offer.pricePerMinute;
        _totalSpendingPerMinute[msg.sender] += offer.pricePerMinute;

        emit BookingCreated(bookingCount, offer.pricePerMinute, offer.client, provider);

        bookingCount++;
        nonce[msg.sender]++;
    }

    function cancelBooking(uint256 bookingId, bool satisfyed) external {
        Booking memory booking = bookings[bookingId];

        uint8 reason = 0;

        if (booking.client == msg.sender) {
            if (satisfyed) {
                reason = CANCEL_REASON_NOT_NEEDED;
            } else {
                reason = CANCEL_REASON_NOT_SATISFIED;
            }
        } else if (booking.provider == msg.sender) {
            reason = CANCEL_REASON_PROVIDER;
        } else {
            revert("Only client or provider can cancel the booking");
        }

        _executeClaimPayment(booking.provider, booking.client);

        if (getTotalBalance(booking.client) == 0) {
            reason = CANCEL_REASON_NON_PAYMENT;
        }

        userProviderAccounting[booking.provider][booking.client].pricePerMinute -= booking.pricePerMinute;
        _totalSpendingPerMinute[booking.client] -= booking.pricePerMinute;

        delete bookings[bookingId];

        emit BookingCancelled(bookingId, reason);
    }

    function listClientBookings(address client) external view returns (Booking[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].client == client) {
                count++;
            }
        }

        Booking[] memory result = new Booking[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].client == client) {
                result[index] = bookings[i];
                index++;
            }
        }

        return result;
    }

    function listProvidersBookings(address provider) external view returns (Booking[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].provider == provider) {
                count++;
            }
        }

        Booking[] memory result = new Booking[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
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
