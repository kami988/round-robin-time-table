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
		personsFilePath     = flag.String("i", "./files/persons.txt", "人物のリストファイルのパス")
		combinationFilePath = flag.String("o", "./files/combinations.csv", "出力ファイル（CSV）のパス")
	)
	flag.Parse()

	persons, err := txt.ReadTxtLine(*personsFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(persons) < 3 {
		fmt.Fprintln(os.Stderr, "Must be 3 or more.")
		os.Exit(1)
	}

	combiArray := combination.MakeRoundRobinCombi(len(persons))

	if err = csvconvert.OutputCSV(*combinationFilePath, combiArray, persons); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("入力ファイルパス: ", *personsFilePath)
	fmt.Println("出力ファイルパス: ", *combinationFilePath)
	fmt.Println("組み合わせの生成に成功したにゃ！")
}
