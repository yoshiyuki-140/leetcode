package subtreeofanothertree

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isSubtree(root, subRoot *TreeNode) bool {
	// 1. もしrootがnilならばfalseを返却する。なぜならば、その木は部分木がないからだ
	if root == nil {
		return false
	}
	// 2. もしsubRootがnilならばtrueを返却する。なぜならば、その部分木はすべての場合に部分木になるからだ
	if subRoot == nil {
		return true
	}
	// 3. もしrootとsubRootが同じ(isIdent()がtrue)ならば、trueを返却する
	if isIdent(root, subRoot) {
		return true
	}
	// 4. 上記の処理を左右の木について再帰的に行い、どちらかの木がtrueになればtrueを返却する
	return isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
}

// 配下の木まですべて同じ木であること判定する関数
// もし、同じ木であれば、trueを返却する
func isIdent(a, b *TreeNode) bool {
	// 判定アルゴリズム
	// 1. a,bが共にnilならば、同じノードだと判定してtureを返却する
	if a == nil && b == nil {
		return true
	}
	// 2. a,bのどちらかがnilでどちらかがnilではない場合は、別の形をしている部分木になるのでfalseを返却する
	if a == nil || b == nil {
		return false
	}
	// 3. a,bの値が違う場合は、別の木になるのでfalseを返却する
	if a.Val != b.Val {
		return false
	}
	// 4. a,b各々の左右の部分木について、上記の判定ロジックを走らせて、どちらもtrueになった場合はtrueを返却する
	return isIdent(a.Right, b.Right) && isIdent(a.Left, b.Left)
}
