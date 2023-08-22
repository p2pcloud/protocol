// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./Storage.sol";

abstract contract VerifiableKYCV3 is  StorageV3 {
    using ECDSA for bytes32;

    function _validSignature(address _address, bytes2 country, bytes memory signature) private view returns (bool) {
        bytes32 data = keccak256(abi.encodePacked(_address, country));
        return data.toEthSignedMessageHash().recover(signature) == KYCSigner;
    }

    // Submit KYC
    function submitKYC(address _address, bytes2 country, bytes memory signature) public {
        require(_validSignature(_address, country, signature), "Invalid signature");
        KYCStatus[_address] = country;
    }

    // Check KYC status
    function isKYCCompleted(address _address) public view returns (bool) {
        return (KYCStatus[_address] != bytes2(0x0000));
    }

    function allowUserCountry(bytes2 country) public onlyOwner {
        AllowedUserCountries[country] = true;
    }

    function disallowUserCountry(bytes2 country) public onlyOwner {
        AllowedUserCountries[country] = false;
    }

    function allowProviderCountry(bytes2 country) public onlyOwner {
        AllowedProviderCountries[country] = true;
    }

    function disallowProviderCountry(bytes2 country) public onlyOwner {
        AllowedProviderCountries[country] = false;
    }

    function checkUserKYC(address _address) public view returns (bool) {
        return (AllowedUserCountries[KYCStatus[_address]]);
    }

    function checkProviderKYC(address _address) public view returns (bool) {
        return (AllowedProviderCountries[KYCStatus[_address]]);
    }

    function setKYCSigner(address _KYCSigner) public onlyOwner {
        KYCSigner = _KYCSigner;
    }
}
