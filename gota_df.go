package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	// step 1: open the csv
	csvfile, err := os.Open("example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// step 2: read the csv into memory.
	reader := csv.NewReader(csvfile)
	// note: you can also iterate through csv records with `.Read()` to handle row-level errors / manipulate data
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// step 3: load the csv into a dataframe
	df := dataframe.LoadRecords(records)
	fmt.Println("df: ", df)

	// step 4: add column is_even to the dataframe
	isEven := func(s series.Series) series.Series {
		num, _ := s.Int()
		isFavoriteNumberEven := num[0]%2 == 0
		return series.Bools(isFavoriteNumberEven)
	}
	isEvenSeries := df.Select("favorite_number").Rapply(isEven)
	isEvenSeries.SetNames("is_even")
	df = df.CBind(isEvenSeries)
	fmt.Println("df with is even: ", df)

	// step 5: filter the dataframe
	df = df.Filter(dataframe.F{"is_even", "==", true})
	fmt.Println("df filtered: ", df)

	// step 6: write csv
	f, err := os.Create("example_edited.csv")
	if err != nil {
		log.Fatal(err)
	}

	//df.WriteCSV(f)
}
