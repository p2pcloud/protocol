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
    - Functions: `getMinersOffers`, `updateOffer`, `addOffer`
1. User sends a stablecoin to the Broker. All stablecoin is split between free and locked states.
    - Tasks: [Deposit and withdraw steblecoin](https://github.com/p2pcloud/protocol/issues/1), [Check stablecoin balance](https://github.com/p2pcloud/protocol/issues/2)
    - Functions: `depositCoin`, `checkBalance`, `withdrawCoin`
1. User gets list of offers
    - Functions: `getAvailableOffers`
1. User books VM. His stablecoin gets locked. 
    - Functions: `bookVM`
    - Tasks: [bookVM: Lock pps*time of user's stablecoin and check if user has enough](https://github.com/p2pcloud/protocol/issues/3), [Emit bookingStarted event](https://github.com/p2pcloud/protocol/issues/4)
    - Tasks: [Emit booking event with bookingId, timeUsed, miner, user and vmType fields](https://github.com/p2pcloud/protocol/issues/5)
1. User aborts booking because of miners misbehaviour. Unused stablecoin gets unlocked. Miner gets paid 1/2 of the price. The rest 1/2 goes into the community. Booking gets deleted.
    - Functions: `reportBooking` 
    - Events: `bookingReported`, `minerPayout`
    - Tasks: [Emit bookingReported event with bookingId, timeUsed, miner, user and vmType fields](https://github.com/p2pcloud/protocol/issues/6)
1. User aborts booking, but no problems with miner. Unused stablecoin gets unlocked. Miner gets paid 95% of the price. The rest 5% goes into the community. Booking gets deleted.
    - Functions: `stopBooking`
    - Events: `bookingStopped`, `minerPayout`
    - Tasks: [Emit bookingStopped event with bookingId, timeUsed, miner, user and vmType fields](https://github.com/p2pcloud/protocol/issues/7)
1. User extends booking. More stablecoin gets locked. New PPS is used. No checks for slot availability.
    - Functions: `extendBooking`
    - Tasks: [Emit bookingExtended event with bookingId, timeUsed, miner, user and vmType fields](https://github.com/p2pcloud/protocol/issues/8) [Propose logic of refund in case if extended with different pps booking gets aborted](https://github.com/p2pcloud/protocol/issues/9)
1. User does nothing and booking expires.
1. Miner claims expired booking. Miner gets paid 95% of the price. The rest 5% goes into the community. Booking gets deleted.
    - Functions: `claimBookingFinished`
    - Events: `minerPayout`
    - Tasks: [claimBookingFinished function](https://github.com/p2pcloud/protocol/issues/10)

Events are used to calculate miner's reputation.

### Service functions
1. Stablecoin address. Set is callable only by community contract/wallet.
    - Functions: `setStablecoinAddress`, `getStablecoinAddress`
    - Tasks: [Set and get stablecoin address](https://github.com/p2pcloud/protocol/issues/11) [Propose logic of work in case stablecoin address is changed during active vm bookings](https://github.com/p2pcloud/protocol/issues/12)
1. Community wallet/contract address. Set is callable only by community contract/wallet. Just a regular wallet for now.
    - Functions: `setCommunityContract`, `getCommunityContract`
    - Tasks: [Set and get community address](https://github.com/p2pcloud/protocol/issues/13) 
1. Community fee. 5% for now.
    - Functions: `setCommunityFee`, `getCommunityFee`
    - Tasks: [Set and get community fee](https://github.com/p2pcloud/protocol/issues/14) 
