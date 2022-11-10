package main

import (
	"fmt"
	"os"

	"github.com/kami988/round-robin-time-table/combination"
	csvconvert "github.com/kami988/round-robin-time-table/csvConvert"
	"github.com/kami988/round-robin-time-table/txt"
)

func main() {
	persons, err := txt.ReadTxtLine("./files/persons.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	combiArray := combination.MakeRoundRobinCombi(len(persons))

	if err = csvconvert.OutputCSV("./files/combination.csv", combiArray, persons); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
