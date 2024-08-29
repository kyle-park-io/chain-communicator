package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"chain-communicator/logger"
)

type Contract struct {
	Abi      abi.ABI
	Bytecode string
}

func GetContractInfo(contract string) (*Contract, error) {

	var contract2 string
	if strings.Contains(contract, "token") {
		contract2 = "TokenBurnable"
	} else {
		contract2 = contract
	}

	file, err := os.Open(fmt.Sprintf("artifacts/%s.json", contract2))
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	defer file.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	abiJSON, err := json.Marshal(data["abi"])
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	parsedABI, err := abi.JSON(bytes.NewReader(abiJSON))
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	bytecode, ok := data["bytecode"].(string)
	if !ok {
		return nil, fmt.Errorf("convert error")
	}

	return &Contract{Abi: parsedABI, Bytecode: bytecode}, nil
}

func GetContractAddress(contract string) (string, error) {

	file, err := os.Open(fmt.Sprintf("contracts/%s.contract.json", contract))
	if err != nil {
		logger.Log.Error(err)
		return "", err
	}
	defer file.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		logger.Log.Error(err)
		return "", err
	}

	address, ok := data["contractAddress"].(string)
	if !ok {
		return "", fmt.Errorf("convert error")
	}
	return address, nil
}
