package constant

const (
	Testnet_RpcUrl = "http://xxx.xx/"
	Test_ChainId   = "testnet"

	Mainnet_RpcUrl = "https://xxx.xx/"
	Main_ChainId   = "mainnet"

	// ex. "\x19ETHEREUM(mainnet) Signed Message:\n100"
	TxPrefix = "\x19<NETWORK>(<SetYourChainId>) Signed Message:\n<SetYourEncodedDataLength>"

	ZeroAddress = "0x0000000000000000000000000000000000000000"
)
