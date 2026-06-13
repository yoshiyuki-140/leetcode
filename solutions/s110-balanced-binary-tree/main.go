package balancedbinarytree

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

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
