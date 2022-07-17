package evm

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/p2pcloud/protocol/pkg/keyring"
	"strings"
	"testing"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/implementations/evm/token"
)

type WrappedSimulatedBlockchainEnv struct {
	Origin *SimulatedBlockchainEnv
}

func (b *WrappedSimulatedBlockchainEnv) GetNextPrivateKey() (*ecdsa.PrivateKey, error) {
	return b.Origin.GetNextPrivateKey()
}

func (b *WrappedSimulatedBlockchainEnv) Commit() {
	b.Origin.Backend.Commit()
}

func NewWrappedSimulatedBlockchainEnv(t *testing.T) *WrappedSimulatedBlockchainEnv {
	bc, err := NewSimulatedBlockchainEnv()
	if err != nil {
		t.Fatal(err)
	}

	return &WrappedSimulatedBlockchainEnv{Origin: bc}
}

type GanacheBCHelper struct {
	privateKeys  []*ecdsa.PrivateKey
	nextKeyIndex int
}

func NewGanacheBCHelper(count int, jsonKeys string) (*GanacheBCHelper, error) {
	var pks []string
	if err := json.NewDecoder(strings.NewReader(jsonKeys)).Decode(&pks); err != nil {
		return nil, err
	}

	result := make([]*ecdsa.PrivateKey, 0, count)
	for i := range pks {
		pk, err := keyring.DecodePrivateKey(pks[i])
		if err != nil {
			return nil, err
		}

		result = append(result, pk)
	}

	return &GanacheBCHelper{
		privateKeys: result,
	}, nil
}

func (s *GanacheBCHelper) GetNextPrivateKey() (*ecdsa.PrivateKey, error) {
	if s.nextKeyIndex >= len(s.privateKeys) {
		return nil, fmt.Errorf("no more private keys to test, got %d", len(s.privateKeys))
	}
	key := s.privateKeys[s.nextKeyIndex]
	s.nextKeyIndex++
	return key, nil
}

func (s *GanacheBCHelper) Commit() {}

type BlockChainHelper interface {
	GetNextPrivateKey() (*ecdsa.PrivateKey, error)
	Commit()
}

type TestInstances struct {
	Count         int
	Decimals      int64
	InitialSupply int64

	DeployerToken       *token.Token
	BrokerDeployAddress *common.Address
	Contracts           []protocol.BrokerIface
	BcHelper            BlockChainHelper
	Backend             bind.ContractBackend
}

type Gifts struct {
	userInitialBalances map[int]int64
	userAllowances      map[int]int64
}

func NewGifts(userInitialBalances, userAllowances map[int]int64) *Gifts {
	return &Gifts{
		userInitialBalances: userInitialBalances,
		userAllowances:      userAllowances,
	}
}

func (g *Gifts) requiredSupply() (initialSupply int64) {
	for _, tokens := range g.userInitialBalances {
		initialSupply += tokens
	}

	return initialSupply
}

func (g *Gifts) initialUserAllowances(p *TestInstances) error {
	for userIdx, tokens := range g.userAllowances {
		user := p.Contracts[userIdx]

		userToken := token.NewToken(&token.Params{
			Decimals:           p.Decimals,
			Backend:            p.Backend,
			PrivateKey:         user.GetPrivateKey(),
			ContractAddressStr: p.DeployerToken.GetContractAddress().Hex(),
			ChainID:            ChainIDSimulated,
			Commit:             p.BcHelper.Commit,
		})

		if err := userToken.RegenerateSession(); err != nil {
			return err
		}

		if err := userToken.TestApprove(user.ContractAddress(), tokens); err != nil {
			return err
		}
	}

	return nil
}

func (g *Gifts) initialUsersTokens(p *TestInstances) error {
	for userIdx, tokens := range g.userInitialBalances {
		userPk := p.Contracts[userIdx].GetPrivateKey()

		err := p.DeployerToken.TransferForTests(
			crypto.PubkeyToAddress(p.DeployerToken.GetPrivateKey().PublicKey),
			crypto.PubkeyToAddress(userPk.PublicKey),
			tokens,
		)
		if err != nil {
			return err
		}

	}

	return nil
}

func InitializeTestInstances(
	count int, decimals int64, g Gifts,
	backend bind.ContractBackend, bcHelper BlockChainHelper,
) (*TestInstances, error) {
	p := &TestInstances{
		Count:         count,
		Decimals:      decimals,
		InitialSupply: g.requiredSupply() + 10,
		Backend:       backend,
		BcHelper:      bcHelper,
	}

	if err := instantiateToken(p); err != nil {
		return nil, err
	}

	if err := instantiateUsers(p); err != nil {
		return nil, err
	}

	if err := g.initialUsersTokens(p); err != nil {
		return nil, err
	}

	if err := g.initialUserAllowances(p); err != nil {
		return nil, err
	}

	return p, nil
}

func instantiateToken(p *TestInstances) error {
	tokenPk, err := p.BcHelper.GetNextPrivateKey()
	if err != nil {
		return err
	}

	tkn := token.NewToken(&token.Params{
		Decimals:   p.Decimals,
		Backend:    p.Backend,
		PrivateKey: tokenPk,
		ChainID:    ChainIDSimulated,
		Commit:     p.BcHelper.Commit,
	})
	if err != nil {
		return err
	}

	_, err = tkn.DeployContract(p.InitialSupply)
	if err != nil {
		return err
	}

	p.DeployerToken = tkn

	return nil
}

func instantiateUsers(p *TestInstances) error {
	pk, err := p.BcHelper.GetNextPrivateKey()
	if err != nil {
		return err
	}

	deployContract, err := broker.NewBroker(
		p.Backend, pk, "", ChainIDSimulated, p.DeployerToken.GetContractAddress(),
	)
	if err != nil {
		return err
	}

	addresses, err := deployContract.DeployContracts()
	if err != nil {
		return err
	}
	contractAddressStr := addresses[0]
	contractAddress := common.HexToAddress(contractAddressStr)

	p.BcHelper.Commit()

	userContracts := make([]protocol.BrokerIface, 0)
	for i := 0; i < p.Count; i++ {
		pk, err = p.BcHelper.GetNextPrivateKey()
		if err != nil {
			return err
		}

		userContract, err := broker.NewBroker(
			p.Backend, pk, contractAddressStr, ChainIDSimulated, p.DeployerToken.GetContractAddress(),
		)
		if err != nil {
			return err
		}

		userContracts = append(userContracts, userContract)
	}

	p.Contracts = userContracts
	p.BrokerDeployAddress = &contractAddress

	return nil
}
