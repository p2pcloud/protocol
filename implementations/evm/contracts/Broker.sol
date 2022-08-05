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
        uint256 bookedTill;
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

    uint256 public constant version = 1;

    mapping(uint256 => VMOffer) vmOffers;
    uint256 nextVmOfferId = 0;

    mapping(uint256 => Booking) bookings;
    uint256 nextBookingId = 0;

    mapping(address => bytes32) minerUrls;

    mapping(address => uint256) deposits;
    mapping(address => uint256) locked;

    IERC20 token;

    address community;
    uint256 communityFee;

    constructor(address communityAddress) {
        community = communityAddress;
    }

    function setMunerUrl(bytes32 url) public {
        minerUrls[msg.sender] = url;
    }

    function getMinerUrl(address _user) public view returns (bytes32) {
        return minerUrls[_user];
    }

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

    function updateOffer(
        uint256 offerIndex,
        uint256 vmTypeId,
        uint256 machinesAvailable
    ) public {
        require(
            vmOffers[offerIndex].miner == msg.sender,
            "Only the owner can update an offer"
        );
        vmOffers[offerIndex].machinesAvailable = machinesAvailable;
        vmOffers[offerIndex].vmTypeId = vmTypeId;
    }

    function bookVM(uint256 offerIndex, uint256 secs) public returns (uint256) {
        require(
            vmOffers[offerIndex].machinesAvailable > 0 &&
                secs > 0 &&
                userBalance() >= vmOffers[offerIndex].pricePerSecond * secs,
            "No machines available OR secs should be positive OR user does not have required amount to book machine"
        );
        require(
            vmOffers[offerIndex].miner != address(0),
            "offer with that index does not exist"
        );

        locked[msg.sender] += vmOffers[offerIndex].pricePerSecond * secs;

        Booking memory booking = Booking(
            nextBookingId,
            vmOffers[offerIndex].vmTypeId,
            vmOffers[offerIndex].miner,
            msg.sender,
            vmOffers[offerIndex].pricePerSecond,
            block.timestamp,
            block.timestamp + secs
        );
        bookings[nextBookingId] = booking;
        nextBookingId++;

        emit BookingStarted(
            booking.index,
            booking.miner,
            booking.user,
            booking.vmTypeId
        );

        return nextBookingId - 1;
    }

    function findBookingsByUser(address _owner)
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

    function findBookingsByMiner(address _miner)
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

    function getBooking(uint64 index)
        public
        view
        returns (Booking memory booking)
    {
        booking = bookings[index];
    }

    function extendBooking(uint64 index, uint256 secs) public {
        Booking memory booking = bookings[index];
        uint256 currentTime = GetTime();

        require(msg.sender == booking.user, "only owner of booking can extend");
        require(currentTime < booking.bookedTill, "booking is expired");
        require(
            userBalance() >= booking.pricePerSecond * secs,
            "insufficient funds to extend booking"
        );

        locked[msg.sender] += booking.pricePerSecond * secs;

        booking.bookedTill += secs;
        bookings[index] = booking;

        emit BookingExtended(
            booking.index,
            booking.miner,
            booking.user,
            booking.vmTypeId
        );
    }

    function abortBooking(uint64 index, uint8 abortType) public {
        Booking memory booking = bookings[index];
        uint256 currentTime = GetTime();

        require(msg.sender == booking.user, "only owner of booking can report");
        require(
            abortType == 1 || abortType == 2,
            "report or stop booking is required"
        );
        require(
            currentTime < booking.bookedTill,
            "booking is expired, miner should claim expired booking"
        );

        uint256 minersPercentage = 0;
        uint256 communityPercentage = 0;

        if (abortType == 1) {
            minersPercentage = 50;
            communityPercentage = 50;
        } else {
            minersPercentage = 100 - communityFee;
            communityPercentage = communityFee;
        }

        uint256 toPay = booking.pricePerSecond *
            (currentTime - booking.bookedAt);

        uint256 minerPayout = (toPay * minersPercentage) / 100;
        uint256 communityPayout = toPay - minerPayout;

        deposits[msg.sender] -= toPay;
        locked[msg.sender] -=
            booking.pricePerSecond *
            (booking.bookedTill - booking.bookedAt);

        delete bookings[booking.index];

        if (abortType == 1) {
            emit BookingReported(
                minerPayout,
                booking.index,
                booking.miner,
                booking.user,
                currentTime - booking.bookedAt,
                booking.vmTypeId
            );
        } else {
            emit BookingStopped(
                minerPayout,
                booking.index,
                booking.miner,
                booking.user,
                currentTime - booking.bookedAt,
                booking.vmTypeId
            );
        }

        if (!token.transfer(booking.miner, minerPayout)) {
            revert("could not payout for miner");
        }

        if (!token.transfer(community, communityPayout)) {
            revert("could not payout for community");
        }
    }

    function claimExpired(uint64 index) public {
        Booking memory booking = bookings[index];
        uint256 currentTime = GetTime();

        require(msg.sender == booking.miner, "only miner of booking can claim");
        require(currentTime >= booking.bookedTill, "booking is still in use");

        uint256 toPay = booking.pricePerSecond *
            (booking.bookedTill - booking.bookedAt);

        uint256 minerPayout = (toPay * (100 - communityFee)) / 100;
        uint256 communityPayout = toPay - minerPayout;

        deposits[booking.user] -= toPay;
        locked[booking.user] -= toPay;

        delete bookings[booking.index];

        emit BookingClaimed(
            minerPayout,
            booking.index,
            booking.miner,
            booking.user,
            booking.bookedTill - booking.bookedAt,
            booking.vmTypeId
        );

        if (!token.transfer(booking.miner, minerPayout)) {
            revert("could not payout for miner");
        }

        if (!token.transfer(community, communityPayout)) {
            revert("could not payout for community");
        }
    }

    function getAvailableOffers(uint256 _vmTypeId)
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

    function getUsersBookings(address user)
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

    function getMinersOffers(address miner)
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

    function setStablecoinAddress(IERC20 t) public returns (bool) {
        require(
            msg.sender == community,
            "only community contract can set stablecoin"
        );

        token = t;

        return false;
    }

    function getStablecoinAddress() public view returns (IERC20) {
        return token;
    }

    function getCoinDecimals() public view returns (uint8) {
        return token.decimals();
    }

    function depositCoin(uint256 numTokens) public returns (bool) {
        if (!token.transferFrom(msg.sender, address(this), numTokens)) {
            return false;
        }

        deposits[msg.sender] = deposits[msg.sender] + numTokens;
        return true;
    }

    function withdrawCoin(uint256 numTokens) public returns (bool) {
        require(userBalance() >= numTokens, "insufficient funds to withdraw");

        if (!token.transfer(msg.sender, numTokens)) {
            return false;
        }

        deposits[msg.sender] = deposits[msg.sender] - numTokens;
        return true;
    }

    function userBalance() public view returns (uint256) {
        return deposits[msg.sender] - locked[msg.sender];
    }

    function userDeposit() public view returns (uint256) {
        return deposits[msg.sender];
    }

    function userLockedBalance() public view returns (uint256) {
        return locked[msg.sender];
    }

    function userTokenBalance() public view returns (uint256) {
        return token.balanceOf(msg.sender);
    }

    function userAllowance() public view returns (uint256) {
        return token.allowance(msg.sender, address(this));
    }

    function setCommunityContract(address newCommunityAddress)
        public
        returns (bool)
    {
        require(
            msg.sender == community,
            "only community contract can set new community contract"
        );

        community = newCommunityAddress;
        return false;
    }

    function getCommunityContract() public view returns (address) {
        return community;
    }

    function setCommunityFee(uint256 fee) public returns (bool) {
        require(
            fee > 0 && fee < 101 && msg.sender == community,
            "community fee should be in range of 1 to 100"
        );

        communityFee = fee;
        return false;
    }

    function getCommunityFee() public view returns (uint256) {
        return communityFee;
    }
}
