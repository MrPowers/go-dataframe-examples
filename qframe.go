package main

import (
	"fmt"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/function"
	"log"
	"os"
)

func concatExample() {
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

func isEven(x int) bool {
	return x%2 == 0
}

func mainExample() {
	csvfile, err := os.Open("data/example.csv")
	if err != nil {
		log.Fatal(err)
	}

	f := qframe.ReadCSV(csvfile)
	fmt.Println(f)

	f = f.Apply(
		qframe.Instruction{
			Fn:      isEven,
			DstCol:  "is_even",
			SrcCol1: "favorite_number"})

	fmt.Println(f)

	newF := f.Filter(qframe.Filter{Column: "is_even", Comparator: "=", Arg: true})
	fmt.Println(newF)

	file, err := os.Create("tmp/qframe_main_ouput.csv")
	if err != nil {
		log.Fatal(err)
	}
	newF.ToCSV(file)
}

func main() {
	//concatExample()
	mainExample()
}
