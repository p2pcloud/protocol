// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./ProviderRegistry.sol";
import "./VerifiableOffer.sol";
import "./Payments.sol";
import "./AddressBook.sol";

abstract contract BrokerV3 is VerifiableOfferV3, ProviderRegistryV3, PaymentsV3, AddressBookV3 {
    function bookResource(UnsignedOffer calldata offer, bytes calldata signature) external {
        address provider = _getOfferProvider(offer, signature);
        require(isProviderRegistered(provider), "Provider is not registered");
        _executeClaimPayment(provider, offer.client);

        uint32 providerId = idByAddress(provider);
        uint32 clientId = idByAddress(offer.client);

        bookings[bookingCount] = Booking({
            specs: offer.specs,
            pricePerMinute: offer.pricePerMinute,
            clientId: clientId,
            providerId: providerId,
            startTime: uint32(block.timestamp)
        });

        require(
            getFreeBalance(msg.sender) >= offer.pricePerMinute * MONEY_LOCK_MINUTES,
            "Not enough balance to add a new the booking"
        );

        userProviderAccounting[provider][msg.sender].pricePerMinute += offer.pricePerMinute;
        _totalSpendingPerMinute[msg.sender] += offer.pricePerMinute;

        emit BookingCreated(bookingCount, offer.pricePerMinute, offer.client, provider);

        bookingCount = bookingCount + 1;
        nonce[msg.sender]++;
    }

    function cancelBooking(uint32 bookingId, bool satisfyed) external {
        Booking memory booking = bookings[bookingId];

        uint8 reason;

        address client = addressById(booking.clientId);
        address provider = addressById(booking.providerId);

        if (client == msg.sender) {
            if (satisfyed) {
                reason = CANCEL_REASON_NOT_NEEDED;
            } else {
                reason = CANCEL_REASON_NOT_SATISFIED;
            }
        } else if (provider == msg.sender) {
            reason = CANCEL_REASON_PROVIDER;
        } else {
            revert("Only client or provider can cancel the booking");
        }

        _executeClaimPayment(provider, client);

        if (getTotalBalance(client) == 0) {
            reason = CANCEL_REASON_NON_PAYMENT;
        }

        userProviderAccounting[provider][client].pricePerMinute -= booking.pricePerMinute;
        _totalSpendingPerMinute[client] -= booking.pricePerMinute;

        delete bookings[bookingId];

        emit BookingCancelled(bookingId, reason);
    }

    function listClientBookings(address client) external view returns (BookingFull[] memory) {
        uint32 clientId = idByAddressReadOnly(client);

        uint256 count = 0;
        for (uint32 i = 0; i < bookingCount; i++) {
            if (bookings[i].clientId == clientId) {
                count++;
            }
        }

        BookingFull[] memory result = new BookingFull[](count);
        uint256 index = 0;
        for (uint32 i = 0; i < bookingCount; i++) {
            if (bookings[i].clientId == clientId) {
                result[index] = BookingFull({
                    id: i,
                    specs: bookings[i].specs,
                    pricePerMinute: bookings[i].pricePerMinute,
                    client: client,
                    provider: addressById(bookings[i].providerId),
                    startTime: bookings[i].startTime
                });
                index++;
            }
        }

        return result;
    }

    function listProvidersBookings(address provider) external view returns (BookingFull[] memory) {
        uint32 providerId = idByAddressReadOnly(provider);

        uint32 count = 0;
        for (uint32 i = 0; i < bookingCount; i++) {
            if (bookings[i].providerId == providerId) {
                count++;
            }
        }

        BookingFull[] memory result = new BookingFull[](count);
        uint32 index = 0;
        for (uint32 i = 0; i < bookingCount; i++) {
            if (bookings[i].providerId == providerId) {
                result[index] = BookingFull({
                    id: i,
                    specs: bookings[i].specs,
                    pricePerMinute: bookings[i].pricePerMinute,
                    client: addressById(bookings[i].clientId),
                    provider: provider,
                    startTime: bookings[i].startTime
                });
                index++;
            }
        }

        return result;
    }

    function getBooking(uint32 id) external view returns (BookingFull memory) {
        return
            BookingFull({
                id: id,
                specs: bookings[id].specs,
                pricePerMinute: bookings[id].pricePerMinute,
                client: addressById(bookings[id].clientId),
                provider: addressById(bookings[id].providerId),
                startTime: bookings[id].startTime
            });
    }
}
