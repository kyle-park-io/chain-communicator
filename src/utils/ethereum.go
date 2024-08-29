package utils

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"

	"chain-communicator/logger"
)

func EncodeFunctionSignature(method abi.Method) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(method.Sig))
	signatureHash := hex.EncodeToString(hash.Sum(nil))

	fS, err := HexStringToBytes("0x" + signatureHash[:8])
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return fS, nil
}
