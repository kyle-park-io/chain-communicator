package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"

	"chain-communicator/logger"
)

func PrivteToAddress(input string) (string, error) {
	if !IsHexPriv(input) {
		return "", fmt.Errorf("not hex private key")
	}

	privateKey, err := crypto.HexToECDSA(input[2:])
	if err != nil {
		logger.Log.Error(err)
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("convert public key error")
	}
	publicKeyBytes := crypto.CompressPubkey(publicKeyECDSA)

	return PubToAddress(publicKeyBytes), nil
}

func PubToAddress(publicKey []byte) string {
	pubKeyHash := sha256.Sum256(publicKey)

	rip := ripemd160.New()
	rip.Write(pubKeyHash[:])
	ripHash := rip.Sum(nil)
	address := hex.EncodeToString(ripHash)

	return "0x" + address
}
