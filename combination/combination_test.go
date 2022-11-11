package combination

import (
	"reflect"
	"testing"
)

func Test_rotatePersons(t *testing.T) {
	type args struct {
		firstHalf  []int
		secondHalf []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "正常系 配列長が0",
			args: args{
				firstHalf:  []int{},
				secondHalf: []int{},
			},
			want:  []int{},
			want1: []int{},
		},
		{
			name: "正常系 配列長が1",
			args: args{
				firstHalf:  []int{0},
				secondHalf: []int{1},
			},
			want:  []int{1},
			want1: []int{0},
		},
		{
			name: "正常系 配列長が3",
			args: args{
				firstHalf:  []int{0, 1, 2},
				secondHalf: []int{3, 4, 5},
			},
			want:  []int{0, 2, 5},
			want1: []int{1, 3, 4},
		},
		{
			name: "正常系 配列長が4",
			args: args{
				firstHalf:  []int{0, 1, 2, 3, -1},
				secondHalf: []int{4, 5, 6, 7, 8},
			},
			want:  []int{0, 2, 3, -1, 8},
			want1: []int{1, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := rotatePersons(tt.args.firstHalf, tt.args.secondHalf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotatePersons() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("rotatePersons() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_divideAlterHalf(t *testing.T) {
	tests := []struct {
		name  string
		arg   []int
		want  []int
		want1 []int
	}{
		{
			name:  "正常系 配列長が0",
			arg:   []int{},
			want:  []int{},
			want1: []int{},
		},
		{
			name:  "正常系 配列長が1",
			arg:   []int{0},
			want:  []int{0},
			want1: []int{-1},
		},
		{
			name:  "正常系 配列長が2",
			arg:   []int{0, 1},
			want:  []int{0},
			want1: []int{1},
		},
		{
			name:  "正常系 配列長が5",
			arg:   []int{0, 1, 2, 3, 4},
			want:  []int{0, 2, 4},
			want1: []int{1, 3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := divideAlterHalf(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideAlterHalf() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("divideAlterHalf() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
