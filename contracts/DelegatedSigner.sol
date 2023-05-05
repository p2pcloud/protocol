// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

contract DelegatedSigner {
    mapping(address => address) private _walletToSigner;
    mapping(address => address) private _signerToWallet;

    function setSigner(address signer) public {
        if (_signerToWallet[signer] != address(0)) {
            revert("Signer already in use");
        }

        address oldValue = _walletToSigner[msg.sender];
        if (oldValue != address(0)) {
            delete _signerToWallet[oldValue];
        }

        _walletToSigner[msg.sender] = signer;
        _signerToWallet[signer] = msg.sender;
    }

    function getSigner(address user) public view returns (address) {
        return _walletToSigner[user];
    }

    function resolveSigner(address signer) public view returns (address) {
        return _signerToWallet[signer];
    }
}
