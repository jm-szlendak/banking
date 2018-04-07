package main

import (
	"fmt"
	"os"

	"github.com/jm-szlendak/banking/dataimport"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	checkError(err)

	defer file.Close()
	transactions := dataimport.ImportPKOBPData(file, "1")
	fmt.Println(transactions)

}
