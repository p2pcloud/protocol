// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

interface IERC20 {
    event Transfer(address indexed from, address indexed to, uint256 value);

    event Approval(address indexed owner, address indexed spender, uint256 value);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address to, uint256 amount) external returns (bool);

    function allowance(address owner, address spender) external view returns (uint256);

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
        uint256 bookedTill;
    }

    struct VMOffer {
        uint256 index;
        address miner;
        uint256 pricePerSecond;
        uint256 machinesAvailable;
        uint256 vmTypeId;
    }

    event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId);

    uint256 public constant version = 1;

    mapping(uint256 => VMOffer) vmOffers;
    uint256 nextVmOfferId = 0;

    mapping(uint256 => Booking) bookings;
    uint256 nextBookingId = 0;

    mapping(address => bytes32) minerUrls;
    mapping(address => bytes20) mtlsHashes;

    mapping(address => uint256) deposits;
    mapping(address => uint256) locked;

    IERC20 token;

    address community;
    uint256 communityFee;

    constructor(address communityAddress) {
        community = communityAddress;
    }

    function SetMtlsHash(bytes20 _signature) public {
        mtlsHashes[msg.sender] = _signature;
    }

    function getMtlsHash(address _user) public view returns (bytes20) {
        return mtlsHashes[_user];
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
        uint256 pricePerSecond,
        uint256 vmTypeId,
        uint256 machinesAvailable
    ) public {
        require(
            vmOffers[offerIndex].miner == msg.sender,
            "Only the owner can update an offer"
        );
        vmOffers[offerIndex].pricePerSecond = pricePerSecond;
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

        locked[msg.sender] += vmOffers[offerIndex].pricePerSecond * secs;

        Booking memory booking = Booking(
            nextBookingId,
            vmOffers[offerIndex].vmTypeId,
            vmOffers[offerIndex].miner,
            msg.sender,
            vmOffers[offerIndex].pricePerSecond,
            block.timestamp + secs
        );
        bookings[nextBookingId] = booking;
        nextBookingId++;

        emit BookingStarted(booking.index, booking.miner, booking.user, booking.vmTypeId);

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
        require(
            userBalance() >= numTokens,
            "insufficient funds to withdraw"
        );

        if (!token.transfer(msg.sender, numTokens)) {
            return false;
        }

        deposits[msg.sender] = deposits[msg.sender] - numTokens;
        return true;
    }

    function userBalance() public view returns (uint256) {
        return deposits[msg.sender] - locked[msg.sender];
    }

    function userTokenBalance() public view returns (uint256) {
        return token.balanceOf(msg.sender);
    }

    function userAllowance() public view returns (uint256) {
        return token.allowance(msg.sender, address(this));
    }

    function setCommunityContract(address newCommunityAddress) public returns (bool) {
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
            fee > 0 && fee < 101 &&
            msg.sender == community,
            "community fee should be in range of 1 to 100"
        );

        communityFee = fee;
        return false;
    }

    function getCommunityFee() public view returns (uint256) {
        return communityFee;
    }
}
