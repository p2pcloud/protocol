// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

interface IERC20 {
    event Transfer(address indexed from, address indexed to, uint256 value);

    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    function balanceOf(address account) external view returns (uint256);

    function transfer(address to, uint256 amount) external returns (bool);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    function decimals() external view returns (uint8);

    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);
}

//TODO: optimize uint types to use less gas

contract Broker {
    struct Booking {
        uint64 index;
        uint64 vmTypeId;
        address miner;
        address user;
        uint64 pricePerSecond;
        uint256 bookedAt;
        uint256 lastPayment;
    }

    struct VMOffer {
        uint64 index;
        address miner;
        uint64 pricePerSecond;
        uint64 machinesAvailable;
        uint64 vmTypeId;
    }

    event BookingStarted(
        uint64 index,
        address indexed miner,
        address indexed user,
        uint64 vmTypeId
    );
    event BookingReported(
        uint64 minerPayout,
        uint64 index,
        address indexed miner,
        address indexed user,
        uint64 timeUsed,
        uint64 vmTypeId
    );
    event BookingStopped(
        uint64 minerPayout,
        uint64 index,
        address indexed miner,
        address indexed user,
        uint64 timeUsed,
        uint64 vmTypeId
    );
    event BookingClaimed(
        uint64 minerPayout,
        uint64 index,
        address indexed miner,
        address indexed user,
        uint64 timeUsed,
        uint64 vmTypeId
    );
    event BookingExtended(
        uint64 index,
        address indexed miner,
        address indexed user,
        uint64 vmTypeId
    );

    uint64 public constant SECONDS_IN_WEEK = 604800;

    uint64 public constant version = 1;

    mapping(uint64 => VMOffer) vmOffers;
    uint64 nextVmOfferId = 0;

    mapping(uint64 => Booking) bookings;
    uint64 nextBookingId = 0;

    mapping(address => bytes32) minerUrls;

    mapping(address => uint256) stablecoinBalance;
    mapping(address => uint64) userTotalPps;

    IERC20 stablecoin;

    address communityContract;
    uint64 communityFee = 500; // 5%

    constructor() {
        communityContract = msg.sender;
    }

    function setMunerUrl(bytes32 url) public {
        minerUrls[msg.sender] = url;
    }

    function getMinerUrl(address _user) public view returns (bytes32) {
        return minerUrls[_user];
    }

    //TODO: do we need this?!
    function GetTime() public view returns (uint256) {
        return block.timestamp;
    }

    function AddOffer(
        uint64 pricePerSecond,
        uint64 vmTypeId,
        uint64 machinesAvailable
    ) public returns (uint64) {
        vmOffers[nextVmOfferId] = VMOffer(
            nextVmOfferId,
            msg.sender,
            pricePerSecond,
            machinesAvailable,
            vmTypeId
        );
        nextVmOfferId++;
        return nextVmOfferId - 1;
    }

    function UpdateOffer(
        uint64 offerIndex,
        uint64 machinesAvailable,
        uint64 pps
    ) public {
        require(
            vmOffers[offerIndex].miner == msg.sender,
            "Only the owner can update an offer"
        );
        vmOffers[offerIndex].machinesAvailable = machinesAvailable;
        vmOffers[offerIndex].pricePerSecond = pps;
    }

    function RemoveOffer(uint64 offerIndex) public {
        require(
            vmOffers[offerIndex].miner == msg.sender,
            "Only the owner can remove an offer"
        );
        delete vmOffers[offerIndex];
    }

    function DepositStablecoin(uint256 numTokens) public returns (bool) {
        require(
            stablecoin.transferFrom(msg.sender, address(this), numTokens),
            "Failed to transfer tokens"
        );

        stablecoinBalance[msg.sender] =
            stablecoinBalance[msg.sender] +
            numTokens;
        return true;
    }

    function getLockedStablecoinBalance(address user)
        private
        view
        returns (uint256)
    {
        return userTotalPps[user] * SECONDS_IN_WEEK;
    }

    function WithdrawStablecoin(uint256 amt) public returns (bool) {
        uint256 freeBalance = stablecoinBalance[msg.sender] -
            getLockedStablecoinBalance(msg.sender);

        require(freeBalance >= amt, "Not enough balance to withdraw");

        require(stablecoin.transfer(msg.sender, amt), "ERC20 transfer failed");

        stablecoinBalance[msg.sender] -= amt;
        return true;
    }

    function ClaimPayment(uint64 bookingId) public {
        require(
            bookings[bookingId].miner == msg.sender,
            "Only the miner can claim a payment"
        );

        uint256 timeUsed = block.timestamp - bookings[bookingId].lastPayment;

        uint256 totalPayout = timeUsed * bookings[bookingId].pricePerSecond;

        if (stablecoinBalance[bookings[bookingId].user] < totalPayout) {
            totalPayout = stablecoinBalance[bookings[bookingId].user];
        }

        uint256 communityPayout = (totalPayout * communityFee) / (100 * 100);
        uint256 minerPayout = totalPayout - communityPayout;

        bookings[bookingId].lastPayment = block.timestamp;

        stablecoinBalance[communityContract] += communityPayout;
        stablecoinBalance[bookings[bookingId].miner] += minerPayout;
        stablecoinBalance[bookings[bookingId].user] -= totalPayout;
    }

    function GetStablecoinBalance(address user)
        public
        view
        returns (uint256, uint256)
    {
        uint256 locked = getLockedStablecoinBalance(user);
        return (stablecoinBalance[user] - locked, locked);
    }

    function BookVM(uint64 offerIndex) public returns (uint64) {
        require(
            vmOffers[offerIndex].machinesAvailable > 0,
            "No machines available"
        );

        uint256 willBeLocked = getLockedStablecoinBalance(msg.sender) +
            vmOffers[offerIndex].pricePerSecond *
            SECONDS_IN_WEEK;

        require(
            willBeLocked <= stablecoinBalance[msg.sender],
            "You don't have enough balance to pay for this and all other VMs for 7 days "
        );

        Booking memory booking = Booking(
            nextBookingId,
            vmOffers[offerIndex].vmTypeId,
            vmOffers[offerIndex].miner,
            msg.sender,
            vmOffers[offerIndex].pricePerSecond,
            block.timestamp,
            block.timestamp
        );
        bookings[nextBookingId] = booking;
        nextBookingId++;

        userTotalPps[msg.sender] += vmOffers[offerIndex].pricePerSecond;

        vmOffers[offerIndex].machinesAvailable -= 1;

        return nextBookingId - 1;
    }

    function FindBookingsByUser(address _owner)
        public
        view
        returns (Booking[] memory filteredBookings)
    {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint64 count;
        for (uint64 i = 0; i < nextBookingId; i++) {
            if (bookings[i].user == _owner) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint64 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function FindBookingsByMiner(address _miner)
        public
        view
        returns (Booking[] memory filteredBookings)
    {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint64 count;
        for (uint64 i = 0; i < nextBookingId; i++) {
            if (bookings[i].miner == _miner) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint64 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function GetBooking(uint64 index)
        public
        view
        returns (Booking memory booking)
    {
        booking = bookings[index];
    }

    // function extendBooking(uint64 index, uint256 secs) public {
    //     Booking memory booking = bookings[index];
    //     uint256 currentTime = GetTime();

    //     require(msg.sender == booking.user, "only owner of booking can extend");
    //     require(currentTime < booking.bookedTill, "booking is expired");
    //     require(
    //         userBalance() >= booking.pricePerSecond * secs,
    //         "insufficient funds to extend booking"
    //     );

    //     locked[msg.sender] += booking.pricePerSecond * secs;

    //     booking.bookedTill += secs;
    //     bookings[index] = booking;

    //     emit BookingExtended(
    //         booking.index,
    //         booking.miner,
    //         booking.user,
    //         booking.vmTypeId
    //     );
    // }

    // function stopBooking(uint64 index, bool complain) public {
    //     Booking memory booking = bookings[index];
    //     uint256 currentTime = GetTime();

    //     require(
    //         msg.sender == booking.user,
    //         "only owner of booking can stop it"
    //     );

    //     //TODO: pay to miner
    //     uint256 toPay = booking.pricePerSecond *
    //         (currentTime - booking.bookedAt);

    //     uint256 minerPayout = (toPay * minersPercentage) / 100;
    //     uint256 communityPayout = toPay - minerPayout;

    //     deposits[msg.sender] -= toPay;
    //     locked[msg.sender] -=
    //         booking.pricePerSecond *
    //         (booking.bookedTill - booking.bookedAt);

    //     delete bookings[booking.index];

    //     if (abortType == 1) {
    //         emit BookingReported(
    //             minerPayout,
    //             booking.index,
    //             booking.miner,
    //             booking.user,
    //             currentTime - booking.bookedAt,
    //             booking.vmTypeId
    //         );
    //     } else {
    //         emit BookingStopped(
    //             minerPayout,
    //             booking.index,
    //             booking.miner,
    //             booking.user,
    //             currentTime - booking.bookedAt,
    //             booking.vmTypeId
    //         );
    //     }

    //     if (!token.transfer(booking.miner, minerPayout)) {
    //         revert("could not payout for miner");
    //     }

    //     if (!token.transfer(community, communityPayout)) {
    //         revert("could not payout for community");
    //     }
    // }

    // function claimExpired(uint64 index) public {
    //     Booking memory booking = bookings[index];

    //     require(msg.sender == booking.miner, "only miner of booking can claim");
    //     require(
    //         block.timestamp >= booking.bookedTill,
    //         "booking is still in use"
    //     );

    //     uint256 toPay = booking.pricePerSecond *
    //         (booking.bookedTill - booking.bookedAt);

    //     uint256 minerPayout = (toPay * (100 - communityFee)) / 100;
    //     uint256 communityPayout = toPay - minerPayout;

    //     deposits[booking.user] -= toPay;
    //     locked[booking.user] -= toPay;

    //     delete bookings[booking.index];

    //     emit BookingClaimed(
    //         minerPayout,
    //         booking.index,
    //         booking.miner,
    //         booking.user,
    //         booking.bookedTill - booking.bookedAt,
    //         booking.vmTypeId
    //     );

    //     if (!token.transfer(booking.miner, minerPayout)) {
    //         revert("could not payout for miner");
    //     }

    //     if (!token.transfer(community, communityPayout)) {
    //         revert("could not payout for community");
    //     }
    // }

    function GetAvailableOffersByType(uint64 _vmTypeId)
        public
        view
        returns (VMOffer[] memory filteredOffers)
    {
        VMOffer[] memory offersTemp = new VMOffer[](nextVmOfferId);
        uint64 count;
        for (uint64 i = 0; i < nextVmOfferId; i++) {
            if (
                vmOffers[i].vmTypeId == _vmTypeId &&
                vmOffers[i].machinesAvailable > 0
            ) {
                offersTemp[count] = vmOffers[i];
                count += 1;
            }
        }

        filteredOffers = new VMOffer[](count);
        for (uint64 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function GetAvailableOffers()
        public
        view
        returns (VMOffer[] memory filteredOffers)
    {
        VMOffer[] memory offersTemp = new VMOffer[](nextVmOfferId);
        uint64 count;
        for (uint64 i = 0; i < nextVmOfferId; i++) {
            if (vmOffers[i].machinesAvailable > 0) {
                offersTemp[count] = vmOffers[i];
                count += 1;
            }
        }

        filteredOffers = new VMOffer[](count);
        for (uint64 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function GetUsersBookings(address user)
        public
        view
        returns (Booking[] memory filteredBookings)
    {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint64 count;
        for (uint64 i = 0; i < nextBookingId; i++) {
            if (bookings[i].user == user) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint64 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function GetMinersOffers(address miner)
        public
        view
        returns (VMOffer[] memory filteredOffers)
    {
        VMOffer[] memory offersTemp = new VMOffer[](nextVmOfferId);
        uint64 count;
        for (uint64 i = 0; i < nextVmOfferId; i++) {
            if (vmOffers[i].miner == miner) {
                offersTemp[count] = vmOffers[i];
                count += 1;
            }
        }

        filteredOffers = new VMOffer[](count);
        for (uint64 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function SetStablecoinAddress(IERC20 newStablecoinAddress) public {
        require(
            msg.sender == communityContract,
            "only community contract can set stablecoin"
        );

        stablecoin = newStablecoinAddress;
    }

    //TODO: do we need it?
    function getStablecoinAddress() public view returns (IERC20) {
        return stablecoin;
    }

    function SetCommunityContract(address newCommunityAddress) public {
        require(
            msg.sender == communityContract,
            "only community contract can set new community contract"
        );
        communityContract = newCommunityAddress;
    }

    //TODO: we may not need this with a public field
    function getCommunityContract() public view returns (address) {
        return communityContract;
    }

    function setCommunityFee(uint64 fee) public returns (bool) {
        require(
            fee < 10000,
            "community fee should be in range of 0 (0%) to 10000 (100%)"
        );
        require(
            msg.sender == communityContract,
            "only community contract can set community fee"
        );

        communityFee = fee;
        return false;
    }

    //TODO: we may not need this with a public field
    function getCommunityFee() public view returns (uint64) {
        return communityFee;
    }
}
