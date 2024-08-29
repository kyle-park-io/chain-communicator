package utils

import (
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/crypto"

	"chain-communicator/logger"
)

func Sign(fromPriv string, input []byte) ([]byte, error) {
	pB, err := HexStringToBytes(fromPriv)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	priv, err := crypto.ToECDSA(pB)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	hashBytes := sha256.Sum256(input)
	signature, err := crypto.Sign(hashBytes[:], priv)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	return signature, nil
}
