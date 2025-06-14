package ethscan

type EtherScanNormalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Hash            string `json:"hash"`
		TimeStamp       string `json:"timeStamp"`
		From            string `json:"from"`
		To              string `json:"to"`
		Value           string `json:"value"`
		GasUsed         string `json:"gasUsed"`
		GasPrice        string `json:"gasPrice"`
		ContractAddress string `json:"contractAddress"`
	} `json:"result"`
}

type EtherScanInternalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Hash            string `json:"hash"`
		TimeStamp       string `json:"timeStamp"`
		From            string `json:"from"`
		To              string `json:"to"`
		Value           string `json:"value"`
		ContractAddress string `json:"contractAddress"`
	} `json:"result"`
}

type EtherScanERC20Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Hash            string `json:"hash"`
		TimeStamp       string `json:"timeStamp"`
		From            string `json:"from"`
		To              string `json:"to"`
		Value           string `json:"value"`
		GasUsed         string `json:"gasUsed"`
		GasPrice        string `json:"gasPrice"`
		TokenName       string `json:"tokenName"`
		TokenSymbol     string `json:"tokenSymbol"`
		TokenID         string `json:"tokenID"`
		ContractAddress string `json:"contractAddress"`
	} `json:"result"`
}

type EtherScanERC721Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Hash            string `json:"hash"`
		TimeStamp       string `json:"timeStamp"`
		From            string `json:"from"`
		To              string `json:"to"`
		Value           string `json:"value"`
		GasUsed         string `json:"gasUsed"`
		GasPrice        string `json:"gasPrice"`
		TokenName       string `json:"tokenName"`
		TokenSymbol     string `json:"tokenSymbol"`
		TokenID         string `json:"tokenID"`
		ContractAddress string `json:"contractAddress"`
	} `json:"result"`
}

type EtherScanERC1155Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Hash            string `json:"hash"`
		TimeStamp       string `json:"timeStamp"`
		From            string `json:"from"`
		To              string `json:"to"`
		Value           string `json:"value"`
		GasUsed         string `json:"gasUsed"`
		GasPrice        string `json:"gasPrice"`
		TokenName       string `json:"tokenName"`
		TokenSymbol     string `json:"tokenSymbol"`
		TokenID         string `json:"tokenID"`
		ContractAddress string `json:"contractAddress"`
	} `json:"result"`
}
