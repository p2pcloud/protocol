package main

import (
	"github.com/sirupsen/logrus"

	"github.com/p2pcloud/protocol/implementations/evm/evmcompiler"
)

func main() {
	err := evmcompiler.CompileContracts(true)
	if err != nil {
		panic(err)
	}
	logrus.Println("Contracts compiled")
}
