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

contract Broker {
    struct Booking {
        uint256 index;
        uint256 vmTypeId;
        address miner;
        address user;
        uint256 pricePerSecond;
        uint256 bookedAt;
        uint256 lastPayment;
    }

    struct VMOffer {
        uint256 index;
        address miner;
        uint256 pricePerSecond;
        uint256 machinesAvailable;
        uint256 vmTypeId;
    }

    event BookingStarted(
        uint256 index,
        address indexed miner,
        address indexed user,
        uint256 vmTypeId
    );
    event BookingReported(
        uint256 minerPayout,
        uint256 index,
        address indexed miner,
        address indexed user,
        uint256 timeUsed,
        uint256 vmTypeId
    );
    event BookingStopped(
        uint256 minerPayout,
        uint256 index,
        address indexed miner,
        address indexed user,
        uint256 timeUsed,
        uint256 vmTypeId
    );
    event BookingClaimed(
        uint256 minerPayout,
        uint256 index,
        address indexed miner,
        address indexed user,
        uint256 timeUsed,
        uint256 vmTypeId
    );
    event BookingExtended(
        uint256 index,
        address indexed miner,
        address indexed user,
        uint256 vmTypeId
    );

    uint256 public constant SECONDS_IN_WEEK = 604800;

    uint256 public constant version = 1;

    mapping(uint256 => VMOffer) vmOffers;
    uint256 nextVmOfferId = 0;

    mapping(uint256 => Booking) bookings;
    uint256 nextBookingId = 0;

    mapping(address => bytes32) minerUrls;

    mapping(address => uint256) stablecoinBalance;
    mapping(address => uint256) userTotalPps;

    IERC20 stablecoin;

    address communityContract;
    uint256 communityFee;

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

    function addOffer(
        uint256 pricePerSecond,
        uint256 vmTypeId,
        uint256 machinesAvailable
    ) public returns (uint256) {
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
        uint256 offerIndex,
        uint256 machinesAvailable,
        uint256 pps
    ) public {
        require(
            vmOffers[offerIndex].miner == msg.sender,
            "Only the owner can update an offer"
        );
        vmOffers[offerIndex].machinesAvailable = machinesAvailable;
        vmOffers[offerIndex].pricePerSecond = pps;
    }

    function RemoveOffer(uint256 offerIndex) public {
        require(
            vmOffers[offerIndex].miner == msg.sender,
            "Only the owner can remove an offer"
        );
        delete vmOffers[offerIndex];
    }

    function DepositCoins(uint256 numTokens) public returns (bool) {
        if (!stablecoin.transferFrom(msg.sender, address(this), numTokens)) {
            return false;
        }

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

    function WithdrawStablecoin() public returns (bool) {
        uint256 freeBalance = stablecoinBalance[msg.sender] -
            getLockedStablecoinBalance(msg.sender);

        if (!stablecoin.transfer(msg.sender, freeBalance)) {
            return false;
        }

        stablecoinBalance[msg.sender] -= freeBalance;
        return true;
    }

    function GetStablecoinBalance(address user)
        public
        view
        returns (uint256, uint256)
    {
        uint256 locked = getLockedStablecoinBalance(user);
        return (stablecoinBalance[user] - locked, locked);
    }

    function BookVM(uint256 offerIndex) public returns (uint256) {
        require(
            vmOffers[offerIndex].machinesAvailable > 0,
            "No machines available"
        );

        require(
            (userTotalPps[msg.sender] + vmOffers[offerIndex].pricePerSecond) *
                SECONDS_IN_WEEK >=
                stablecoinBalance[msg.sender],
            "You don't have enough balance to pay for this and all other VMs for 7 days"
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

        return nextBookingId - 1;
    }

    function FindBookingsByUser(address _owner)
        public
        view
        returns (Booking[] memory filteredBookings)
    {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint256 count;
        for (uint256 i = 0; i < nextBookingId; i++) {
            if (bookings[i].user == _owner) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint256 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function FindBookingsByMiner(address _miner)
        public
        view
        returns (Booking[] memory filteredBookings)
    {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint256 count;
        for (uint256 i = 0; i < nextBookingId; i++) {
            if (bookings[i].miner == _miner) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint256 i = 0; i < count; i++) {
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

    function GetAvailableOffersByType(uint256 _vmTypeId)
        public
        view
        returns (VMOffer[] memory filteredOffers)
    {
        VMOffer[] memory offersTemp = new VMOffer[](nextVmOfferId);
        uint256 count;
        for (uint256 i = 0; i < nextVmOfferId; i++) {
            if (
                vmOffers[i].vmTypeId == _vmTypeId &&
                vmOffers[i].machinesAvailable > 0
            ) {
                offersTemp[count] = vmOffers[i];
                count += 1;
            }
        }

        filteredOffers = new VMOffer[](count);
        for (uint256 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function GetAvailableOffers()
        public
        view
        returns (VMOffer[] memory filteredOffers)
    {
        VMOffer[] memory offersTemp = new VMOffer[](nextVmOfferId);
        uint256 count;
        for (uint256 i = 0; i < nextVmOfferId; i++) {
            if (vmOffers[i].machinesAvailable > 0) {
                offersTemp[count] = vmOffers[i];
                count += 1;
            }
        }

        filteredOffers = new VMOffer[](count);
        for (uint256 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function GetUsersBookings(address user)
        public
        view
        returns (Booking[] memory filteredBookings)
    {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint256 count;
        for (uint256 i = 0; i < nextBookingId; i++) {
            if (bookings[i].user == user) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint256 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function GetMinersOffers(address miner)
        public
        view
        returns (VMOffer[] memory filteredOffers)
    {
        VMOffer[] memory offersTemp = new VMOffer[](nextVmOfferId);
        uint256 count;
        for (uint256 i = 0; i < nextVmOfferId; i++) {
            if (vmOffers[i].miner == miner) {
                offersTemp[count] = vmOffers[i];
                count += 1;
            }
        }

        filteredOffers = new VMOffer[](count);
        for (uint256 i = 0; i < count; i++) {
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

    function setCommunityFee(uint256 fee) public returns (bool) {
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
    function getCommunityFee() public view returns (uint256) {
        return communityFee;
    }
}
