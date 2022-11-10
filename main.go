package main

import (
	"fmt"
	"os"

	"github.com/kami988/round-robin-time-table/combination"
	"github.com/kami988/round-robin-time-table/txt"
)

func main() {
	persons, err := txt.ReadTxtLine("./persons.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	combiArray := combination.MakeRoundRobinCombi(len(persons))

	for _, v := range combiArray {
		for _, v2 := range v {
			fmt.Printf("[%2d %2d] ", v2[0], v2[1])
		}
		fmt.Println()
	}
}
