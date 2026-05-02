package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 返却時に先頭を特定するためdummyを作る
	dummy := &ListNode{}
	current := dummy

	// 値を比較して、小さいほうをcurrentに入れる処理
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next // list1の更新
		} else {
			current.Next = list2
			list2 = list2.Next // listの更新
		}
		current = current.Next
	}

	// どちらかの要素がnilになったらnilではない要素をそのままcurrentに結合する
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	// dummyのnextを返却する
	return dummy.Next
}
