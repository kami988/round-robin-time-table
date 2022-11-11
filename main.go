package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kami988/round-robin-time-table/combination"
	csvconvert "github.com/kami988/round-robin-time-table/csvConvert"
	"github.com/kami988/round-robin-time-table/txt"
)

func main() {
	var (
		personsFilePath     = flag.String("i", "./files/persons.txt", "人物のリストファイル")
		combinationFilePath = flag.String("o", "./files/combinations.csv", "人物のリストファイル")
	)
	flag.Parse()
	// Set up a connection to the server.
	persons, err := txt.ReadTxtLine(*personsFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	combiArray := combination.MakeRoundRobinCombi(len(persons))

	if err = csvconvert.OutputCSV(*combinationFilePath, combiArray, persons); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Combination generation completed successfully.")
	fmt.Println("input file path: ", *personsFilePath)
	fmt.Println("output file path: ", *combinationFilePath)
}
