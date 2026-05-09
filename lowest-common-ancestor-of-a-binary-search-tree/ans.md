# 解答

アルゴリズム

現在のノードをrootとすると、以下の3つのパターンが考えられる.

1. 両方の値が、rootより小さい場合pもqも左の部分木にあるはずなので、左に移動する
2. 両方の値が、rootより大きい場合pもqも右の部分木にあるはずなので、右に移動する
3. 1,2以外の場合は、片方が小さく片方が大きいか、一方がroot自身となるので、このときのrootを返却する


```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 可読性のため、走査中の対象ノードをcurr変数としてコピーする
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
```