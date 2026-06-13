# 解答

アルゴリズム言語化

1. 高さとエラーを伝播させる関数を再帰的に動かす
2. checkHeightは高さを返却する関数で、もし左右の木の高さの差が1より大きくなれば-1を返却する
3. もしcheckHeightが-1を返却したらそれを上のノードに伝播させる
4. 最終的にcheckHeightの戻り値が-1であればfalseを返却し、-1以外を返却すればtrueを返却する

```go
func isBalanced(root *TreeNode) bool {
	// checkHeightは、平衡であればその高さを
	// 平衡でなければ-1を返却する
	if checkHeight(root) == -1 {
		return false
	}
	return true
}

// checkHeightは、平衡であればその高さを
// 平衡でなければ-1を返却する
func checkHeight(node *TreeNode) int {
	// nodeがnilであれば0を返却する
	if node == nil {
		return 0
	}
	// 左部分木部分のチェック
	leftHeight := checkHeight(node.Left)
	if leftHeight == -1 {
		return -1 // 上に伝播させる
	}
	// 右部分木部分のチェック
	rightHeight := checkHeight(node.Right)
	if rightHeight == -1 {
		return -1 // 上に伝播させる
	}

	// 左右の高さの差を計算
	diff := abs(rightHeight, leftHeight)
	if diff > 1 {
		return -1
	}

	return max(rightHeight, leftHeight) + 1
}

// 差を返す
func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
```