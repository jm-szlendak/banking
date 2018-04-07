package jobs

import (
	"github.com/jm-szlendak/banking/dataimport"
)

type DataImportJob struct{}

func (j *DataImportJob) Run(args []string) {
	var pkoImporter dataimport.PKOBPDataImporter
	importers := [...]dataimport.Importer{
		pkoImporter,
	}

}
