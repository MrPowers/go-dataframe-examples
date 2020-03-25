package main

import (
	"encoding/csv"
	"fmt"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"log"
	"os"
	"strings"
)

func main() {
	// step 1: open the csv
	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// step 2: read the csv into memory.
	reader := csv.NewReader(csvfile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	df, err := imports.LoadFromCSV(ctx, strings.NewReader(records))

	fmt.Print(df.Table())

}
