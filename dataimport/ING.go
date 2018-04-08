package dataimport

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/jm-szlendak/banking/models"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type INGDataImporter struct{}

func (p INGDataImporter) Import(data io.Reader, accountId string) []models.Transaction {
	const (
		operationDateInd   = 1
		currencyDateInd    = 0
		transactionTypeInd = 6
		amountInd          = 8
		balanceInd         = 14
		titleInd           = 3
		counterpartInd     = 2
	)

	dec := transform.NewReader(data, charmap.Windows1250.NewDecoder())
	r := csv.NewReader(dec)
	r.Comma = ';'
	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		records = append(records, record)
	}

	dateLayout := "2006-01-02"

	for _, record := range records {
		fmt.Println(record)
	}

	var transactions []models.Transaction
	for _, line := range records[5:] {
		if len(line[amountInd]) == 0 {
			continue
		}
		operationDate, _ := time.Parse(dateLayout, line[operationDateInd])
		currencyDate, _ := time.Parse(dateLayout, line[currencyDateInd])
		amount, _ := strconv.ParseFloat(strings.Replace(line[amountInd], ",", ".", -1), 64)
		balance, _ := strconv.ParseFloat(strings.Replace(line[balanceInd], ",", ".", -1), 64)
		var transactionType models.TransactionType = models.Unknown

		switch {
		case strings.Index(line[transactionTypeInd], "TR.KART") != -1:
			transactionType = models.CardPayment
		case strings.Index(line[transactionTypeInd], "PRZELEW") != -1:
			transactionType = models.TransferIncoming
		case strings.Index(line[transactionTypeInd], "TR.BLIK") != -1:
			transactionType = models.TransferWeb
		}

		transaction := models.Transaction{
			TransactionDate: operationDate.Unix(),
			CurrencyDate:    currencyDate.Unix(),
			Type:            transactionType,
			Amount:          amount,
			Balance:         balance,
			Title:           line[titleInd],
			Counterpart:     line[counterpartInd],
			Account:         accountId,
		}

		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
	return transactions
}
