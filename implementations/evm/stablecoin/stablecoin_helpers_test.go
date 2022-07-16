package stablecoin_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/implementations/evm/token"
)

const (
	ChainIDSimulated = 1337
)

type params struct {
	count         int
	decimals      int64
	initialSupply int64

	deployerToken       *token.Token
	brokerDeployAddress *common.Address
	contracts           []protocol.BrokerIface
	bcSim               *evm.SimulatedBlockchainEnv
}

type gifts struct {
	userInitialBalances map[int]int64
	userAllowances      map[int]int64
}

func (g *gifts) requiredSupply() (initialSupply int64) {
	for _, tokens := range g.userInitialBalances {
		initialSupply += tokens
	}

	return initialSupply
}

func (g *gifts) initialUserAllowances(p *params) error {
	for userIdx, tokens := range g.userAllowances {
		user := p.contracts[userIdx]

		userToken := token.NewToken(&token.Params{
			Decimals:           p.decimals,
			Backend:            p.bcSim.Backend,
			PrivateKey:         user.GetPrivateKey(),
			ContractAddressStr: p.deployerToken.GetContractAddress().Hex(),
			ChainID:            ChainIDSimulated,
			Commit:             p.bcSim.Backend.Commit,
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

func (g *gifts) initialUsersTokens(p *params) error {
	for userIdx, tokens := range g.userInitialBalances {
		userPk := p.contracts[userIdx].GetPrivateKey()

		err := p.deployerToken.TransferForTests(
			crypto.PubkeyToAddress(p.deployerToken.GetPrivateKey().PublicKey),
			crypto.PubkeyToAddress(userPk.PublicKey),
			tokens,
		)
		if err != nil {
			return err
		}

	}

	return nil
}

func instantiate(count int, decimals int64, g gifts) (*params, error) {
	p := &params{
		count:         count,
		decimals:      decimals,
		initialSupply: g.requiredSupply() + 10,
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

func instantiateToken(p *params) error {
	blockchainSim, err := evm.NewSimulatedBlockchainEnv()
	if err != nil {
		return err
	}

	tokenPk, err := blockchainSim.GetNextPrivateKey()
	if err != nil {
		return err
	}

	tkn := token.NewToken(&token.Params{
		Decimals:   p.decimals,
		Backend:    blockchainSim.Backend,
		PrivateKey: tokenPk,
		ChainID:    ChainIDSimulated,
		Commit:     blockchainSim.Backend.Commit,
	})
	if err != nil {
		return err
	}

	_, err = tkn.DeployContract(p.initialSupply)
	if err != nil {
		return err
	}

	blockchainSim.Backend.Commit()

	p.bcSim = blockchainSim
	p.deployerToken = tkn

	return nil
}

func instantiateUsers(p *params) error {
	pk, err := p.bcSim.GetNextPrivateKey()
	if err != nil {
		return err
	}

	deployContract, err := broker.NewBroker(
		p.bcSim.Backend, pk, "", ChainIDSimulated, p.deployerToken.GetContractAddress(),
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

	p.bcSim.Backend.Commit()

	userContracts := make([]protocol.BrokerIface, 0)
	for i := 0; i < p.count; i++ {
		pk, err = p.bcSim.GetNextPrivateKey()
		if err != nil {
			return err
		}

		userContract, err := broker.NewBroker(
			p.bcSim.Backend, pk, contractAddressStr, ChainIDSimulated, p.deployerToken.GetContractAddress(),
		)
		if err != nil {
			return err
		}

		userContracts = append(userContracts, userContract)
	}

	p.contracts = userContracts
	p.brokerDeployAddress = &contractAddress

	return nil
}
