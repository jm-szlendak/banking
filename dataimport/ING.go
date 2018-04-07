package dataimport

import (
	"encoding/csv"
	"io"

	"github.com/jm-szlendak/banking/models"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type INGDataImporter struct{}

func (p INGDataImporter) Import(data io.Reader, accountId string) []models.Transaction {
	dec := transform.NewReader(data, charmap.Windows1250.NewDecoder())
	records, e := csv.NewReader(dec).ReadAll()
	if e != nil {
		panic(e)
	}

	transactions := make([]models.Transaction, len(records)-1, len(records)-1)

	for i, line := range records[1:] {
		dateLayout := "2006-01-02"

		transaction := models.Transaction{
			TransactionDate: operationDate.Unix(),
			CurrencyDate:    currencyDate.Unix(),
			Type:            transactionType,
			Amount:          amount,
			Balance:         balance,
			Title:           title,
			Counterpart:     counterpart,
			Account:         accountId,
		}

		transactions[i] = transaction
	}
	return transactions
}
