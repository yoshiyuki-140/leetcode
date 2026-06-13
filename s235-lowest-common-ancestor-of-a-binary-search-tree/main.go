package lowestcommonancestorofabinarysearchtree

// pre defined
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root

	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			// currがp,qより大きい場合は、pもqも左の部分木にあるはずなので、左に移動
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			// p,qがcurrより小さい場合は、pもqも右の部分木にあるはずなので、右に移動
			curr = curr.Right
		} else {
			// それ以外の場合は、片方が大きく片方が小さい、または一方がLCA自身なので、返却
			return curr
		}
	}
	// これで出てこない場合はp,qがBSTに無いのでnilを返却する
	return nil
}
