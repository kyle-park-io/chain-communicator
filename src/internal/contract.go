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

func QueryTx(fromPriv string, contract string, function string, args []interface{}) ([]interface{}, error) {

	cI, err := utils.GetContractInfo(contract)
	if err != nil {
		return nil, err
	}

	cA, err := utils.GetContractAddress(contract)
	if err != nil {
		return nil, err
	}

	from, err := utils.PrivteToAddress(fromPriv)
	if err != nil {
		return nil, err
	}

	// encode parameters
	method := cI.Abi.Methods[function]
	fS, err := utils.EncodeFunctionSignature(method)
	if err != nil {
		return nil, err
	}
	eP, err := method.Inputs.Pack(args...)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	payload := append(fS, eP...)

	tS := base64.StdEncoding.EncodeToString(payload)

	call, err := rpc.QueryTx(from, cA, tS)
	if err != nil {
		return nil, err
	}

	var parsedCall map[string]interface{}
	err = json.Unmarshal([]byte(call), &parsedCall)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	result3 := parsedCall["result"].(map[string]interface{})
	value3 := result3["value"].(map[string]interface{})
	rD, ok := value3["returnData"].(string)
	if !ok {
		return nil, fmt.Errorf("returnData is not string")
	}
	rDB, err := base64.StdEncoding.DecodeString(rD)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	k, err := method.Outputs.Unpack(rDB)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	return k, nil
}

func QueryTxFromAddress(fromPriv string, contract string, function string, args []interface{}, address string) ([]interface{}, error) {

	cI, err := utils.GetContractInfo(contract)
	if err != nil {
		return nil, err
	}

	cA := address

	from, err := utils.PrivteToAddress(fromPriv)
	if err != nil {
		return nil, err
	}

	// encode parameters
	method := cI.Abi.Methods[function]
	fS, err := utils.EncodeFunctionSignature(method)
	if err != nil {
		return nil, err
	}
	eP, err := method.Inputs.Pack(args...)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	payload := append(fS, eP...)

	tS := base64.StdEncoding.EncodeToString(payload)

	call, err := rpc.QueryTx(from, cA, tS)
	if err != nil {
		return nil, err
	}

	var parsedCall map[string]interface{}
	err = json.Unmarshal([]byte(call), &parsedCall)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	result3 := parsedCall["result"].(map[string]interface{})
	value3 := result3["value"].(map[string]interface{})
	rD, ok := value3["returnData"].(string)
	if !ok {
		return nil, fmt.Errorf("returnData is not string")
	}
	rDB, err := base64.StdEncoding.DecodeString(rD)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	k, err := method.Outputs.Unpack(rDB)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	return k, nil
}

func SubmitTxWithETH(fromPriv string, eth string, contract string, function string, args []interface{}) error {

	cI, err := utils.GetContractInfo(contract)
	if err != nil {
		return err
	}

	cA, err := utils.GetContractAddress(contract)
	if err != nil {
		return err
	}

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

	// encode parameters
	method := cI.Abi.Methods[function]
	fS, err := utils.EncodeFunctionSignature(method)
	if err != nil {
		return err
	}
	eP, err := method.Inputs.Pack(args...)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	payload := append(fS, eP...)

	// build Transaction
	pb, err := utils.BuildTransaction(nS, from, cA, eth, gS, gPS, 6, payload)
	if err != nil {
		return err
	}

	// encode
	eR, err := encode.EncodeRlp(pb, payload)
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

	call, err := rpc.SubmitTx(tS)
	if err != nil {
		return err
	}
	logger.Log.Info(call)

	return nil
}

func SubmitTx(fromPriv string, contract string, function string, args []interface{}) error {

	cI, err := utils.GetContractInfo(contract)
	if err != nil {
		return err
	}

	cA, err := utils.GetContractAddress(contract)
	if err != nil {
		return err
	}

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
		fmt.Println(err)
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

	// encode parameters
	method := cI.Abi.Methods[function]
	fS, err := utils.EncodeFunctionSignature(method)
	if err != nil {
		return err
	}
	eP, err := method.Inputs.Pack(args...)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	payload := append(fS, eP...)

	// build Transaction
	pb, err := utils.BuildTransaction(nS, from, cA, "0", gS, gPS, 6, payload)
	if err != nil {
		return err
	}

	// encode
	eR, err := encode.EncodeRlp(pb, payload)
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

	call, err := rpc.SubmitTx(tS)
	if err != nil {
		return err
	}
	logger.Log.Info(call)

	return nil
}

func SubmitTxFromAddress(fromPriv string, contract string, function string, args []interface{}, address string) error {

	cI, err := utils.GetContractInfo(contract)
	if err != nil {
		return err
	}

	cA := address

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
		fmt.Println(err)
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

	// encode parameters
	method := cI.Abi.Methods[function]
	fS, err := utils.EncodeFunctionSignature(method)
	if err != nil {
		return err
	}
	eP, err := method.Inputs.Pack(args...)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	payload := append(fS, eP...)

	// build Transaction
	pb, err := utils.BuildTransaction(nS, from, cA, "0", gS, gPS, 6, payload)
	if err != nil {
		return err
	}

	// encode
	eR, err := encode.EncodeRlp(pb, payload)
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

	call, err := rpc.SubmitTx(tS)
	if err != nil {
		return err
	}
	logger.Log.Info(call)

	return nil
}
