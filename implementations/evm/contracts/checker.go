package contracts

// func (session *BrokerTransactorSession) DebugEstimateGas(backend bind.ContractBackend, contractAddress string, myAddress string, method string, args ...interface{}) (uint64, error) {
// 	abi, err := abi.JSON(strings.NewReader(BrokerABI))
// 	if err != nil {
// 		return 0, err
// 	}

// 	input, err := abi.Pack(method, args...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var b bind.ContractBackend

// 	sendTo := common.HexToAddress(contractAddress)

// 	callMsg := ethereum.CallMsg{
// 		From:     common.HexToAddress(contractAddress),
// 		To:       &sendTo,
// 		Gas:      0,
// 		GasPrice: big.NewInt(0),
// 		Value:    big.NewInt(0),
// 		Data:     input,
// 	}

// 	gasLimit, err := backend.EstimateGas(context.Background(), callMsg)
// 	return gasLimit, err
// }
