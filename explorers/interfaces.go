package explorers

import model "CoinTracker/models"

type Explorer interface {
	FetchTransactions() ([]model.Transaction, error)
}
