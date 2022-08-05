package evm

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/implementations/evm/token"
	"github.com/p2pcloud/protocol/pkg/keyring"
)

type GanacheBCHelper struct {
	privateKeys  []*ecdsa.PrivateKey
	pksString    []string
	pksToString  map[*ecdsa.PrivateKey]string
	nextKeyIndex int
	waitForTx    func(common.Hash) error
}

func NewGanacheBCHelper(count int, web3client *ethclient.Client) (*GanacheBCHelper, error) {
	result := make([]*ecdsa.PrivateKey, 0, count)

	// ganache seed 123, mnemonic phrase:
	// tumble then poet spot sail spike forward system theory ankle pottery cute
	var pks = []string{
		"b751ca3c38c91528ff87cb04c98eb85f78693f8eaab5e45300270a9dc18168db",
		"af355f2914da0ea0361a5f93ca4224f60ae6a00d648e63f8fdebacc26e5a06f3",
		"db4aa7094e4a444cfe44667e57c2110ac5174d6736dc36ecb45165a6c583523b",
		"bcd820e1bff433173d7ad7d855ca6d44150c91d873faff98dadbf8435da5b6be",
		"4559d2a8327fa19c254256d79d1dcfa8168b1a6b567ff80062629b6aab94bdef",
		"462f146e90f4eb9006f382ba8d86af492f74f747c225979fce0d182b33a5b81c",
		"a99dd215ddb7fea877ec19a21d8a11f1e45d7a650a510f637af67018db159d7c",
		"46ada3ce889c3478d808bf47529b83da749a078f15794042fdc800af36f3340d",
		"bf11a242a33fb33fba0a54b8ecb032a2a5d43112c78be500433f4f28e75927eb",
		"f56760887faffc45797bb5d6563bcd4032254c44524fca8191fe6009c9cef5cd",
		"0f9e2f3a4afd287342731d483db789fb50aca9ea8aca7f763d5b640cb315acf7",
		"bea74505d50c0caee8ca4b44a157cc78538e512755b2146ae1f79fa1c730b87d",
		"6c01d43b67c7cde066dec41c58128baeab5fefcfe1c96103a31e54b7b19d2974",
		"a4081e80d738502c98ae491d2f334a61bdef105e9a11e2da851c5bda06f71ce3",
		"7a4a8ac13ac681cc2886ac0c2f3c18b4b723de9c12f58907d178775d8b4a6b8d",
		"130fdc39b5a044cf802e71c9cbac8a7bb937029da7b261ba7c94ebfba018587d",
		"b399f2a03189b9a16406ede65a38178bda58ceb81a3dcbebd91c419e018217fc",
		"c22e32d55a217ac6a3e52bde9e0ab83020dd6e68850f7a7e5df4d39bff2efbd6",
		"c75b72fec5ed55ff556dd6a95f1c64a03885ca6008eb52e1504f963423ca0fb4",
		"11a42ca61ca17d0b708e4cd2165ded2f84d8a8e866969adfe4b9f7644569c690",
	}

	pksToString := make(map[*ecdsa.PrivateKey]string, len(pks))

	for i := range pks {
		pk, err := keyring.DecodePrivateKey(pks[i])
		if err != nil {
			return nil, err
		}

		pksToString[pk] = pks[i]
		result = append(result, pk)
	}

	return &GanacheBCHelper{
		privateKeys: result,
		pksString:   pks,
		pksToString: pksToString,
		waitForTx:   NewTransactionWaiter(web3client, time.Minute).WaitForTx,
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

func (s *GanacheBCHelper) GetPrivateKeyString(pk *ecdsa.PrivateKey) string {
	return s.pksToString[pk]
}

func (s *GanacheBCHelper) WaitForTx(hash common.Hash) error {
	return s.waitForTx(hash)
}

type BlockChainHelper interface {
	GetNextPrivateKey() (*ecdsa.PrivateKey, error)
	WaitForTx(common.Hash) error
}

type TestInstances struct {
	Count         int
	Decimals      uint8
	InitialSupply float64
	UpdateCh      chan common.Address

	DeployerToken       *token.Token
	DeployerBroker      protocol.BrokerIface
	BrokerDeployAddress *common.Address
	CommunityInitialPk  *ecdsa.PrivateKey
	CommunityAccount    protocol.BrokerIface
	Contracts           []protocol.BrokerIface
	BcHelper            BlockChainHelper
	Backend             bind.ContractBackend
}

type Gifts struct {
	userInitialBalances map[int]float64
	userAllowances      map[int]float64
}

func NewGifts(userInitialBalances, userAllowances map[int]float64) *Gifts {
	return &Gifts{
		userInitialBalances: userInitialBalances,
		userAllowances:      userAllowances,
	}
}

func (g *Gifts) requiredSupply() (initialSupply float64) {
	if g == nil {
		return 0
	}

	for _, tokens := range g.userInitialBalances {
		initialSupply += tokens
	}

	return initialSupply
}

func (g *Gifts) InitialUserAllowances(p *TestInstances) error {
	if g == nil {
		return nil
	}

	for userIdx, tokens := range g.userAllowances {
		user := p.Contracts[userIdx]

		userToken := token.NewToken(&token.Params{
			Decimals:           &p.Decimals,
			Backend:            p.Backend,
			PrivateKey:         user.GetPrivateKey(),
			ContractAddressStr: p.DeployerToken.GetContractAddress().Hex(),
			ChainID:            ChainIDSimulated,
			WaitForTx:          p.BcHelper.WaitForTx,
		}).(*token.Token)

		if err := userToken.RegenerateSession(); err != nil {
			return err
		}

		if err := userToken.Approve(user.ContractAddress(), tokens); err != nil {
			return err
		}
	}

	return nil
}

func (g *Gifts) InitialUserTokens(p *TestInstances) error {
	if g == nil {
		return nil
	}

	tkn := token.NewToken(&token.Params{
		Decimals:           &p.Decimals,
		Backend:            p.Backend,
		PrivateKey:         p.DeployerToken.GetPrivateKey(),
		ContractAddressStr: p.DeployerToken.GetContractAddress().Hex(),
		ChainID:            ChainIDSimulated,
		WaitForTx:          p.BcHelper.WaitForTx,
	}).(*token.Token)
	if err := tkn.RegenerateSession(); err != nil {
		return err
	}

	for userIdx, tokens := range g.userInitialBalances {
		if err := tkn.Transfer(
			crypto.PubkeyToAddress(p.Contracts[userIdx].GetPrivateKey().PublicKey), tokens,
		); err != nil {
			return err
		}
	}

	return nil
}

func InitializeTestInstances(
	count int, decimals uint8, g *Gifts,
	backend bind.ContractBackend, bcHelper BlockChainHelper,
	communityPk *ecdsa.PrivateKey,
) (*TestInstances, error) {
	p := &TestInstances{
		Count:              count,
		Decimals:           decimals,
		InitialSupply:      g.requiredSupply() + 10,
		Backend:            backend,
		BcHelper:           bcHelper,
		CommunityInitialPk: communityPk,
		UpdateCh:           make(chan common.Address, count),
	}

	if err := BuildToken(p); err != nil {
		return nil, err
	}

	if err := BuildBroker(p); err != nil {
		return nil, err
	}

	if err := p.DeployerBroker.SetStablecoinAddress(p.DeployerToken.GetContractAddress()); err != nil {
		return nil, err
	}

	if err := BuildUsers(p); err != nil {
		return nil, err
	}

	if err := g.InitialUserTokens(p); err != nil {
		return nil, err
	}

	if err := g.InitialUserAllowances(p); err != nil {
		return nil, err
	}

	return p, nil
}

func BuildToken(p *TestInstances) error {
	tokenPk, err := p.BcHelper.GetNextPrivateKey()
	if err != nil {
		return err
	}

	tkn := token.NewToken(&token.Params{
		Decimals:   &p.Decimals,
		Backend:    p.Backend,
		PrivateKey: tokenPk,
		ChainID:    ChainIDSimulated,
		WaitForTx:  p.BcHelper.WaitForTx,
		UpdCh:      p.UpdateCh,
	}).(*token.Token)

	if _, err = tkn.DeployContract(p.InitialSupply); err != nil {
		return err
	}

	p.DeployerToken = tkn

	return nil
}

func BuildBroker(p *TestInstances) error {
	pk, err := p.BcHelper.GetNextPrivateKey()
	if err != nil {
		return err
	}

	deployContract, err := broker.NewBroker(
		p.Backend, pk, "", ChainIDSimulated, p.BcHelper.WaitForTx, p.UpdateCh,
	)
	if err != nil {
		return err
	}

	_, err = deployContract.DeployContracts(crypto.PubkeyToAddress(p.CommunityInitialPk.PublicKey))

	if err != nil {
		return err
	}

	p.DeployerBroker = deployContract

	return nil
}

func BuildUsers(p *TestInstances) error {
	userContracts := make([]protocol.BrokerIface, 0)
	for i := 0; i < p.Count; i++ {
		pk, err := p.BcHelper.GetNextPrivateKey()
		if err != nil {
			return err
		}

		userContract, err := broker.NewBroker(
			p.Backend, pk, p.DeployerBroker.ContractAddress().Hex(),
			ChainIDSimulated, p.BcHelper.WaitForTx, p.UpdateCh,
		)
		if err != nil {
			return err
		}

		userContracts = append(userContracts, userContract)
	}

	p.Contracts = userContracts

	return nil
}

func SetFeeAndStablecoin(p *TestInstances) error {
	communityContract, err := broker.NewBroker(
		p.Backend, p.CommunityInitialPk, p.DeployerBroker.ContractAddress().Hex(),
		ChainIDSimulated, p.BcHelper.WaitForTx, make(chan common.Address, 1),
	)
	if err != nil {
		return err
	}

	if err = communityContract.(*broker.Broker).RegenerateSession(); err != nil {
		return err
	}

	if err = communityContract.SetCommunityFee(CommunityFee); err != nil {
		return err
	}

	if err = communityContract.SetStablecoinAddress(p.DeployerToken.GetContractAddress()); err != nil {
		return err
	}

	p.CommunityAccount = communityContract

	return nil
}
