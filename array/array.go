package array

// 値とインデックスが同じ配列を生成する
func MakeIndexArray(len int) []int {
	array := make([]int, len)
	for i := 0; i < len; i++ {
		array[i] = i
	}
	return array
}

// 配列を1つシフトする
func Shift(array []int) []int {
	lastIndex := len(array) - 1
	res := []int{array[lastIndex]}
	return append(res, array[:lastIndex]...)
}
