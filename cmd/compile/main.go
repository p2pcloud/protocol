package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Incognida/protocol/implementations/evm"
)

func main() {
	err := evm.CompileContracts()
	if err != nil {
		panic(err)
	}
	logrus.Println("Contracts compiled")
}
