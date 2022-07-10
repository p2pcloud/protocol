package broker

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"math/big"
)

func (b *Broker) GetCommunityWallet() (string, error) {
	hashBytes, err := b.session.GetCommunityWallet()
	if err != nil {
		return "", err
	}
	logrus.WithField("result", "0x"+hex.EncodeToString(hashBytes[:])).Debug("GetCommunityWallet called")
	return "0x" + hex.EncodeToString(hashBytes[:]), nil
}

func (b *Broker) SetCommunityWalletIfNeeded(address string) error {
	logrus.WithField("communityWallet", address).Debug("SetCommunityWalletIIfNeeded called")
	oldCommunityWallet, err := b.GetCommunityWallet()
	if err != nil {
		return err
	}

	if oldCommunityWallet == address {
		return nil
	}

	var bytes20 [20]byte

	hashBytes := common.FromHex(address)

	if len(hashBytes) != 20 {
		return fmt.Errorf("community wallet is not 20 bytes long")
	}

	copy(bytes20[:], hashBytes)

	_, err = b.session.SetCommunityWallet(bytes20)
	if err != nil {
		return fmt.Errorf("could not register community wallet: %v", err)
	}
	return nil
}

func (b *Broker) GetCommunityFee() (*big.Int, error) {
	fee, err := b.session.GetCommunityFee()
	if err != nil {
		return nil, err
	}
	logrus.WithField("result", fee).Debug("GetCommunityWallet called")
	return fee, nil
}

func (b *Broker) SetCommunityFeeIfNeeded(fee int64) error {
	bigFee := big.NewInt(fee)
	logrus.WithField("communityFee", fee).Debug("SetCommunityFeeIIfNeeded called")
	oldCommunityWallet, err := b.GetCommunityFee()
	if err != nil {
		return err
	}

	if oldCommunityWallet == bigFee {
		return nil
	}

	_, err = b.session.SetCommunityFee(bigFee)
	if err != nil {
		return fmt.Errorf("could not register community fee: %v", err)
	}
	return nil
}
