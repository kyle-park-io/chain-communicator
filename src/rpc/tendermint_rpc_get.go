package rpc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"chain-communicator/constant"
	"chain-communicator/logger"
	"chain-communicator/utils"
)

func BasicRpc() {
	endpoint := "block"
	values := url.Values{}
	values.Add("height", "514550")

	requestURL := constant.Testnet_RpcUrl + endpoint + "?" + values.Encode()

	resp, err := http.Get(requestURL)
	if err != nil {
		logger.Log.Errorln("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorln("Error reading response: ", err)
		return
	}
	logger.Log.Infoln("Response: ", string(body))
}

func GetAccount(address string) (string, error) {
	if !utils.IsHexAddress(address) {
		return "", fmt.Errorf("not hex Address")
	}

	endpoint := "account"
	values := url.Values{}
	values.Add("addr", address)

	requestURL := constant.Testnet_RpcUrl + endpoint + "?" + values.Encode()

	resp, err := http.Get(requestURL)
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

func GetRule() (string, error) {
	endpoint := "rule"

	requestURL := constant.Testnet_RpcUrl + endpoint + "?"

	resp, err := http.Get(requestURL)
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
