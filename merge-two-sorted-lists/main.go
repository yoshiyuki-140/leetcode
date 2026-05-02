package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	// Valを比較してcurrentに格納する
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// nilではないほうをcurrentのNextにつなげる
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// dummyのNextを返却する

	return dummy.Next
}
