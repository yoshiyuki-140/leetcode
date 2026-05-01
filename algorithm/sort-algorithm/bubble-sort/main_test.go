package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ソートできてるか",
			args: args{
				nums: []int{9, 59, 3, 24, 5, 0},
			},
			want: []int{0, 3, 5, 9, 24, 59},
		},
		// 最悪時間計算量
		{
			name: "最悪時間計算量",
			args: args{
				nums: []int{59, 24, 9, 5, 3, 0},
			},
			want: []int{0, 3, 5, 9, 24, 59},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ベンチマーク
func BenchmarkBubbleSort(b *testing.B) {
	// 最悪時間計算量で実行
	BubbleSort([]int{59, 24, 9, 5, 3, 0})
}
