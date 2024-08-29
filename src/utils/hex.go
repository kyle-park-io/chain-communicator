package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"chain-communicator/logger"
)

func IsHexString(input string) bool {
	lower := strings.ToLower(input)
	if !(len(input) >= 2 && input[0] == '0' && input[1] == 'x') {
		return false
	}

	lower = lower[2:]
	_, err := hex.DecodeString(lower)
	if err != nil {
		logger.Log.Error(err)
		return false
	}
	return true
}

func IsHexAddress(input string) bool {
	if !IsHexString(input) {
		return false
	}
	if len(input) == 42 {
		return true
	}
	return false
}

func IsHexPriv(input string) bool {
	if !IsHexString(input) {
		return false
	}
	if len(input) == 66 {
		return true
	}
	return false
}

func HexStringToBytes(input string) ([]byte, error) {
	if !IsHexString(input) {
		return nil, fmt.Errorf("not hex string")
	}

	bS, err := hex.DecodeString(input[2:])
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return bS, nil
}

func BytesToHexString(input []byte) string {
	return "0x" + hex.EncodeToString(input)
}

func HexAddressToBytes(input string) ([]byte, error) {
	if !IsHexAddress(input) {
		return nil, fmt.Errorf("not hex address")
	}

	bA, err := hex.DecodeString(input[2:])
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return bA, nil
}

func BytesToHexAdress(input []byte) (string, error) {
	if len(input) != 20 {
		return "", fmt.Errorf("length is not 20")
	}
	return "0x" + hex.EncodeToString(input), nil
}
