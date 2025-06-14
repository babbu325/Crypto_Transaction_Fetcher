package utils

import (
	"CoinTracker/models"
	"encoding/csv"
	"os"
)

func WriteCSV(txs []models.Transaction, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"TransferType", "Symbol", "Hash", "Date", "From", "To", "Contract", "TokenID", "Amount", "Fee"}
	writer.Write(headers)

	for _, tx := range txs {
		writer.Write([]string{
			tx.TransferType, tx.AssetSymbol, tx.Hash, tx.Date, tx.From, tx.To,
			tx.Contract, tx.TokenID, tx.Amount, tx.Fee,
		})
	}

	return nil
}
