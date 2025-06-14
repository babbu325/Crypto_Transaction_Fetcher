package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	PreferredExplorer string
	FallbackExplorers []string
	EtherscanAPIKey   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PreferredExplorer: os.Getenv("PREFERRED_EXPLORER"),
		FallbackExplorers: strings.Split(os.Getenv("FALLBACK_EXPLORERS"), ","),
		EtherscanAPIKey:   os.Getenv("ETHERSCAN_API_KEY"),
	}

}
