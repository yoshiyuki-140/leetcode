package mergetwosortedlists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	mockList1 := MakeListNodeFromArray([]int{1, 2, 4})
	mockList2 := MakeListNodeFromArray([]int{1, 3, 4})
	want := MakeListNodeFromArray([]int{1, 1, 2, 3, 4, 4})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[正常系] きちんとマージソートできる",
			args: args{
				list1: mockList1,
				list2: mockList2,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ListNodeを作成するヘルパー関数
func MakeListNodeFromArray(array []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, element := range array {
		current.Next = &ListNode{Val: element} // 新しいノードを作る
		current = current.Next
	}
	return dummy.Next
}

// 配列に変換するヘルパー関数
func MakeArrayFromListNode(node *ListNode) []int {
	array := make([]int, 1)
	if node != nil {
		array[0] = node.Val
		node = node.Next
	}
	for node != nil {
		array = append(array, node.Val)
		node = node.Next
	}
	return array
}

// ListNodeの中身を可視化するヘルパー関数
func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Printf("%d", node.Val)
		node = node.Next
		if node != nil {
			fmt.Printf(",")
		}
	}
	fmt.Printf("\n") // 改行用
}
