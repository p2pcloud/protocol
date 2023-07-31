// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/// @custom:security-contact security@p2pcloud.io
abstract contract BaseERC20 is
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

    function __BaseERC20_init() internal {
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

    function _beforeTokenTransfer(address from, address to, uint256 amount) internal virtual override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);
    }
}

contract P2PCloudCredit is BaseERC20 {
    using ECDSA for bytes32;
    //allowedRecipient is the address of marketplace to make sure the tokens are only used for the intended purpose
    address public allowedRecipient;
    address public trustedMinter;

    function initialize() public initializer {
        __BaseERC20_init();
    }

    mapping(bytes12 => bool) public mintedIds;

    function setAllowedRecipient(address _allowedRecipient) public onlyOwner {
        require(_allowedRecipient != address(0), "Invalid allowedRecipient address");
        allowedRecipient = _allowedRecipient;
    }

    function setTrustedMinter(address _trustedMinter) public onlyOwner {
        require(_trustedMinter != address(0), "Invalid trustedMinter address");
        trustedMinter = _trustedMinter;
    }

    //mint with idempotency
    function idemopotentMint(address payable to, uint256 amount, bytes12 mintId) public payable {
        require(msg.sender == trustedMinter, "Only trustedMinter can mint");
        require(mintId != 0, "Mint id cannot be zero");
        require(!mintedIds[mintId], "Mint id already used");
        mintedIds[mintId] = true;
        _mint(to, amount);

        // Transfer any ether sent with the transaction to the 'to' address.
        if (msg.value > 0) {
            to.transfer(msg.value);
        }
    }

    //restrict token transfers to only allowedRecipient
    function _beforeTokenTransfer(address from, address to, uint256 amount) internal override whenNotPaused {
        require(
            to == allowedRecipient || from == allowedRecipient || from == address(0) || to == address(0),
            "ERC20: Recipient not allowed"
        );
        super._beforeTokenTransfer(from, to, amount);
    }
}
