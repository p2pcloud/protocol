// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/// @custom:security-contact security@p2pcloud.io
abstract contract P2PCloudCreditBaseERC20 is
    Initializable,
    ERC20Upgradeable,
    ERC20BurnableUpgradeable,
    PausableUpgradeable,
    OwnableUpgradeable
{
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function __P2PCloudCreditBaseERC20_init() internal {
        __ERC20_init("P2PCloud Credit", "PPC");
        __ERC20Burnable_init();
        __Pausable_init();
        __Ownable_init();
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount) internal override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);
    }
}

contract P2PCloudCredit is P2PCloudCreditBaseERC20 {
    using ECDSA for bytes32;

    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");

    bytes32 public DOMAIN_SEPARATOR;

    //voucher logic
    bytes32 public constant VOUCHER_TYPEHASH = keccak256("UnsignedVoucher(uint256 amount,bytes32 paymentId)");

    struct UnsignedVoucher {
        uint256 amount;
        bytes32 paymentId;
    }

    event VoucherClaimed(address indexed client, uint256 amount, bytes32 paymentId);
    event CoinBurned(address indexed client, uint256 amount);

    mapping(bytes32 => bool) public usedVouchers;
    address public voucherSigner;

    function __P2PCloudCredit_init() internal onlyInitializing {
        uint256 chainId;
        assembly {
            chainId := chainid()
        }

        DOMAIN_SEPARATOR = keccak256(
            abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("p2pcloud.io")), keccak256(bytes("2")), chainId, address(this))
        );
    }

    function initialize(address _voucherSigner) public initializer {
        __P2PCloudCreditBaseERC20_init();
        __P2PCloudCredit_init();
        voucherSigner = _voucherSigner;
        //TODO: check if second init is possible
    }

    function setVoucherSigner(address _voucherSigner) public onlyOwner {
        voucherSigner = _voucherSigner;
    }

    function isVoucherAlreadyClaimed(bytes32 paymentId) public view returns (bool) {
        return usedVouchers[paymentId];
    }

    function claimVoucher(UnsignedVoucher calldata voucher, bytes calldata signature, address receiver) public payable {
        require(usedVouchers[voucher.paymentId] == false, "Voucher already used");
        usedVouchers[voucher.paymentId] = true;

        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(VOUCHER_TYPEHASH, voucher.amount, voucher.paymentId))
            )
        );

        address recoveredSigner = digest.recover(signature);
        require(recoveredSigner == voucherSigner, "Invalid signature");

        _mint(receiver, voucher.amount);

        if (msg.value > 0) {
            payable(receiver).transfer(msg.value);
        }

        emit VoucherClaimed(receiver, voucher.amount, voucher.paymentId);
    }
}
