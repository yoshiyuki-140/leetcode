package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "ソートできるか",
			args: args{
				nums: []int{64, 25, 12, 22, 11},
			},
			want: []int{11, 12, 22, 25, 64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	MergeSort([]int{64, 25, 12, 22, 11})
}
