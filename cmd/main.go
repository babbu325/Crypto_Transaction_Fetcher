package main

import (
	"CoinTracker/config"
	"CoinTracker/explorers"
	"CoinTracker/service"
	"flag"
	"fmt"
	"os"
)

func main() {

	address := flag.String("address", "", "Ethereum address to fetch transactions for")
	output := flag.String("output", "transactions.csv", "Output CSV file")
	flag.Parse()

	if *address == "" {
		fmt.Println("Please provide an Ethereum address using.")
		os.Exit(1)
	}

	cfg := config.LoadConfig()
	exp, err := explorers.NewExplorer(cfg, *address)
	if err != nil {
		fmt.Println("Error while creating explorer:", err)
		os.Exit(1)
	}

	exporter := service.NewExporter(exp)

	if err := exporter.Run(*output); err != nil {
		fmt.Println("Error:", err)
	}

}
