# Blockchain logic for P2PCloud

![Badge](https://github.com/p2pcloud/protocol/actions/workflows/go.yml/badge.svg)


## What is p2pcloud
[p2pcloud.io](https://p2pcloud.io) is a protocol and utility set for booking and securing trustless virtual machines over the blockchain.

## Development 
### Agreements
All implementations should be in the implementations folder in a separate subfolder. 
All implementations should comply with BlockchainIface
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
