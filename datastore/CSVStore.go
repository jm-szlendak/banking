package datastore

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/jm-szlendak/banking/models"
)

type CSVTransactionStore struct {
	writer io.WriteCloser
}

func (s *CSVTransactionStore) GetAll() ([]models.Transaction, error) {
	return make([]models.Transaction, 0, 0), nil
}

func (s *CSVTransactionStore) Insert(transactions []models.Transaction, replace bool) error {
	rows := make([][]string, len(transactions)+1)
	rows[0] = make([]string, 10)
	rows[0][0] = "Id"
	rows[0][1] = "Account"
	rows[0][2] = "TransactionDate"
	rows[0][3] = "CurrencyDate"
	rows[0][4] = "Type"
	rows[0][5] = "Amount"
	rows[0][6] = "Balance"
	rows[0][7] = "Title"
	rows[0][8] = "Counterpart"
	rows[0][9] = "Details"

	for i, transaction := range transactions {
		rows[i+1] = make([]string, 10)
		rows[i+1][0] = transaction.Id
		rows[i+1][1] = transaction.Account
		rows[i+1][2] = strconv.FormatInt(transaction.TransactionDate, 10)
		rows[i+1][3] = strconv.FormatInt(transaction.CurrencyDate, 10)
		rows[i+1][4] = strconv.Itoa(int(transaction.Type))
		rows[i+1][5] = strconv.FormatFloat(transaction.Amount, 'g', -1, 64)
		rows[i+1][6] = strconv.FormatFloat(transaction.Balance, 'g', -1, 64)
		rows[i+1][7] = transaction.Title
		rows[i+1][8] = transaction.Counterpart
		rows[i+1][9] = transaction.Details
	}

	csvWriter := csv.NewWriter(s.writer)
	csvWriter.WriteAll(rows)

	return nil
}

func (s *CSVTransactionStore) Close() {
	s.writer.Close()
}

func NewCSVTransactionStore(filename string) *CSVTransactionStore {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	return &CSVTransactionStore{file}
}
