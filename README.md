# Blockchain logic for P2PCloud

![Badge](https://github.com/p2pcloud/protocol/actions/workflows/go.yml/badge.svg)


## What is p2pcloud
[p2pcloud.io](https://p2pcloud.io) is a protocol and utility set for booking and securing trustless virtual machines over the blockchain.

## Development 
### Agreements
- All implementations should be in the implementations folder in a separate subfolder;
- All implementations should comply with `BlockchainIface`;
- Good test coverage with edge cases;
- Deliver changes through pull reqeusts.

### How to 
Compile contracts
```
go run ./cmd/compile/
```
Test:
```
go test ./...
```
### Current state of things
Undertested and a bit messy

## Business logic

### Actors
There are 3 actors - miner, user and community. User books and pays for VM. Miner runs VMs. Community gets comission from every VM booking. 

### Contracts
1. Broker contract. Facilitates deals between miners and users.
2. Community contract. Some contract used to vote on SDK and protocol updates. Also accepts and distributes fees. Not implemented yet.
3. Satblecoin contract. Some external stablecoin, probably USDC. May be replaced.

### Workflow
1. Miner creates offer setting type and price of VM per second
    - Functions: `GetMinersOffers`, `UpdateOffer`, `AddOffer`, `RemoveOffer`
1. User sends a stablecoin to the Broker.
    - Functions: `depositCoin`, `checkBalance`, `withdrawCoin`
1. User/Miner withdraws a stablecoin from the Broker. Vatiable `totalPPS` get's checked to make sure there is enough money for a week of work.
    - Functions: `withdrawCoin`
1. User gets list of offers
    - Functions: `getAvailableOffers`
1. User books VM. System checks that he has enough money to pay for all his VMs for 7 days (variable `totalPPS`).
    - Functions: `bookVM`
1. User aborts booking. With reason (0 - ok, 1 - miner misbehaviour)
    - Functions: `stopBooking` 
    - Events: `bookingReported`, `minerPayout`
1. Miner claims Booking payment. New date of claim get's recorded. 
    - Functions: `claimPayment`
    - Events: `minerPayout`

Events are used to calculate miner's reputation.

### Service functions
1. Stablecoin address. Set is callable only by community contract/wallet.
    - Functions: `setStablecoinAddress`, `getStablecoinAddress`
1. Community wallet/contract address. Set is callable only by community contract/wallet. Just a regular wallet for now.
    - Functions: `setCommunityContract`, `getCommunityContract`
1. Community fee. 5% for now.
    - Functions: `setCommunityFee`, `getCommunityFee`
