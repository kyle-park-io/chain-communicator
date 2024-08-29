package rpc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"chain-communicator/constant"
	"chain-communicator/logger"
	"chain-communicator/utils"
)

func GetBlockByHash(data string) (string, error) {
	req := &StringParamsRequest{
		Method: "block_by_hash",
		Params: []string{data},
		ID:     1,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Log.Errorln("Error marshalling request: ", err)
		return "", err
	}

	resp, err := http.Post(constant.Testnet_RpcUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return "", err
	}
	logger.Log.Infoln("Response: ", string(body))

	return string(body), nil
}

func GetTxByHash(data string) (string, error) {
	if !utils.IsHexString(data) {
		return "", fmt.Errorf("not hex string")
	}

	t, err := utils.HexStringToBytes(data)
	if err != nil {
		return "", err
	}
	tS := base64.StdEncoding.EncodeToString(t)

	var a []interface{}
	a = append(a, tS)
	a = append(a, true)

	req := &GenericParamsRequest{
		Method: "tx",
		Params: a,
		ID:     1,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Log.Errorln("Error marshalling request: ", err)
		return "", err
	}

	resp, err := http.Post(constant.Testnet_RpcUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return "", err
	}

	return string(body), nil
}

func Transfer(data string) (string, error) {
	req := &StringParamsRequest{
		// Method: "broadcast_tx_async",
		// Method: "broadcast_tx_sync",
		Method: "broadcast_tx_commit",
		Params: []string{data},
		ID:     1,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Log.Errorln("Error marshalling request: ", err)
		return "", err
	}

	resp, err := http.Post(constant.Testnet_RpcUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return "", err
	}

	return string(body), nil
}

func Deploy(data string) (string, error) {
	req := &StringParamsRequest{
		// Method: "broadcast_tx_sync",
		Method: "broadcast_tx_commit",
		Params: []string{data},
		ID:     1,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Log.Errorln("Error marshalling request: ", err)
		return "", err
	}

	resp, err := http.Post(constant.Testnet_RpcUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return "", err
	}

	return string(body), nil
}

func QueryTx(from, to, data string) (string, error) {
	var t []interface{}
	t = append(t, from)
	t = append(t, to)
	t = append(t, "0")
	t = append(t, data)

	req := &GenericParamsRequest{
		Method: "vm_call",
		Params: t,
		ID:     1,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Log.Errorln("Error marshalling request: ", err)
		return "", err
	}

	resp, err := http.Post(constant.Testnet_RpcUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return "", err
	}

	return string(body), nil
}

func SubmitTx(data string) (string, error) {
	req := &StringParamsRequest{
		// Method: "broadcast_tx_sync",
		Method: "broadcast_tx_commit",
		Params: []string{data},
		ID:     1,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Log.Errorln("Error marshalling request: ", err)
		return "", err
	}

	resp, err := http.Post(constant.Testnet_RpcUrl, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return "", err
	}

	return string(body), nil
}
