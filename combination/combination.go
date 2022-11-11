package combination

import (
	"github.com/kami988/round-robin-time-table/array"
)

const dummyVal = -1

func MakeRoundRobinCombi(len int) [][][2]int {
	persons := array.MakeIndexArray(len)
	combiArray := makeCombi(persons)

	// 人数が奇数の場合は休憩順が最後の人からになるようソート
	if len%2 == 1 {
		combiArray = sortCombi(combiArray)
	}
	return combiArray
}

func makeCombi(persons []int) [][][2]int {
	combiArray := make([][][2]int, len(persons), len(persons)+2)
	firstHalf, secondHalf := divideAlterHalf(persons)
	for i := range combiArray {
		for j := 0; j < len(firstHalf); j++ {
			if firstHalf[j] == dummyVal || secondHalf[j] == dummyVal {
				continue
			}
			combiArray[i] = append(combiArray[i], [2]int{firstHalf[j], secondHalf[j]})
		}
		firstHalf, secondHalf = rotatePersons(firstHalf, secondHalf)
	}
	return combiArray
}

func rotatePersons(firstHalf []int, secondHalf []int) ([]int, []int) {
	// 謎にメモリバグが起きて、firstHalfとsecondHalfのメモリが共用されるため再確保する
	secondHalf = append([]int{}, secondHalf...)

	if len(firstHalf) < 2 {
		return secondHalf, firstHalf
	}

	lastSecondIndex := len(secondHalf) - 1
	firstHalf = append(firstHalf, secondHalf[lastSecondIndex])
	secondHalf = append([]int{firstHalf[1]}, secondHalf...)
	firstHalf[1] = firstHalf[0]
	return firstHalf[1:], secondHalf[:lastSecondIndex+1]
}

func divideAlterHalf(array []int) ([]int, []int) {
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

// 休憩する人が最後の人から順番になるようにソート
func sortCombi(combiArray [][][2]int) [][][2]int {
	// 休憩する人を特定
	idSum := make([]int, len(combiArray))
	maxSum := 0
	for i, combis := range combiArray {
		for _, combi := range combis {
			for _, personID := range combi {
				idSum[i] += personID
			}
		}
		if maxSum < idSum[i] {
			maxSum = idSum[i]
		}
	}

	sortedCombi := make([][][2]int, len(combiArray))
	lastCombisIndex := len(combiArray) - 1
	for i, combis := range combiArray {
		sortedCombi[lastCombisIndex-(maxSum-idSum[i])] = combis
	}
	return sortedCombi
}
