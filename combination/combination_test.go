package combination

import (
	"reflect"
	"testing"
)

func Test_mergeCombi(t *testing.T) {
	type args struct {
		firstCombi   [][][2]int
		secondiCombi [][][2]int
	}
	tests := []struct {
		name string
		args args
		want [][][2]int
	}{
		{
			name: "正常系 中身がnilの場合、空配列を返す",
			args: args{
				firstCombi:   nil,
				secondiCombi: nil,
			},
			want: [][][2]int{},
		},
		{
			name: "正常系 中身が空の場合、そのまま返す",
			args: args{
				firstCombi:   [][][2]int{},
				secondiCombi: [][][2]int{},
			},
			want: [][][2]int{},
		},
		{
			name: "正常系 同じ長さの場合",
			args: args{
				firstCombi: [][][2]int{
					{{0, 1}, {2, 3}},
				},
				secondiCombi: [][][2]int{
					{{4, 5}},
				},
			},
			want: [][][2]int{
				{{0, 1}, {2, 3}, {4, 5}},
			},
		},
		{
			name: "正常系 secondCombiの長さが大きいの場合",
			args: args{
				firstCombi: [][][2]int{
					{{0, 1}, {2, 3}},
				},
				secondiCombi: [][][2]int{
					{{4, 5}},
					{{1, 3}, {0, 2}},
				},
			},
			want: [][][2]int{
				{{0, 1}, {2, 3}, {4, 5}},
				{{1, 3}, {0, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeCombi(tt.args.firstCombi, tt.args.secondiCombi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeCombi() = %v, want %v", got, tt.want)
			}
		})
	}
}
