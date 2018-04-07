package dataimport

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/jm-szlendak/banking/models"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

const (
	operationDateInd   = 0
	currencyDateInd    = 1
	transactionTypeInd = 2
	amountInd          = 3
	balanceInd         = 5
	descriptionInd1    = 6
	descriptionInd2    = 7
	descriptionInd3    = 8
	descriptionInd4    = 9
	descriptionInd5    = 10
)

type PKOBPDataImporter struct{}

func (p PKOBPDataImporter) Import(data io.Reader, accountId string) []models.Transaction {
	dec := transform.NewReader(data, charmap.Windows1250.NewDecoder())
	records, e := csv.NewReader(dec).ReadAll()
	if e != nil {
		panic(e)
	}

	transactions := make([]models.Transaction, len(records)-1, len(records)-1)

	for i, line := range records[1:] {
		dateLayout := "2006-01-02"

		operationDate, _ := time.Parse(dateLayout, line[operationDateInd])
		currencyDate, _ := time.Parse(dateLayout, line[currencyDateInd])
		amount, _ := strconv.ParseFloat(line[amountInd], 64)
		balance, _ := strconv.ParseFloat(line[balanceInd], 64)

		var transactionType models.TransactionType = models.Unknown
		var counterpart string
		var title string

		switch line[transactionTypeInd] {
		case "Płatność kartą":
			transactionType = models.CardPayment
			counterpart = parseCardPaymentDescription(line)
		case "Przelew na rachunek":
			transactionType = models.TransferIncoming
			title, counterpart = parseIncomingTransferDescription(line)
		case "Przelew z rachunku":
			transactionType = models.TransferOutgoing
			title, counterpart = parseOutgoingTransferDescription(line)
		case "Płatność web - kod mobilny":
			transactionType = models.TransferWeb
		case "Wypłata z bankomatu":
			transactionType = models.ATMWithdrawal
			counterpart = parseCardPaymentDescription(line)
		case "Opłata":
			transactionType = models.Charge
			title = line[descriptionInd1]
		case "Przelew na telefon przychodz. wew.":
			transactionType = models.TransferMobile
			title, counterpart = parseIncomingTransferDescription(line)
		}

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
		transaction.Id = transaction.Hash()

		transactions[i] = transaction
	}

	return transactions
}

func parseCardPaymentDescription(line []string) string {
	localization := strings.Replace(line[descriptionInd2], "Lokalizacja: Kraj: ", "", -1)
	localization = strings.Replace(localization, "Adres: ", "", -1)
	localization = strings.Replace(localization, "Miasto: ", "", -1)
	return localization
}

func parseIncomingTransferDescription(line []string) (string, string) {
	title := strings.Replace(line[descriptionInd4], "Tytuł: ", "", -1)
	counterpart := strings.Replace(line[descriptionInd2], "Nazwa nadawcy: ", "", -1)
	counterpartAddress := strings.Replace(line[descriptionInd3], "Adres nadawcy: ", "", -1)
	return title, strings.Join([]string{counterpart, counterpartAddress}, " ")
}

func parseOutgoingTransferDescription(line []string) (string, string) {
	var titleRaw string
	var counterpartAddress string
	counterpart := strings.Replace(line[descriptionInd2], "Nazwa odbiorcy: ", "", -1)
	if strings.Index(line[descriptionInd3], "Tytuł") == -1 {
		counterpartAddress = strings.Replace(line[descriptionInd3], "Adres odbiorcy: ", "", -1)
		titleRaw = line[descriptionInd4]
	} else {
		titleRaw = line[descriptionInd3]

	}
	title := strings.Replace(titleRaw, "Tytuł: ", "", -1)

	return title, strings.Join([]string{counterpart, counterpartAddress}, " ")
}
