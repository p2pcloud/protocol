// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./BalanceHolder.sol";
import "./ProviderRegistry.sol";
import "./Broker.sol";

contract Marketplace is BalanceHolder, ProviderRegistry, Broker {
    bool private initialized;

    function initialize() public {
        require(initialized != true, "Already initialized");
        initialized = true;

        uint256 chainId;
        assembly {
            chainId := chainid()
        }

        DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                DOMAIN_TYPEHASH,
                keccak256(bytes("p2pcloud.io")),
                keccak256(bytes("2")),
                chainId,
                address(this)
            )
        );

        _transferOwnership(msg.sender);
        agreementCount = 1;

        _disableInitializers();

        communityFee = 2000;
    }
}
