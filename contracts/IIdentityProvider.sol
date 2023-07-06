// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface IIdentityProvider {
    /*
     * @dev Returns the country and a region of a customer
     * @param _customer The address of the customer
     * @return verified true if the customer is verified, false otherwise
     * @return country ISO 3166-1 alpha 2 country code of the government ID
     * @return subdivision ISO 3166-2 subdivision code of the government ID (State for the US)
     */
    function getRegion(address) external view returns (bool, bytes2, bytes3);

    function getCountry(address) external view returns (bool, bytes2);

    function getVerificationUrl() external view returns (string memory);
}
