package combination

import (
	"github.com/kami988/round-robin-time-table/array"
)

const dummyVal = -1

func MakeRoundRobinCombi(len int) [][][]int {
	persons := array.MakeIndexArray(len)
	return makeCombi(persons)
}

func makeCombi(persons []int) [][][]int {
	combiArray := [][][]int{}
	firstHalf, secondHalf := divideHalf(persons)

	for fi, fv := range firstHalf {
		if fv == dummyVal {
			continue
		}
		passedValidation := false
		for _, sv := range secondHalf {
			if sv == dummyVal {
				continue
			}
			// 初期は配列長をを確保する
			if !passedValidation {
				passedValidation = true
				combiArray = append(combiArray, [][]int{})
			}
			combiArray[fi] = append(combiArray[fi], []int{fv, sv})
		}
		secondHalf = array.LeftShift(secondHalf)
	}

	// 分割可能な場合、再起処理で組み合わせを生成する
	if len(firstHalf) >= 2 { // secondHalfも同じ長さのため、どちらでも良い
		firstCombi := makeCombi(firstHalf)
		for i, v := range firstCombi {
			combiArray[i] = append(combiArray[i], v...)
		}
		firstLen := len(firstCombi)
		secondCombi := makeCombi(secondHalf)
		for i, v := range secondCombi {
			combiArray[firstLen+i] = append(combiArray[firstLen+i], v...)
		}
	}
	return combiArray
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
