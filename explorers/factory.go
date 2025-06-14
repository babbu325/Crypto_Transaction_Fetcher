package explorers

import (
	"CoinTracker/config"
	"CoinTracker/explorers/ethscan"
	"errors"
)

func NewExplorer(cfg *config.Config, address string) (Explorer, error) {
	switch cfg.PreferredExplorer {
	case "etherscan":
		return ethscan.NewEtherScan(cfg.EtherscanAPIKey, address), nil
	default:
		return nil, errors.New("unsupported explorers: " + cfg.PreferredExplorer)
	}
}
