package datastore

import (
	"github.com/jm-szlendak/banking/models"
)

type TransactionStore interface {
	// Open() *TransactionStore
	Close()
	GetAll() ([]models.Transaction, error)
	Insert(transactions []models.Transaction, replace bool) error
}
