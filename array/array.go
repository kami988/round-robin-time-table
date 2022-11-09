package array

// 値とインデックスが同じ配列を生成する
func MakeIndexArray(len int) []int {
	array := make([]int, len)
	for i := 0; i < len; i++ {
		array[i] = i
	}
	return array
}

// 配列を1つ分左へシフトする
func LeftShift(array []int) []int {
	if len(array) < 2 {
		return array
	}
	return append(array[1:], array[0])
}
