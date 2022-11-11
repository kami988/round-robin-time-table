package combination

import (
	"github.com/kami988/round-robin-time-table/array"
)

const dummyVal = -1

func MakeSmartRoundRobinCombi(len int) [][][2]int {
	persons := array.MakeIndexArray(len)
	return makeSmartCombi(persons)
}

func makeSmartCombi(persons []int) [][][2]int {
	combiArray := make([][][2]int, len(persons), len(persons)+2)
	firstHalf, secondHalf := smartDivideHalf(persons)
	for i := range combiArray {
		for j := 0; j < len(firstHalf); j++ {
			if firstHalf[j] == dummyVal || secondHalf[j] == dummyVal {
				continue
			}
			combiArray[i] = append(combiArray[i], [2]int{firstHalf[j], secondHalf[j]})
		}
		// fmt.Println(firstHalf, secondHalf)
		// fmt.Println(combiArray[i])
		firstHalf, secondHalf = rotatePersons(firstHalf, secondHalf)
	}
	return combiArray
}

func rotatePersons(firstHalf []int, secondHalf []int) ([]int, []int) {
	// 謎にメモリバグが起きて、firstHalfとsecondHalfのメモリが共用されるため再確保する
	secondHalf = append([]int{}, secondHalf...)

	lastSecondIndex := len(secondHalf) - 1
	firstHalf = append(firstHalf, secondHalf[lastSecondIndex])
	secondHalf = append([]int{firstHalf[1]}, secondHalf...)
	firstHalf[1] = firstHalf[0]
	return firstHalf[1:], secondHalf[:lastSecondIndex+1]
}

func smartDivideHalf(array []int) ([]int, []int) {
	oddArray := []int{}
	evenArray := []int{}
	for i := range array {
		if i%2 == 0 {
			evenArray = append(evenArray, i)
		} else {
			oddArray = append(oddArray, i)
		}
	}
	if len(array)%2 == 1 {
		oddArray = append(oddArray, dummyVal)
	}
	return evenArray, oddArray
}
