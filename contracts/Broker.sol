// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

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

    uint32 public constant MONEY_LOCK_MINUTES = 60 * 24 * 7; // 7 days

    event BookingCreated(uint256 bookingId, uint256 pricePerMinute, address client, address provider);
    event BookingCancelled(uint256 bookingId, uint8 reason);

    function bookResource(UnsignedOffer calldata offer, bytes calldata signature) external {
        address provider = _getOfferProvider(offer, signature);

        _executeClaimPayment(provider, offer.client);

        //TODO: check enough non-locked balance
        //TODO: check provider is registered

        bookings[bookingCount] = Booking({
            specs: offer.specs,
            pricePerMinute: offer.pricePerMinute,
            client: offer.client,
            provider: provider
        });
        emit BookingCreated(bookingCount, offer.pricePerMinute, offer.client, provider);

        require(
            getFreeBalance(msg.sender) >= offer.pricePerMinute * MONEY_LOCK_MINUTES,
            "Not enough balance to add a new the booking"
        );
        userProviderAccounting[provider][msg.sender].pricePerMinute += offer.pricePerMinute;

        bookingCount++;
        nonce[msg.sender]++;
        _totalSpendingPerMinute[msg.sender] += offer.pricePerMinute;
    }

    function _executeClaimPayment(address provider, address client) internal {
        // userProviderAccounting[provider][client].lastPaymentTs = block.timestamp;
        uint256 pricePerMinute = userProviderAccounting[provider][client].pricePerMinute;

        if (pricePerMinute == 0) {
            //initialize account
            userProviderAccounting[provider][client].lastPaymentTs = block.timestamp;
            return;
        }

        uint256 minutesPassed = (block.timestamp - userProviderAccounting[provider][client].lastPaymentTs) / 60;
        uint256 amount = minutesPassed * pricePerMinute;

        _spendWithComission(client, provider, amount);
        userProviderAccounting[provider][client].lastPaymentTs += minutesPassed * 60;
    }

    uint8 public constant CANCEL_REASON_NOT_NEEDED = 0;
    uint8 public constant CANCEL_REASON_NOT_SATISFIED = 1;
    uint8 public constant CANCEL_REASON_PROVIDER = 2;
    uint8 public constant CANCEL_REASON_NON_PAYMENT = 3;

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

    function getLockedBalance(address user) public view override returns (uint256) {
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

    function claimPayment(address client) external {
        _executeClaimPayment(msg.sender, client);
    }

    function getBooking(uint256 id) external view returns (Booking memory) {
        return bookings[id];
    }
}
