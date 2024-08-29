package types

type Contract struct {
	Deployer        string `json:"deployer"`
	Contract        string `json:"contract"`
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
}
