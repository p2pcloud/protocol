// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

abstract contract AddressBook {
    mapping(address => uint32) private addressToID;
    mapping(uint32 => address) private idToAddress;
    uint32 private nextID;

    function idByAddress(address _address) internal returns (uint32) {
        uint32 existingID = addressToID[_address];
        if (existingID != 0) {
            return existingID; // If the address is already in the registry, return its uint32 ID.
        }

        uint32 newID = nextID;
        addressToID[_address] = newID;
        idToAddress[newID] = _address;
        nextID += 1;
        return newID;
    }

    function idByAddressReadOnly(address _address) internal view returns (uint32) {
        return addressToID[_address];
    }

    function addressById(uint32 _id) internal view returns (address) {
        return idToAddress[_id];
    }
}
