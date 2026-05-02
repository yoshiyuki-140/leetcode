# 解説

---

### 実装コード

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// ダミーノードを作成して、マージ後のリストの開始地点を保持します
	dummy := &ListNode{}
	current := dummy

	// 両方のリストに要素がある間、値を比較して小さい方を連結していきます
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

	// どちらかのリストが先に空になった場合、残っている方のリストをそのまま連結します
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// ダミーノードの次が、実際のソート済みリストの先頭になります
	return dummy.Next
}
```

---

### アルゴリズムの解説

1.  **ダミーノード（Dummy Node）の活用**: 
    空の `dummy` ノードを用意することで、返却するリストの先頭を特定しやすくなります。これがないと、最初のノードを `list1` にするか `list2` にするかで個別の条件分岐が必要になり、コードが少し不格好になります。
2.  **ポインタの操作**:
    `current` ポインタを使って、新しく形成されるリストの最後尾を常に追跡します。
3.  **残った要素の処理**:
    `for` ループを抜けた時点で、少なくとも一方は `nil` になっています。連結リストはソート済みなので、残っているノード（およびその先に続くノード）をそのままガッチャンコと繋げるだけで完了です。

### 計算量

*   **時間計算量**: $O(n + m)$
    *   $n$ と $m$ はそれぞれのリストの長さです。すべてのノードを一度だけ走査します。
*   **空間計算量**: $O(1)$
    *   既存のノードを繋ぎ変えるだけなので、追加のメモリ（ノード）を生成しません（ダミーノード1つ分を除いて定数時間です）。

これで、効率的かつクリーンに2つのリストを統合できましたね。
