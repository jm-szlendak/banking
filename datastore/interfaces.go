package datastore

import (
	"github.com/jm-szlendak/banking/models"
)

type TransactionStore interface {
	GetAll() []models.Transaction
	Insert(transactions []models.Transaction, replace bool)
}
