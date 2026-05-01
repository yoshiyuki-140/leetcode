package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		nums   []int
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ソートされたかテスト",
			args: args{
				nums:   []int{9, 3, 4, 2, 0, 10},
				length: 6,
			},
			want: []int{0, 2, 3, 4, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.nums, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ベンチマークをとる
func BenchmarkSelectionSort(b *testing.B) {
	SelectionSort(generateDemoData(100), 100)
}
