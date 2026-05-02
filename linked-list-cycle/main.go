package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 要素が0個の時と1つの時はサイクルにならないため、false
	if head == nil || head.Next == nil {
		return false
	}

	// メモ用
	pointerMap := make(map[*ListNode]bool)

	// 最後まで探索する
	for head != nil {
		if pointerMap[head] {
			return true
		}
		pointerMap[head] = true
		head = head.Next
	}
	return false
}
