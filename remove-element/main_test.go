package removeelement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"Test1",
			args{nums: []int{3, 2, 2, 3}, val: 3},
			2,
		},
		{
			"Test2",
			args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
