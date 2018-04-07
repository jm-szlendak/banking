package dataimport

import (
	"io"

	"github.com/jm-szlendak/banking/models"
)

type Importer interface {
	Import(data io.Reader, accountID string) []models.Transaction
}
