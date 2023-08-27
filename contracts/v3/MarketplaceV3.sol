// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./BrokerV3.sol";
import "../v2/FiatMarketplace.sol";

contract MarketplaceV3 is BrokerV3 {
    function initialize(IERC20 _coin) public initializer {
        _transferOwnership(msg.sender);

        __VerifiableOffer_init();
        __AddressBook_init();
        coin = _coin;
    }

    //one-off migration logic
    //needed, because upgrade is not possible

    function performMigration(MarketplaceV2 V2migrationSource) public onlyOwner {
        require(bookingCount == 0, "MGRTN_BOOK_EX");
        require(v2MigrationComplete == false, "MGRTN_COMPLETE");
        v2MigrationComplete = true;

        uint32 totalBookings = V2migrationSource.bookingCount();
        bookingCount = totalBookings;

        for (uint32 i = 0; i < totalBookings; i++) {
            MarketplaceV2.BookingFull memory oldBooking = V2migrationSource.getBooking(i);

            if (oldBooking.pricePerMinute == 0) {
                continue;
            }

            bookings[i] = Booking({
                specs: oldBooking.specs,
                pricePerMinute: oldBooking.pricePerMinute,
                clientId: idByAddress(oldBooking.client),
                providerId: idByAddress(oldBooking.provider),
                startTime: uint32(block.timestamp)
            });

            userProviderAccounting[oldBooking.provider][oldBooking.client].pricePerMinute += oldBooking.pricePerMinute;
            _totalSpendingPerMinute[oldBooking.client] += oldBooking.pricePerMinute;

            emit BookingCreated(i, oldBooking.pricePerMinute, oldBooking.client, oldBooking.provider);

            migrateUser(oldBooking.client, V2migrationSource);
            migrateUser(oldBooking.provider, V2migrationSource);

            nonce[oldBooking.client]++;
        }
    }

    function migrateUser(address addr, MarketplaceV2 V2migrationSource) private {
        if (V2balanceMigrationComplete[addr]) {
            return;
        }
        V2balanceMigrationComplete[addr] = true;
        _coinBalance[addr] = V2migrationSource.getTotalBalance(addr);
        if (V2migrationSource.isProviderRegistered(addr)) {
            _addProviderToDB(addr);
        }
    }
}
