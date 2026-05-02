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

	// フロイドの循環検出アルゴリズム(空間計算量がO(1)になる)
	fast := head
	slow := head

	// 最後まで探索する
	for fast != nil && fast.Next != nil {
		// 始めはfast == slowになっていることは確実なので先に更新する
		slow = slow.Next      // 1個進む
		fast = fast.Next.Next // 2個進む
		if fast == slow {
			return true
		}
	}
	return false
}
