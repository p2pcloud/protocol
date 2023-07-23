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

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount) internal virtual override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);
    }
}

contract P2PCloudCredit is BaseERC20 {
    using ECDSA for bytes32;
    address allowedRecipient;
    address trustedMinter;

    function initialize(address _trustedMinter, address _allowedRecipient) public initializer {
        require(_trustedMinter != address(0), "Invalid trustedMinter address");
        require(_allowedRecipient != address(0), "Invalid allowedRecipient address");
        __BaseERC20_init();
        allowedRecipient = _allowedRecipient;
        trustedMinter = _trustedMinter;
    }

    mapping(bytes12 => bool) public mintedIds;

    //mint with idempotency
    function idemopotentMint(address to, bytes12 mintId, uint256 amount) public {
        require(msg.sender == trustedMinter, "Only trustedMinter can mint");
        require(mintId != 0, "Mint id cannot be zero");
        require(!mintedIds[mintId], "Mint id already used");
        mintedIds[mintId] = true;
        _mint(to, amount);
    }

    //restrict token transfers to only allowedRecipient
    function _beforeTokenTransfer(address from, address to, uint256 amount) internal override whenNotPaused {
        require(to == allowedRecipient || to == address(this), "Only allowedRecipient can receive tokens");
        super._beforeTokenTransfer(from, to, amount);
    }
}
