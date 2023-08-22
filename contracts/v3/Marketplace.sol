// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "./Broker.sol";
import "../v2/FiatMarketplace.sol";

contract MarketplaceV3 is BrokerV3 {
    function initialize(IERC20 _coin, address _migrationSource) public initializer {
        _transferOwnership(msg.sender);

        __VerifiableOffer_init();
        coin = _coin;

        V2migrationSource = _migrationSource;
    }

    //one-off migration logic
    //needed, because upgrade is not possible

    function performMigration() public onlyOwner {
        require(V2migrationSource != address(0), "Migration source is not set or migration already performed");
        require(bookingCount == 0, "Can not perform migration after bookings were created");
        V2migrationSource = address(0);

        uint32 totalBookings = MarketplaceV2(V2migrationSource).bookingCount();
        for (uint32 i = 0; i < totalBookings; i++) {
            MarketplaceV2.BookingFull memory oldBooking = MarketplaceV2(V2migrationSource).getBooking(i);
            bookingCount = bookingCount + 1;

            if (oldBooking.client == address(0) || oldBooking.provider == address(0)) {
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

            nonce[msg.sender]++;
        }
    }

    function migrateBalance(address addr) private {
        if (V2balanceMigrationComplete[addr]) {
            return;
        }
        V2balanceMigrationComplete[addr] = true;
    }
}
