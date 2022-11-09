package main

import (
	"fmt"

	"github.com/kami988/round-robin-time-table/combination"
)

func main() {
	combiArray := combination.MakeRoundRobinCombi(10)

	for _, v := range combiArray {
		for _, v2 := range v {
			fmt.Printf("[%2d %2d] ", v2[0], v2[1])
		}
		fmt.Println()
	}
}
