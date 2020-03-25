package main

import (
	"fmt"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/function"
	"log"
	"os"
)

func main() {
	// open a CSV file
	csvfile, err := os.Open("data/basic.csv")
	if err != nil {
		log.Fatal(err)
	}

	// view the DataFrame
	f := qframe.ReadCSV(csvfile)
	fmt.Println(f)

	// concatenate two DataFrame columns
	f = f.Apply(
		qframe.Instruction{
			Fn:      function.ConcatS,
			DstCol:  "col3",
			SrcCol1: "col1",
			SrcCol2: "col2"})

	fmt.Println(f)

	// write the DataFrame
	file, err := os.Create("tmp/qframe_ouput.csv")
	if err != nil {
		log.Fatal(err)
	}
	f.ToCSV(file)

}
