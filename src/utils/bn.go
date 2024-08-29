package utils

import (
	"fmt"
	"math/big"
)

func BnInt64ToBytes(input int64) []byte {
	value := new(big.Int)
	value.SetInt64(input)
	return value.Bytes()
}

func BnStringToBytes(input string) ([]byte, error) {
	value := new(big.Int)
	_, ok := value.SetString(input, 10)
	if !ok {
		return nil, fmt.Errorf("not big number string")
	}
	return value.Bytes(), nil
}

func BytesToBnInt64(input []byte) *big.Int {
	value := new(big.Int)
	value.SetBytes(input)
	return value
}
