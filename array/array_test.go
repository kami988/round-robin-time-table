package array

import (
	"reflect"
	"testing"
)

func TestMakeIndexArray(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want []int
	}{
		{
			name: "正常 配列長が0",
			arg:  0,
			want: []int{},
		},
		{
			name: "正常 配列長が1",
			arg:  1,
			want: []int{0},
		},
		{
			name: "正常 配列長が2",
			arg:  2,
			want: []int{0, 1},
		},
		{
			name: "正常 配列長が5",
			arg:  5,
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeIndexArray(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeIndexArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
