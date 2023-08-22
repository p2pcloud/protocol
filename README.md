# P2PCloud smart contracts
Smart conracts for distributed confidential computing platform. Website: [https://p2pcloud.io](https://p2pcloud.io)

## Contract inheritance

Marketplace     -> Broker   -> VerifiableOffer  -> DelegatedSigner      -> Storage  
                            -> ProviderRegistry -> VerifiableKYC        -> Storage  
                            -> Payments         -> BalanceHolder        -> VerifiableKYC    -> Storage  
                            -> AddressBook      -> Storage  

## Contract Descriptions

- **Broker**: The primary contract that encompasses all logic, excluding migration.
- **Marketplace**: The gateway or entry point, responsible for migrations.
- **VerifiableOffer**: Assists with signature verification, particularly useful for gasless bookings on the provider's side.
- **DelegatedSigner**: Enables the delegation of gasless interactions, allowing a paying account to delegate signatures to a separate zero-balance account.
- **Storage**: Consolidates all storage variables into one contract, streamlining upgrades.
- **ProviderRegistry**: Manages the registration of new providers and verifies their registration status.
- **VerifiableKYC**: Handles the KYC verification process for both providers and customers.
- **Payments**: Manages transactions from clients to providers.
- **BalanceHolder**: Retains balances and manages both withdrawals and deposits.
- **AddressBook**: Converts addresses into uint32, offering significant gas savings during bookings.