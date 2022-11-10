package combination

import (
	"github.com/kami988/round-robin-time-table/array"
)

const dummyVal = -1

func MakeRoundRobinCombi(len int) [][][2]int {
	persons := array.MakeIndexArray(len)
	return makeCombi(persons)
}

func makeCombi(persons []int) [][][2]int {
	combiArray := [][][2]int{}
	firstHalf, secondHalf := divideHalf(persons)

	for fi := 0; fi < len(firstHalf); fi++ {
		passedValidation := false
		for si := 0; si < len(secondHalf); si++ {
			if firstHalf[si] == dummyVal || secondHalf[si] == dummyVal {
				continue
			}
			// 初期は配列長をを確保する
			if !passedValidation {
				passedValidation = true
				combiArray = append(combiArray, [][2]int{})
			}
			combiArray[fi] = append(combiArray[fi], [2]int{firstHalf[si], secondHalf[si]})
		}
		secondHalf = array.Shift(secondHalf)
	}

	// 分割可能な場合、再起処理で組み合わせを生成する
	if len(firstHalf) >= 2 { // secondHalfも同じ長さのため、どちらでも良い
		// combiLen := len(combiArray) - 1
		firstCombi := makeCombi(firstHalf)
		secondCombi := makeCombi(secondHalf)
		mergedCombi := mergeCombi(firstCombi, secondCombi)
		// for i, v := range secondCombi {
		// 	combiArray[combiLen+i] = append(combiArray[combiLen+i], v...)
		// }
		combiArray = append(combiArray, mergedCombi...)
	}
	return combiArray
}

func mergeCombi(firstCombi [][][2]int, secondiCombi [][][2]int) [][][2]int {
	mergedCombi := [][][2]int{}

	// firstCombiの方が長い場合は想定しない
	for fi, fv := range firstCombi {
		mergedCombi = append(mergedCombi, append(fv, secondiCombi[fi]...))
	}
	for si := len(firstCombi); si < len(secondiCombi); si++ {
		mergedCombi = append(mergedCombi, secondiCombi[si])
	}
	return mergedCombi
}

// 配列を2つのグループに分断する
// ex. {1, 2, 3, 4, 5} の場合
// {1, 2, dummyVal},
// {3, 4, 5}
func divideHalf(array []int) ([]int, []int) {
	// 末尾のdummyValを削除
	if array[len(array)-1] == dummyVal {
		array = array[:len(array)-1]
	}

	// 偶数の場合、半分に割って取得
	if len(array)%2 == 0 {
		half := len(array) / 2
		return array[:half], array[half:]
	}
	// 奇数の場合、最後のを先に休憩にするために前半が少なくなるよう取得
	center := (len(array) + 1) / 2
	// https://github.com/golang/go/wiki/SliceTricks#insert
	AddedArray := append(array[:center-1], append([]int{dummyVal}, array[center-1:]...)...)
	return AddedArray[:center], AddedArray[center:]
}
