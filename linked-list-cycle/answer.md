# 解説


解答コード
```go
func hasCycle(head *ListNode) bool {
	// 要素が0個の時と1つの時はサイクルにならないため、false
	if head == nil || head.Next == nil {
		return false
	}

	// フロイドの循環検出アルゴリズム
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
```