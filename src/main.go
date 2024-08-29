package main

import "chain-communicator/logger"

func main() {
	// logger
	logger.InitLogger()
	logger.Log.Info("hi! start chain-communicator")
}
