package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	// 最大高さを格納するための変数を定義
	var maxDiameter int

	// 再帰的に探索するための関数を定義
	// dfs : depth-first search : 深さ優先探索
	// そのノードまでの深さを返す
	var dfs func(*TreeNode) int // 中で再帰的に呼び出すために、一度宣言しておく
	dfs = func(node *TreeNode) int {
		// 関数の中ですることは、
		// 0. もしnodeがnilであれば、0を返却する
		// 1. 左右の高さを再帰的に呼び出す
		// 2. 左右の高さの合計が最大高さを越しているかどうかを判定し、超えている場合は最大高さを更新する
		// 3. 現在のノードの高さ(1)を足してから、親のノードに、このノードのまでの高さを返却する

		// 0
		if node == nil {
			return 0
		}

		// 1
		rightHeight := dfs(node.Right)
		leftHeight := dfs(node.Left)

		// 2
		if rightHeight+leftHeight > maxDiameter {
			maxDiameter = rightHeight + leftHeight
		}

		// 3
		return max(rightHeight, leftHeight) + 1
	}

	// 再帰的に探索用関数を呼び出す
	dfs(root)

	// 保存していた最大高さの変数を返却する
	return maxDiameter
}
