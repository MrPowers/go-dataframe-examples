package main

import (
	"encoding/csv"
	"fmt"
	"github.com/tobgu/qframe"
	"log"
	"os"
	"strings"
)

func main() {
	//input := `COL1,COL2
	//a,1.5
	//b,2.25
	//c,3.0`

	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvfile)
	// note: you can also iterate through csv records with `.Read()` to handle row-level errors / manipulate data
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f := qframe.ReadCSV(strings.NewReader(input))
	fmt.Println(f)

}
