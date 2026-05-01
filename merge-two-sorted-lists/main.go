package mergetwosortedlists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	/*
		1. 先頭を格納
		# 連結リストの末尾にオブジェクトを追加するとき
		1.
	*/
	tmp := &ListNode{}
	result := tmp
	// the one
	if list1 == nil && list2 == nil {
		fmt.Println("1111111111")
		return nil
	}
	if list1 == nil {
		fmt.Println("2222222222")
		return list2
	}
	if list2 == nil {
		fmt.Println("3333333333")
		return list1
	}
	//
	if list1.Val <= list2.Val {
		fmt.Println("4444444444")
		tmp = list1
		list1 = list1.Next
	} else {
		fmt.Println("5555555555")
		tmp = list2
		list2 = list2.Next
	}
	//

	for list1.Next != nil && list2.Next != nil {
		fmt.Println("**********")
		// nullだったときの条件を追加
		if list1 == nil || list2 == nil {
			continue
		}
		if list1.Val <= list2.Val {
			tmp.Next = list1
			list1 = list1.Next
			fmt.Println("list1.Val <= list2.Val")
		} else {
			tmp.Next = list2
			list2 = list2.Next
			fmt.Println("list1.Val > list2.Val")
		}
		tmp = tmp.Next
	}
	return result
}
