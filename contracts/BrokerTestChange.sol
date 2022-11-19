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

contract BrokerTestChange {
    struct Booking {
        uint64 index;
        uint64 vmTypeId;
        address miner;
        address user;
        uint64 pricePerSecond;
        uint256 bookedAt;
        uint256 lastPayment;
        uint64 offerIndex;
    }

    struct Offer {
        uint64 index;
        address miner;
        uint64 pricePerSecond;
        uint64 machinesAvailable;
        uint64 vmTypeId;
    }

    mapping(uint64 => Offer) offers;
    uint64 nextVmOfferId;

    mapping(uint64 => Booking) bookings;
    uint64 nextBookingId;

    mapping(address => uint256) coinBalance;
    mapping(address => uint64) userTotalPps;

    mapping(address => bytes32) minerUrls;

    IERC20 coin;

    address communityContract;
    uint64 communityFee;

    uint64 public constant SECONDS_IN_WEEK = 604800;

    event Payment(address indexed user, address indexed miner, uint256 amount);

    event Complaint(
        address indexed user,
        address indexed miner,
        uint8 indexed reason
    );

    function RenamedGetMinerUrl(address _user) public view returns (bytes32) {
        return minerUrls[_user];
    }
}
