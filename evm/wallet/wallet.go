package wallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/sirupsen/logrus"
)

// var web3Client *ethclient.Client = nil

type ethAnyClient interface {
	BalanceAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (*big.Int, error)
	bind.ContractBackend
}

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Backend    ethAnyClient
	chainId    *big.Int
}

func NewWalletFromPK(pk *ecdsa.PrivateKey, backend ethAnyClient, chanId int64) Wallet {
	return Wallet{PrivateKey: pk, Backend: backend, chainId: big.NewInt(chanId)}
}

func weiToEther(wei *big.Int) float64 {
	result, _ := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether)).Float64()
	return result
}

func etherToWei(ethFloat float64) *big.Int {
	eth := big.NewFloat(ethFloat)
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (*common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, errors.New("error casting public key to ECDSA, probably wrong format")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &address, nil
}

func (w *Wallet) GetBalanceETH() (float64, error) {
	balanceWei, err := w.GetBalanceWei()
	if err != nil {
		return 0, err
	}
	return weiToEther(balanceWei), err
}

func (w *Wallet) GetBalanceWei() (*big.Int, error) {
	address, err := PrivateKeyToAddress(w.PrivateKey)
	if err != nil {
		return nil, err
	}

	return w.Backend.BalanceAt(context.Background(), *address, nil)
}

func (w *Wallet) TransferEth(to common.Address, amount float64) error {
	amountWei := etherToWei(amount)
	return w.TransferWei(to, amountWei)
}

func (w *Wallet) TransferWei(to common.Address, amount *big.Int) error {
	fromAddress := crypto.PubkeyToAddress(w.PrivateKey.PublicKey)
	nonce, err := w.Backend.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logrus.Fatal(err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := w.Backend.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(w.chainId), w.PrivateKey)
	if err != nil {
		return err
	}

	return w.Backend.SendTransaction(context.Background(), signedTx)
}

func (w *Wallet) GetAddress() (*common.Address, error) {
	return PrivateKeyToAddress(w.PrivateKey)
}
