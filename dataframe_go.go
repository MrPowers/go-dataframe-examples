package main

import (
	//"encoding/csv"
	"context"
	"fmt"
	//"github.com/rocketlaunchr/dataframe-go/dataframe"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"log"
	"os"
	//"strings"
)

func main() {
	ctx := context.TODO()

	// step 1: open the csv
	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// step 2: read the csv into memory.
	//reader := csv.NewReader(csvfile)
	//records, err := reader.ReadAll()
	//if err != nil {
	//log.Fatal(err)
	//}

	dataframe, err := imports.LoadFromCSV(ctx, csvfile)

	fmt.Print(dataframe.Table())

	//s := dataframe.Series[2]

	//applyFn := dataframe.ApplySeriesFn(func(val interface{}, row, nRows int) interface{} {
	//return 2 * val.(int64)
	//})

	//dataframe.Apply(ctx, s, applyFn, dataframe.FilterOptions{InPlace: true})

	//fmt.Print(dataframe.Table())

	// looks like this file has relevant stuff: https://github.com/rocketlaunchr/dataframe-go/blob/0ec0a97d0f0d052e95557fee2aba1f3ce5bf94b1/apply.go

	// can't figure it out

}
