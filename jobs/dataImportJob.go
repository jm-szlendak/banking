package jobs

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/jm-szlendak/banking/dataimport"
	"github.com/jm-szlendak/banking/datastore"
)

type DataImportJob struct {
	Store datastore.TransactionStore
}

func (j *DataImportJob) Run(args []string) (JobResult, error) {
	var pkoImporter dataimport.PKOBPDataImporter
	var ingImporter dataimport.INGDataImporter

	workingDir := args[0]

	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return JobResult{}, err
	}
	for _, file := range files {
		var importer dataimport.Importer

		if strings.HasSuffix(file.Name(), ".csv") && strings.HasPrefix(file.Name(), "history_csv") {
			importer = pkoImporter
		} else if strings.HasSuffix(file.Name(), ".csv") && strings.HasPrefix(file.Name(), "Lista_transakcji") {
			importer = ingImporter
		} else {
			continue
		}

		filePath := strings.Join([]string{workingDir, file.Name()}, "/")
		dataFile, err := os.Open(filePath)
		if err != nil {
			return JobResult{}, err
		}
		data := importer.Import(dataFile, "Kuba")

		j.Store.Insert(data, true)
	}

	return JobResult{}, nil
}
