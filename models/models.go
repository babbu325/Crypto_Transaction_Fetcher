package models

type Transaction struct {
	TransferType string
	Hash         string
	Date         string
	From         string
	To           string
	Contract     string
	AssetSymbol  string
	TokenID      string
	Amount       string
	Fee          string
}
