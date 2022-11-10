package array

import (
	"reflect"
	"testing"
)

func TestMakeIndexArray(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []int
	}{
		{
			name: "正常 配列長が0",
			args: 0,
			want: []int{},
		},
		{
			name: "正常 配列長が1",
			args: 1,
			want: []int{0},
		},
		{
			name: "正常 配列長が2",
			args: 2,
			want: []int{0, 1},
		},
		{
			name: "正常 配列長が5",
			args: 5,
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeIndexArray(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeIndexArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "正常系 nilの場合、空配列で返す",
			args: nil,
			want: []int{},
		},
		{
			name: "正常系 配列長が0",
			args: []int{},
			want: []int{},
		},
		{
			name: "正常系 配列長が1",
			args: []int{0},
			want: []int{0},
		},
		{
			name: "正常系 配列長が2",
			args: []int{0, 1},
			want: []int{1, 0},
		},
		{
			name: "正常系 配列長が5",
			args: []int{0, 1, 2, 3, 4},
			want: []int{4, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shift(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}
