package combination

import (
	"reflect"
	"testing"
)

func Test_mergeCombi(t *testing.T) {
	type args struct {
		firstCombi   [][][]int
		secondiCombi [][][]int
	}
	tests := []struct {
		name string
		args args
		want [][][]int
	}{
		{
			name: "正常系 中身がnilの場合、空配列を返す",
			args: args{
				firstCombi:   nil,
				secondiCombi: nil,
			},
			want: [][][]int{},
		},
		{
			name: "正常系 中身が空の場合、そのまま返す",
			args: args{
				firstCombi:   [][][]int{},
				secondiCombi: [][][]int{},
			},
			want: [][][]int{},
		},
		{
			name: "正常系 同じ長さの場合",
			args: args{
				firstCombi: [][][]int{
					{{0, 1}, {2, 3}},
				},
				secondiCombi: [][][]int{
					{{4, 5}},
				},
			},
			want: [][][]int{
				{{0, 1}, {2, 3}, {4, 5}},
			},
		},
		{
			name: "正常系 secondCombiの長さが大きいの場合",
			args: args{
				firstCombi: [][][]int{
					{{0, 1}, {2, 3}},
				},
				secondiCombi: [][][]int{
					{{4, 5}},
					{{1, 3}, {0, 2}},
				},
			},
			want: [][][]int{
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
