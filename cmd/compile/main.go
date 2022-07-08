package main

import (
	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/sirupsen/logrus"
)

func main() {
	err := evm.CompileContracts()
	if err != nil {
		panic(err)
	}
	logrus.Println("Contracts compiled")
}
