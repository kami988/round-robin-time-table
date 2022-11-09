package main

import (
	"fmt"

	"github.com/kami988/round-robin-time-table/combination"
)

func main() {
	combiArray := combination.MakeRoundRobinCombi(6)

	for _, v := range combiArray {
		fmt.Println(v)
	}
}
