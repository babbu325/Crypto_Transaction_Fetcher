package ethscan

import (
	model "CoinTracker/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

type EtherScan struct {
	APIKey  string
	Address string
}

const (
	baseURL = "https://api.etherscan.io/api/?module=account"
)

const (
	NormalTransfer   string = "txlist"
	InternalTransfer string = "txlistinternal"
	ERC20Transfer    string = "tokentx"
	ERC721Transfer   string = "tokennfttx"
	ERC1155Transfer  string = "token1155tx"
)

func NewEtherScan(apiKey, address string) *EtherScan {
	return &EtherScan{
		APIKey:  apiKey,
		Address: address,
	}
}

func (e *EtherScan) FetchTransactions() ([]model.Transaction, error) {

	normalTx, errNormalTx := e.FetchNormalTransactions(e.Address)
	if errNormalTx != nil {
		return nil, errNormalTx
	}

	internalTx, errInternalTx := e.FetchInternalTransactions(e.Address)
	if errInternalTx != nil {
		return nil, errInternalTx
	}

	erc20Tx, errErc20Tx := e.FetchERC20Transactions(e.Address)
	if errErc20Tx != nil {
		return nil, errErc20Tx
	}

	erc721Tx, errErc721Tx := e.FetchETC721Transactions(e.Address)
	if errErc721Tx != nil {
		return nil, errErc721Tx
	}

	erc1155Tx, errErc1155Tx := e.FetchERC1155Transactions(e.Address)
	if errErc1155Tx != nil {
		return nil, errErc1155Tx
	}

	var allTxs []model.Transaction
	allTxs = append(allTxs, normalTx...)
	allTxs = append(allTxs, internalTx...)
	allTxs = append(allTxs, erc20Tx...)
	allTxs = append(allTxs, erc721Tx...)
	allTxs = append(allTxs, erc1155Tx...)

	return allTxs, nil
}

func (e *EtherScan) FetchNormalTransactions(address string) ([]model.Transaction, error) {
	url := fmt.Sprintf("%s&action=%s&address=%s&startblock=0&offset=1000&sort=desc&apikey=%s", baseURL, NormalTransfer, address, e.APIKey)
	res := EtherScanNormalResponse{}

	if err := e.makeRequest(url, &res); err != nil {
		return nil, err
	}

	txs := []model.Transaction{}
	for _, tx := range res.Result {
		fee := calcFee(tx.GasUsed, tx.GasPrice)
		value := weiToEther(tx.Value)
		txs = append(txs, model.Transaction{
			TransferType: "Normal",
			Hash:         tx.Hash,
			Date:         timestampToDate(tx.TimeStamp),
			From:         tx.From,
			To:           tx.To,
			Amount:       value,
			Fee:          fee,
			AssetSymbol:  "ETH",
		})
	}
	return txs, nil
}

func (e *EtherScan) FetchInternalTransactions(address string) ([]model.Transaction, error) {
	url := fmt.Sprintf("%s&action=%s&address=%s&sort=desc&apikey=%s", baseURL, InternalTransfer, address, e.APIKey)
	res := EtherScanInternalResponse{}

	if err := e.makeRequest(url, &res); err != nil {
		return nil, err
	}

	txs := []model.Transaction{}
	for _, tx := range res.Result {

		value := weiToEther(tx.Value)
		txs = append(txs, model.Transaction{
			TransferType: "Internal",
			Hash:         tx.Hash,
			Date:         timestampToDate(tx.TimeStamp),
			From:         tx.From,
			To:           tx.To,
			Amount:       value,
			AssetSymbol:  "ETH",
			Contract:     tx.ContractAddress,
		})
	}
	return txs, nil
}

func (e *EtherScan) FetchERC20Transactions(address string) ([]model.Transaction, error) {
	url := fmt.Sprintf("%s&action=%s&address=%s&sort=desc&apikey=%s", baseURL, ERC20Transfer, address, e.APIKey)
	res := EtherScanERC20Response{}

	if err := e.makeRequest(url, &res); err != nil {
		return nil, err
	}

	txs := []model.Transaction{}
	for _, tx := range res.Result {
		value := weiToEther(tx.Value)
		txs = append(txs, model.Transaction{
			TransferType: "ERC20",
			Hash:         tx.Hash,
			Date:         timestampToDate(tx.TimeStamp),
			From:         tx.From,
			To:           tx.To,
			AssetSymbol:  tx.TokenSymbol,
			TokenID:      tx.TokenID,
			Contract:     tx.ContractAddress,
			Amount:       value,
		})
	}
	return txs, nil
}

func (e *EtherScan) FetchETC721Transactions(address string) ([]model.Transaction, error) {
	url := fmt.Sprintf("%s&action=%s&address=%s&page=1&offset=100&startblock=0&endblock=99999999&sort=asc&apikey=%s", baseURL, ERC721Transfer, address, e.APIKey)
	res := EtherScanERC721Response{}
	if err := e.makeRequest(url, &res); err != nil {
		return nil, err
	}
	txs := []model.Transaction{}
	for _, tx := range res.Result {
		txs = append(txs, model.Transaction{
			TransferType: "ETC721",
			Hash:         tx.Hash,
			Date:         timestampToDate(tx.TimeStamp),
			From:         tx.From,
			To:           tx.To,
			AssetSymbol:  tx.TokenSymbol,
			TokenID:      tx.TokenID,
			Contract:     tx.ContractAddress,
		})
	}
	return txs, nil
}

func (e *EtherScan) FetchERC1155Transactions(address string) ([]model.Transaction, error) {
	url := fmt.Sprintf("%s&action=%s&address=%s&page=1&offset=100&startblock=0&endblock=99999999&sort=asc&apikey=%s", baseURL, ERC1155Transfer, address, e.APIKey)
	res := EtherScanERC1155Response{}
	if err := e.makeRequest(url, &res); err != nil {
		return nil, err
	}
	txs := []model.Transaction{}
	for _, tx := range res.Result {
		value := weiToEther(tx.Value)
		txs = append(txs, model.Transaction{
			TransferType: "ETC1155",
			Hash:         tx.Hash,
			Date:         timestampToDate(tx.TimeStamp),
			From:         tx.From,
			To:           tx.To,
			AssetSymbol:  tx.TokenSymbol,
			TokenID:      tx.TokenID,
			Amount:       value,
			Contract:     tx.ContractAddress,
		})
	}
	return txs, nil
}

func (e *EtherScan) makeRequest(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, target); err != nil {
		return err
	}
	return nil
}

func calcFee(gasUsedStr, gasPriceStr string) string {
	gasUsed := new(big.Float)
	gasUsed.SetString(gasUsedStr)

	gasPrice := new(big.Float)
	gasPrice.SetString(gasPriceStr)

	feeWei := new(big.Float).Mul(gasUsed, gasPrice)
	feeETH := new(big.Float).Quo(feeWei, big.NewFloat(1e18))

	return feeETH.Text('f', 18)
}

func timestampToDate(ts string) string {
	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return ts
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func weiToEther(weiStr string) string {
	wei, _ := new(big.Float).SetString(weiStr)
	ether := new(big.Float).Quo(wei, big.NewFloat(1e18))
	return ether.Text('f', 18) // 18 decimal places
}
