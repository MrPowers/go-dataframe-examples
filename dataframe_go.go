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

	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	df, err := imports.LoadFromCSV(ctx, csvfile, imports.CSVLoadOptions{
		DictateDataType: map[string]interface{}{
			"first_name":      "",       // specify this column as string
			"favorite_number": int64(0), // specify this column as int64
		}})

	fmt.Print(df.Table())

	s := df.Series[1]

	applyFn := dataframe.ApplySeriesFn(func(val interface{}, row, nRows int) interface{} {
		return 2 * val.(int64)
	})

	dataframe.Apply(ctx, s, applyFn, dataframe.FilterOptions{InPlace: true})

	fmt.Print(df.Table())
}
