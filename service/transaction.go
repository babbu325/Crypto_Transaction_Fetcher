package service

import (
	"CoinTracker/explorers"
	"CoinTracker/utils"
)

type Exporter struct {
	Explorer explorers.Explorer
}

func NewExporter(exp explorers.Explorer) *Exporter {
	return &Exporter{Explorer: exp}
}

func (e *Exporter) Run(filename string) error {
	txs, err := e.Explorer.FetchTransactions()
	if err != nil {
		return err
	}
	return utils.WriteCSV(txs, filename)
}
