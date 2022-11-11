package array

// 値とインデックスが同じ配列を生成する
func MakeIndexArray(len int) []int {
	array := make([]int, len)
	for i := 0; i < len; i++ {
		array[i] = i
	}
	return array
}
