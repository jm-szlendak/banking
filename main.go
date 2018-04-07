package main

import (
	"fmt"
	"os"

	"github.com/jm-szlendak/banking/datastore"
	"github.com/jm-szlendak/banking/jobs"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	jobArgs := os.Args[1:]

	csvStore := datastore.NewCSVTransactionStore(jobArgs[1])
	defer csvStore.Close()

	job := jobs.DataImportJob{csvStore}
	_, err := job.Run(jobArgs)
	if err != nil {
		fmt.Println(err)
	}
}
