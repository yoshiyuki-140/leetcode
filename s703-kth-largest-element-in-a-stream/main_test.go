package kthlargestelementinastream

import (
	"reflect"
	"testing"
)

func TestIntHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("IntHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    IntHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("IntHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    IntHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestIntHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    *IntHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
		})
	}
}

func TestIntHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *IntHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want KthLargest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.k, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKthLargest_Add(t *testing.T) {
	type fields struct {
		k    int
		heap *IntHeap
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &KthLargest{
				k:    tt.fields.k,
				heap: tt.fields.heap,
			}
			if got := this.Add(tt.args.val); got != tt.want {
				t.Errorf("KthLargest.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
