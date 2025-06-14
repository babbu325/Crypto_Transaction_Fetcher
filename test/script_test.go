package test

import (
	"CoinTracker/config"
	"CoinTracker/explorers"
	"CoinTracker/service"
	"fmt"
	"os"
	"testing"
)

func TestScript(t *testing.T) {
	cfg := &config.Config{
		PreferredExplorer: "etherscan",
		EtherscanAPIKey:   "H3K8ZX3FU9JYCHCRZ7WB5DN38ZD95N4C5E",
	}
	address := "0xa39b189482f984388a34460636fea9eb181ad1a6"
	exp, err := explorers.NewExplorer(cfg, address)
	if err != nil {
		fmt.Println("Error while creating explorer:", err)
		os.Exit(1)
	}

	exporter := service.NewExporter(exp)

	if err := exporter.Run("transaction_test.csv"); err != nil {
		fmt.Println("Error:", err)
	}

}
