package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"chain-communicator/constant"
	"chain-communicator/encode"
	"chain-communicator/logger"
	"chain-communicator/marshal"
	"chain-communicator/rpc"
	"chain-communicator/utils"
)

func Transfer(fromPriv string, to string, amount string) error {

	from, err := utils.PrivteToAddress(fromPriv)
	if err != nil {
		return err
	}

	// rpc - account
	data, err := rpc.GetAccount(from)
	if err != nil {
		return err
	}

	var parsedAccount map[string]interface{}
	err = json.Unmarshal([]byte(data), &parsedAccount)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	result := parsedAccount["result"].(map[string]interface{})
	value := result["value"].(map[string]interface{})

	// rpc - rule
	data2, err := rpc.GetRule()
	if err != nil {
		return err
	}

	var parsedRule map[string]interface{}
	err = json.Unmarshal([]byte(data2), &parsedRule)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	result2 := parsedRule["result"].(map[string]interface{})
	value2 := result2["value"].(map[string]interface{})

	// nonce
	nS, ok := value["nonce"].(string)
	if !ok {
		return fmt.Errorf("nonce is not string")
	}
	// gas
	gS, ok := value2["maxTrxGas"].(string)
	if !ok {
		return fmt.Errorf("gas is not string")
	}
	// gasPrice
	gPS, ok := value2["gasPrice"].(string)
	if !ok {
		return fmt.Errorf("gasPrice is not string")
	}

	// build Transaction
	pb, err := utils.BuildTransaction(nS, from, to, amount, gS, gPS, 1, []byte{})
	if err != nil {
		return err
	}

	// encode
	eR, err := encode.EncodeRlp(pb, []byte{})
	if err != nil {
		return err
	}

	// prefix
	p1 := strings.Replace(constant.TxPrefix, "<SetYourChainId>", constant.Test_ChainId, 1)
	p2 := strings.Replace(p1, "<SetYourEncodedDataLength>", fmt.Sprint(len(eR)), 1)
	prefixBytes := []byte(p2)
	concat := append(prefixBytes, eR...)

	// sign
	sign, err := utils.Sign(fromPriv, concat)
	if err != nil {
		return err
	}
	pb.Sig = sign

	t, err := marshal.ProtoMarshal(pb)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	tS := base64.StdEncoding.EncodeToString(t)

	transfer, err := rpc.Transfer(tS)
	if err != nil {
		return err
	}
	logger.Log.Info(transfer)

	return nil
}
