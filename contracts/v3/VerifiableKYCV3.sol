// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./StorageV3.sol";

abstract contract VerifiableKYCV3 is StorageV3 {
    using ECDSA for bytes32;

    function _validateSignature(address _address, bytes2 country, bytes memory signature) private view {
        bytes32 data = keccak256(abi.encodePacked(_address, country));
        address actualSigner = data.toEthSignedMessageHash().recover(signature);
        if (actualSigner != KYCSigner) {
            revert InvalidKYCSigner(KYCSigner, actualSigner);
        }
    }

    // Submit KYC
    function submitKYC(address _address, bytes2 country, bytes memory signature) public {
        _validateSignature(_address, country, signature);
        KYCStatus[_address] = country;
    }

    function allowUserCountry(bytes2 country) public onlyOwner {
        AllowedUserCountries[country] = true;
    }

    function allowProviderCountry(bytes2 country) public onlyOwner {
        AllowedProviderCountries[country] = true;
    }

    function isUserKYCPassed(address _address) public view returns (bool) {
        return AllowedUserCountries[KYCStatus[_address]];
    }

    function isProviderKYCPassed(address _address) public view returns (bool) {
        return AllowedProviderCountries[KYCStatus[_address]];
    }

    function setKYCSigner(address _KYCSigner) public onlyOwner {
        KYCSigner = _KYCSigner;
    }
}
