// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

import "./BalanceHolder.sol";
import "./CommunityOwnable.sol";
import "./ProviderRegistry.sol";
import "./Testable.sol";

contract Broker is BalanceHolder, CommunityOwnable, ProviderRegistry, Testable {
    constructor() {}
    //TODO: SubmitAmendment
    //TODO: BreakContract
    //TODO: ClaimPayment
    //TODO: ListClientsContracts
    //TODO: ListProvidersContracts
    //TODO: GetContract
}
