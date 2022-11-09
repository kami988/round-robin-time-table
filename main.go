package main

import "fmt"

const noPerson = -1

func main() {
	persons := []int{0, 1, 2, 3, 4, 5}
	combiArray := makeCombi(persons)
	fmt.Println()
	for _, v := range combiArray {
		fmt.Println(v)
	}
}

func makeCombi(persons []int) [][][]int {
	combiArray := [][][]int{}
	firstHalf, secondHalf := divideHalf(persons)

	for fi, fv := range firstHalf {
		if fv == noPerson {
			continue
		}
		passedValidation := false
		for _, sv := range secondHalf {
			if sv == noPerson {
				continue
			}
			// 初期は配列長をを確保する
			if !passedValidation {
				passedValidation = true
				combiArray = append(combiArray, [][]int{})
			}
			combiArray[fi] = append(combiArray[fi], []int{fv, sv})
		}
		secondHalf = shift(secondHalf)
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
			fmt.Println(secondCombi, "length", len(combiArray))
			combiArray[firstLen+i] = append(combiArray[firstLen+i], v...)
		}
	}
	return combiArray
}

// 配列を2つに分断する
func divideHalf(array []int) ([]int, []int) {
	// 偶数の場合、半分に割って取得
	if len(array)%2 == 0 {
		half := len(array) / 2
		return array[:half], array[half:]
	}
	// 奇数の場合、最後の人を先に休憩にするために前半が少なくなるよう取得
	// ex.
	// 1, 2, 0
	// 3, 4, 5
	center := (len(array) + 1) / 2
	// https://github.com/golang/go/wiki/SliceTricks#insert
	AddedArray := append(array[:center-1], append([]int{noPerson}, array[center-1:]...)...)
	return AddedArray[:center], AddedArray[center:]
}

// 配列を1つ分シフトする
func shift(array []int) []int {
	lastIndex := len(array) - 1
	res := []int{array[lastIndex]}
	return append(res, array[:lastIndex]...)
}
