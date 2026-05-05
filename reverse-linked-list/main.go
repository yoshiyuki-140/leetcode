package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head // 現在の操作対称を選択

	for curr != nil {

		// 現在のノードの次のノードを示すアドレスを一時的に保存
		nextTemp := curr.Next
		// 現在のノードの次のノードを示すアドレスを、現在のノードの前のノードのアドレスで書き換える
		curr.Next = prev

		// 走査対象を次のノードに更新
		prev = curr
		curr = nextTemp
	}

	// 最終的にprevが最後に残るノードとなる
	return prev
}
