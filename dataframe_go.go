package main

import (
	"context"
	"fmt"
	dataframe "github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"log"
	"os"
)

func main() {
	ctx := context.TODO()

	// step 1: open the csv
	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	df, err := imports.LoadFromCSV(ctx, csvfile)

	fmt.Print(df.Table())

	s := df.Series[1]

	applyFn := dataframe.ApplySeriesFn(func(val interface{}, row, nRows int) interface{} {
		return 2 * val.(int64)
	})

	dataframe.Apply(ctx, s, applyFn, dataframe.FilterOptions{InPlace: true})

	fmt.Print(df.Table())
}
