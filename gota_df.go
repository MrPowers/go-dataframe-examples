package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	// load csv
	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(csvfile)

	fmt.Println("df: ", df)

	// add is_even column
	isEven := func(s series.Series) series.Series {
		num, _ := s.Int()
		isFavoriteNumberEven := num[0]%2 == 0
		return series.Bools(isFavoriteNumberEven)
	}
	isEvenSeries := df.Select("favorite_number").Rapply(isEven)
	isEvenSeries.SetNames("is_even")
	df = df.CBind(isEvenSeries)
	fmt.Println("df with is even: ", df)

	// filter the dataframe
	df = df.Filter(dataframe.F{"is_even", "==", true})
	fmt.Println("df filtered: ", df)

	// write csv
	f, err := os.Create("tmp/gota_example_output.csv")
	if err != nil {
		log.Fatal(err)
	}

	df.WriteCSV(f)
}
