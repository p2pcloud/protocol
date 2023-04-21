// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BalanceHolder.sol";
import "./ProviderRegistry.sol";
import "./VerifiableOffer.sol";
import "./Payments.sol";
import "./AddressBook.sol";

abstract contract Broker is VerifiableOffer, BalanceHolder, ProviderRegistry, Payments, AddressBook {
    uint8 public constant CANCEL_REASON_NOT_NEEDED = 0;
    uint8 public constant CANCEL_REASON_NOT_SATISFIED = 1;
    uint8 public constant CANCEL_REASON_PROVIDER = 2;
    uint8 public constant CANCEL_REASON_NON_PAYMENT = 3;

    event BookingCreated(uint256 bookingId, uint64 pricePerMinute, address client, address provider);
    event BookingCancelled(uint256 bookingId, uint8 reason);

    struct Booking {
        bytes32 specs;
        uint64 pricePerMinute;
        uint32 clientId;
        uint32 providerId;
    }

    struct BookingFull {
        bytes32 specs;
        uint64 pricePerMinute;
        address client;
        address provider;
    }

    uint256 public bookingCount;
    mapping(uint256 => Booking) public bookings;

    function bookResource(UnsignedOffer calldata offer, bytes calldata signature) external {
        address provider = _getOfferProvider(offer, signature);

        _executeClaimPayment(provider, offer.client);

        require(isProviderRegistered(provider), "Provider is not registered");

        uint32 providerId = idByAddress(provider);
        uint32 clientId = idByAddress(offer.client);

        bookings[bookingCount] = Booking({
            specs: offer.specs,
            pricePerMinute: offer.pricePerMinute,
            clientId: clientId,
            providerId: providerId
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
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].clientId == clientId) {
                count++;
            }
        }

        BookingFull[] memory result = new BookingFull[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].clientId == clientId) {
                result[index] = BookingFull({
                    specs: bookings[i].specs,
                    pricePerMinute: bookings[i].pricePerMinute,
                    client: client,
                    provider: addressById(bookings[i].providerId)
                });
                index++;
            }
        }

        return result;
    }

    function listProvidersBookings(address provider) external view returns (BookingFull[] memory) {
        uint32 providerId = idByAddressReadOnly(provider);

        uint256 count = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].providerId == providerId) {
                count++;
            }
        }

        BookingFull[] memory result = new BookingFull[](count);
        uint256 index = 0;
        for (uint256 i = 0; i < bookingCount; i++) {
            if (bookings[i].providerId == providerId) {
                result[index] = BookingFull({
                    specs: bookings[i].specs,
                    pricePerMinute: bookings[i].pricePerMinute,
                    client: addressById(bookings[i].clientId),
                    provider: provider
                });
                index++;
            }
        }

        return result;
    }

    function getBooking(uint256 id) external view returns (BookingFull memory) {
        return
            BookingFull({
                specs: bookings[id].specs,
                pricePerMinute: bookings[id].pricePerMinute,
                client: addressById(bookings[id].clientId),
                provider: addressById(bookings[id].providerId)
            });
    }
}
