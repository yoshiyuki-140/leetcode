# 解答コード

```go
// リストの順番を逆転させるプログラム
func reverseList(head *ListNode) *ListNode {
	// 1. 始めは前は存在しないので nil
	var prev *ListNode
	curr := head

	for curr != nil {
		// 2. 矢印を書き換えると次にいけなくなるので、先に保存しておく
		nextTemp := curr.Next

		// 3. 矢印を前のほうへ向ける(ここで逆転)
		curr.Next = prev

		// 4. 次のステップのために prevとcurrをひとつづつ進める
		prev = curr
		curr = nextTemp
	}
	// 5. 最終的に currがnilになった時、prevが新しい先頭(元の末尾)を指している
	return prev
}
```