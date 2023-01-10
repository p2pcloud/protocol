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

    function allowance(
        address owner,
        address spender
    ) external view returns (uint256);

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
        address miner;
        address user;
        uint32 index;
        uint32 offerIndex;
        uint32 pricePerSecond;
        uint32 bookedAt; //TODO: change timestamp to uint48
        uint32 lastPayment; //TODO: change timestamp to uint48
    }

    //TODO: try rearranging fields to optimize gas usage
    struct Offer {
        address miner;
        uint32 index;
        uint32 pricePerSecond;
        bytes32 specsIpfsHash;
        uint16 machinesTotal;
        uint16 machinesBooked;
    }

    mapping(uint32 => Offer) offers;
    uint32 nextVmOfferId;

    mapping(uint32 => Booking) bookings;
    uint32 nextBookingId;

    mapping(address => uint256) coinBalance;

    mapping(address => uint32) userTotalPps;

    IERC20 public coin;

    address public communityContract;
    uint16 public communityFee;

    uint64 public constant SECONDS_IN_WEEK = 604800;

    event Payment(address indexed user, address indexed miner, uint256 amount);

    event Termination(uint32 bookingIndex, uint16 indexed reason);

    event NewBooking(
        address indexed user,
        address indexed miner,
        uint32 bookingIndex,
        uint32 pps
    );

    struct MinerInfo {
        bytes32 url;
        bool isRegistered;
        uint256 feePaid;
    }

    mapping(address => MinerInfo) minerInfo;

    //01_miner_url

    function SetMinerUrl(bytes32 url) public {
        require(
            minerInfo[msg.sender].isRegistered,
            "Miner must be registered to set url"
        );
        minerInfo[msg.sender].url = url;
    }

    function GetMinerUrl(address _user) public view returns (bytes32) {
        return minerInfo[_user].url;
    }

    function IsMinerRegistered(address _user) public view returns (bool) {
        return minerInfo[_user].isRegistered;
    }

    //02_offers

    uint64 public constant MINER_REGISTRATION_FEE = 100 * 1000000;

    function RegisterMiner() public {
        uint256 freeBalance = coinBalance[msg.sender] -
            GetLockedCoinBalance(msg.sender);
        require(
            freeBalance >= MINER_REGISTRATION_FEE,
            "Not enough coin to register "
        );
        require(
            !minerInfo[msg.sender].isRegistered,
            "Miner is already registered"
        );

        coinBalance[msg.sender] -= MINER_REGISTRATION_FEE;
        coinBalance[communityContract] += MINER_REGISTRATION_FEE;

        minerInfo[msg.sender].isRegistered = true;
        minerInfo[msg.sender].feePaid += MINER_REGISTRATION_FEE;
    }

    function AddOffer(
        uint32 pricePerSecond,
        uint16 machinesTotal,
        bytes32 specsIpfsHash
    ) public returns (uint32) {
        require(
            minerInfo[msg.sender].isRegistered,
            "Miner must be registered to add offers"
        );

        offers[nextVmOfferId] = Offer(
            msg.sender,
            nextVmOfferId,
            pricePerSecond,
            specsIpfsHash,
            machinesTotal,
            0
        );
        nextVmOfferId++;
        return nextVmOfferId - 1;
    }

    function UpdateOffer(
        uint32 offerIndex,
        uint16 machinesTotal,
        uint32 pps
    ) public {
        require(
            offers[offerIndex].miner == msg.sender,
            "Only the owner can update an offer"
        );
        require(
            offers[offerIndex].machinesBooked <= machinesTotal,
            "Cannot reduce machinesTotal below machinesBooked"
        );
        offers[offerIndex].machinesTotal = machinesTotal;
        offers[offerIndex].pricePerSecond = pps;
    }

    function RemoveOffer(uint32 offerIndex) public {
        require(
            offers[offerIndex].miner == msg.sender,
            "Only the owner can remove an offer"
        );
        require(
            offers[offerIndex].machinesBooked == 0,
            "Cannot remove an offer with booked machines"
        );
        delete offers[offerIndex];
    }

    function GetMinersOffers(
        address miner
    ) public view returns (Offer[] memory filteredOffers) {
        Offer[] memory offersTemp = new Offer[](nextVmOfferId);
        uint32 count;
        for (uint32 i = 0; i < nextVmOfferId; i++) {
            if (offers[i].miner == miner) {
                offersTemp[count] = offers[i];
                count += 1;
            }
        }

        filteredOffers = new Offer[](count);
        for (uint32 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function GetAvailableOffers()
        public
        view
        returns (Offer[] memory filteredOffers)
    {
        Offer[] memory offersTemp = new Offer[](nextVmOfferId);
        uint32 count;
        for (uint32 i = 0; i < nextVmOfferId; i++) {
            if (offers[i].machinesTotal > offers[i].machinesBooked) {
                offersTemp[count] = offers[i];
                count += 1;
            }
        }

        filteredOffers = new Offer[](count);
        for (uint32 i = 0; i < count; i++) {
            filteredOffers[i] = offersTemp[i];
        }
    }

    function GetOffer(
        uint32 index
    ) public view returns (Offer memory oneOffer) {
        oneOffer = offers[index];
    }

    //03_coin

    function SetCoinAddress(IERC20 newCoinAddress) public {
        require(
            msg.sender == communityContract,
            "only community contract can set coin"
        );

        coin = newCoinAddress;
    }

    function DepositCoin(uint256 numTokens) public returns (bool) {
        require(
            coin.transferFrom(msg.sender, address(this), numTokens),
            "Failed to transfer tokens"
        );

        coinBalance[msg.sender] = coinBalance[msg.sender] + numTokens;
        return true;
    }

    function GetLockedCoinBalance(address user) private view returns (uint256) {
        return userTotalPps[user] * SECONDS_IN_WEEK;
    }

    function WithdrawCoin(uint256 amt) public returns (bool) {
        uint256 freeBalance = coinBalance[msg.sender] -
            GetLockedCoinBalance(msg.sender);

        require(freeBalance >= amt, "Not enough balance to withdraw");

        require(coin.transfer(msg.sender, amt), "ERC20 transfer failed");

        coinBalance[msg.sender] -= amt;
        return true;
    }

    function GetCoinBalance(
        address user
    ) public view returns (uint256 free, uint256 locked) {
        locked = GetLockedCoinBalance(user);
        free = coinBalance[user] - locked;
        return (free, locked);
    }

    //04_bookings

    function Book(uint32 offerIndex) public returns (uint32) {
        require(
            offers[offerIndex].machinesTotal >
                offers[offerIndex].machinesBooked,
            "No machines available"
        );

        uint256 willBeLocked = GetLockedCoinBalance(msg.sender) +
            offers[offerIndex].pricePerSecond *
            SECONDS_IN_WEEK;

        require(
            willBeLocked <= coinBalance[msg.sender],
            "You don't have enough balance to pay for this and all other Bookings for 7 days "
        );

        Booking memory booking = Booking(
            offers[offerIndex].miner,
            msg.sender,
            nextBookingId,
            offerIndex,
            offers[offerIndex].pricePerSecond,
            uint32(block.timestamp),
            uint32(block.timestamp)
        );
        bookings[nextBookingId] = booking;
        nextBookingId++;

        userTotalPps[msg.sender] += offers[offerIndex].pricePerSecond;

        offers[offerIndex].machinesBooked += 1;

        emit NewBooking(
            msg.sender, // address indexed user,
            offers[offerIndex].miner, // address indexed miner,
            nextBookingId - 1, // uint32 bookingIndex,
            offers[offerIndex].pricePerSecond // uint32 pps
        );

        return nextBookingId - 1;
    }

    function _executeBookingDelete(uint32 bookingId, uint16 reason) private {
        offers[bookings[bookingId].offerIndex].machinesBooked -= 1;
        userTotalPps[bookings[bookingId].user] -= bookings[bookingId]
            .pricePerSecond;

        emit Termination(bookingId, reason);

        delete bookings[bookingId];
    }

    uint16 public constant REASON_MINER_TERMINATED = 2;
    uint16 public constant REASON_NON_PAYMENT = 3;
    uint16 public constant REASON_COMMUNITY_TERMINATED = 4;

    function Terminate(uint32 bookingId, uint16 reason) public {
        require(
            bookings[bookingId].user == msg.sender ||
                (bookings[bookingId].miner == msg.sender &&
                    reason == REASON_MINER_TERMINATED) ||
                (msg.sender == communityContract &&
                    reason == REASON_COMMUNITY_TERMINATED),
            "Only the user can stop a VM with reason 0 or 1, only the miner can stop a VM with reason 2"
        );

        _executeClaimPayment(bookingId);
        _executeBookingDelete(bookingId, reason);
    }

    function ClaimPayment(uint32 bookingId) public {
        require(
            bookings[bookingId].miner == msg.sender,
            "Only the miner can claim a payment"
        );
        bool enoughMoney = _executeClaimPayment(bookingId);
        if (!enoughMoney) {
            _executeBookingDelete(bookingId, REASON_NON_PAYMENT);
        }
    }

    function _executeClaimPayment(uint32 bookingId) private returns (bool) {
        bool enoughMoney = true;
        uint256 timeUsed = block.timestamp - bookings[bookingId].lastPayment;

        uint256 totalPayout = timeUsed * bookings[bookingId].pricePerSecond;

        if (coinBalance[bookings[bookingId].user] < totalPayout) {
            totalPayout = coinBalance[bookings[bookingId].user];
            enoughMoney = false;
        }

        uint256 communityPayout = (totalPayout * communityFee) / (100 * 100);
        uint256 minerPayout = totalPayout - communityPayout;

        bookings[bookingId].lastPayment = uint32(block.timestamp);

        coinBalance[communityContract] += communityPayout;
        coinBalance[bookings[bookingId].miner] += minerPayout;
        coinBalance[bookings[bookingId].user] -= totalPayout;

        emit Payment(
            bookings[bookingId].user,
            bookings[bookingId].miner,
            minerPayout
        );
        return enoughMoney;
    }

    function FindBookingsByUser(
        address _owner
    ) public view returns (Booking[] memory filteredBookings) {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint32 count;
        for (uint32 i = 0; i < nextBookingId; i++) {
            if (bookings[i].user == _owner) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint32 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function FindBookingsByMiner(
        address _miner
    ) public view returns (Booking[] memory filteredBookings) {
        Booking[] memory bookingsTemp = new Booking[](nextBookingId);
        uint32 count;
        for (uint32 i = 0; i < nextBookingId; i++) {
            if (bookings[i].miner == _miner) {
                bookingsTemp[count] = bookings[i];
                count += 1;
            }
        }

        filteredBookings = new Booking[](count);
        for (uint32 i = 0; i < count; i++) {
            filteredBookings[i] = bookingsTemp[i];
        }
    }

    function GetBooking(
        uint32 index
    ) public view returns (Booking memory booking) {
        booking = bookings[index];
    }

    //05_community

    function SetCommunityContract(address newCommunityAddress) public {
        require(
            (msg.sender == communityContract) ||
                (communityContract == address(0)),
            "only community contract can set new community contract"
        );
        communityContract = newCommunityAddress;
    }

    function SetCommunityFee(uint16 fee) public returns (bool) {
        require(
            fee < 2500,
            "community fee should be in range of 0 (0%) to 2500 (25%)"
        );
        require(
            msg.sender == communityContract,
            "only community contract can set community fee"
        );

        communityFee = fee;
        return false;
    }

    function GetTime() public view returns (uint256) {
        //TODO: remove. Test function only
        return block.timestamp;
    }
}