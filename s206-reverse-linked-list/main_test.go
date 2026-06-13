package reverselinkedlist

import (
	"reflect"
	"testing"
)

func makeListsFromSlice(slices []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, element := range slices {
		current.Next = &ListNode{Val: element}
		current = current.Next
	}
	return dummy.Next // dummy自体は常にリストの一つ前を指し続けているので、dummy.Nextを返却する
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[正常系] 正常に反転",
			args: args{head: makeListsFromSlice([]int{1, 2, 3, 4, 5})},
			want: makeListsFromSlice([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
