package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"chain-communicator/constant"
	"chain-communicator/encode"
	"chain-communicator/logger"
	"chain-communicator/marshal"
	"chain-communicator/rpc"
	"chain-communicator/types"
	"chain-communicator/utils"
)

func Deploy(fromPriv string, contract string, args []interface{}) error {

	cI, err := utils.GetContractInfo(contract)
	if err != nil {
		return err
	}

	deployer, err := utils.PrivteToAddress(fromPriv)
	if err != nil {
		return err
	}

	// rpc - account
	data, err := rpc.GetAccount(deployer)
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
	method := cI.Abi.Constructor
	eP, err := method.Inputs.Pack(args...)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	bB, err := utils.HexStringToBytes(cI.Bytecode)
	if err != nil {
		return err
	}
	payload := append(bB, eP...)

	// build Transaction
	pb, err := utils.BuildTransaction(nS, deployer, constant.ZeroAddress, "0", gS, gPS, 6, payload)
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

	deploy, err := rpc.Deploy(tS)
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second)

	var parsedDeploy map[string]interface{}
	err = json.Unmarshal([]byte(deploy), &parsedDeploy)
	if err != nil {
		logger.Log.Error(err)
		return err
	}

	result3 := parsedDeploy["result"].(map[string]interface{})
	dT := result3["deliver_tx"].(map[string]interface{})
	aB, ok := dT["data"].(string)
	if !ok {
		return fmt.Errorf("returnData is not string")
	}
	addr, err := base64.StdEncoding.DecodeString(aB)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	addrS, err := utils.BytesToHexAdress(addr)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	logger.Log.Info(addrS)

	cJ := &types.Contract{
		Deployer:        deployer,
		Contract:        contract,
		ContractAddress: addrS,
		TxHash:          result3["hash"].(string),
	}
	jsonData, err := json.MarshalIndent(cJ, "", "  ")
	if err != nil {
		logger.Log.Errorln("Error encoding JSON:", err)
		return err
	}

	if contract == "Token" || contract == "TokenBurnable" {
		file, err := os.Create(fmt.Sprintf("contracts/%s.contract.json", args[0]))
		if err != nil {
			logger.Log.Errorln("Error creating file:", err)
			return err
		}
		defer file.Close()
		_, err = file.Write(jsonData)
		if err != nil {
			logger.Log.Errorln("Error writing to file:", err)
			return err
		}
	} else {
		file, err := os.Create(fmt.Sprintf("contracts/%s.contract.json", contract))
		if err != nil {
			logger.Log.Errorln("Error creating file:", err)
			return err
		}
		defer file.Close()
		_, err = file.Write(jsonData)
		if err != nil {
			logger.Log.Errorln("Error writing to file:", err)
			return err
		}
	}

	return nil
}
