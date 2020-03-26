package main

import (
	"context"
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
)

func main() {
	ctx := context.TODO()

	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)

	fmt.Print(df.Table())

	df.Append(nil, 9, 123.6)

	df.Append(nil, map[string]interface{}{
		"day":   10,
		"sales": nil,
	})

	df.Remove(0)

	fmt.Print(df.Table())

	df.UpdateRow(0, nil, map[string]interface{}{
		"day":   3,
		"sales": 45,
	})

	fmt.Print(df.Table())

	sks := []dataframe.SortKey{
		{Key: "sales", Desc: true},
		{Key: "day", Desc: true},
	}

	df.Sort(ctx, sks)

	fmt.Print(df.Table())

	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true}) // Don't apply read lock because we are write locking from outside.

	df.Lock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		fmt.Println(*row, vals)
	}
	df.Unlock()

	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df, err := imports.LoadFromCSV(ctx, strings.NewReader(csvStr))

	fmt.Print(df.Table())
}
