package wallet

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"

	"github.com/p2pcloud/protocol/implementations/evm/contracts"
)

type Token struct {
	backend         ethAnyClient
	transactOpts    *bind.TransactOpts
	contractAddress common.Address
	session         contracts.FiatTokenSession
	privateKey      *ecdsa.PrivateKey
	chainID         int64
	gasLimit        uint64
}

func NewToken(
	backend ethAnyClient, privateKey *ecdsa.PrivateKey,
	contractAddressStr string, chainId int64, gasLimit uint64,
) (*Token, error) {
	if chainId == 0 {
		return nil, fmt.Errorf("chanId is 0. please set it to a valid value")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}

	return &Token{
		backend:         backend,
		transactOpts:    transactOpts,
		contractAddress: common.HexToAddress(contractAddressStr),
		privateKey:      privateKey,
		chainID:         chainId,
		gasLimit:        gasLimit,
	}, nil
}

func (b *Token) GetPrivateKey() *ecdsa.PrivateKey {
	return b.privateKey
}

func (b *Token) RegenerateSession() error {
	instance, err := contracts.NewFiatToken(b.contractAddress, b.backend)
	if err != nil {
		return err
	}

	b.session = contracts.FiatTokenSession{
		Contract:     instance,
		TransactOpts: *b.transactOpts,
		CallOpts: bind.CallOpts{
			Pending: false,               // Whether to operate on the pending state or the last known one
			From:    b.transactOpts.From, // Optional the sender address, otherwise the first account is used
			Context: context.Background(),
		},
	}
	return nil
}

func (b *Token) Deploy() (string, error) {
	address, _, _, err := contracts.DeployFiatToken(
		b.transactOpts,
		b.backend,
	)

	if err != nil {
		return "", fmt.Errorf("could not deploy fiat token: %v", err)
	}
	b.contractAddress = address

	b.RegenerateSession()

	return address.String(), nil
}

func (b *Token) GetMyAddress() *common.Address {
	return &b.transactOpts.From
}

func (b *Token) TransferTokens(ctx context.Context, from, to common.Address, amount uint64) error {
	b.RegenerateSession()

	var (
		transferFnSignature = []byte("transfer(address,uint64)")
	)

	nonce, err := b.backend.PendingNonceAt(ctx, from)
	if err != nil {
		logrus.Fatal(err)
	}

	gasPrice, err := b.backend.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	intArr := make([]byte, 8)
	binary.LittleEndian.PutUint64(intArr, amount)

	paddedAddress := common.LeftPadBytes(to.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(intArr, 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit := b.gasLimit
	if gasLimit < 1 {
		gasLimit, err = b.backend.EstimateGas(ctx, ethereum.CallMsg{
			To:   &b.contractAddress,
			Data: data,
		})
	}

	tx := types.NewTransaction(nonce, b.contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(b.chainID)), b.privateKey)
	if err != nil {
		return err
	}

	return b.backend.SendTransaction(ctx, signedTx)
}

func (b *Token) BalanceTokens() (uint64, error) {
	b.RegenerateSession()

	return b.session.BalanceOf(*b.GetMyAddress())
}

func (b *Token) GetBalanceWei() (*big.Int, error) {
	return b.backend.BalanceAt(context.Background(), *b.GetMyAddress(), nil)
}
