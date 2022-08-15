package evmcompiler

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

const container_solc = "ethereum/solc:0.8.15"
const container_geth_tools = "ethereum/client-go:alltools-release-1.10"
const container_basic = "debian:buster"

func CompileContracts(compileTestToken bool) error {
	contractsDir, err := filepath.Abs("./implementations/evm/contracts")
	if err != nil {
		return err
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/contracts", contractsDir),
		container_solc,
		"-o", "/contracts", "--abi", "--bin", "/contracts/Broker.sol", "--overwrite",
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error compiling: %v: %s", err, out)
	}

	cmd = exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/contracts", contractsDir),
		container_geth_tools,
		"abigen", "--bin=/contracts/Broker.bin", "--type", "Broker", "--abi=/contracts/Broker.abi", "--pkg=contracts", "--out=/contracts/broker_autogenerated.go",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error abigen: %v: %s", err, out)
	}

	if err = compileERC(contractsDir, compileTestToken); err != nil {
		return err
	}

	cmd = exec.Command(
		"docker", "run", "--rm",
		container_solc,
		"--version",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error getting version: %v: %s", err, out)
	}
	ioutil.WriteFile(filepath.Join(contractsDir, "compiler_version.txt"), []byte(out), 0666)

	cmd = exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/contracts", contractsDir),
		container_basic,
		"bash", "-c", "rm /contracts/*.bin && chmod 777 /contracts/*.go",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error abigen: %v: %s", err, out)
	}

	return nil
}

func compileERC(contractsDir string, should bool) error {
	if !should {
		return nil
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/contracts", contractsDir),
		container_solc,
		"-o", "/contracts", "--abi", "--bin", "/contracts/Token.sol", "--overwrite",
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error compiling: %v: %s", err, out)
	}

	cmd = exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/contracts", contractsDir),
		container_geth_tools,
		"abigen", "--bin=/contracts/Token.bin", "--type", "Token", "--abi=/contracts/Token.abi", "--pkg=contracts", "--out=/contracts/token_autogenerated.go",
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error abigen: %v: %s", err, out)
	}

	return nil
}
