# 解答

アルゴリズムの言語化

1. もし、rootがnilならば、falseを返却する
2. もし、subRootがnilならば、すべての場合において部分木になりえるのでtrueを返却する
3. もし、rootとsubRootが同じであれば、trueを返却する
    > rootとsubRoot配下の部分木が同じであるかどうかを判定するロジックは以下のようなものである
   1. もし、与えられたrootとsubRootのうちどちらのノードもnilであれば、それは同じノードであることを表すので、trueを返却する
   2. もし、どちらかのrootとsubRootのうちどちらかのノードがnilでどちらかがnilでないならば、それはa,bが違うノードであることを表すのでfalseを返却する
   3. もし、aとbのノードの値が違う場合は、違うノードなので、falseを返却する
   4. 上記の処理をa,bのノード配下の気に対して再帰的に行い、左右の木がどちらも同じ木であればtrueを返却する
4. 左右の部分木に対して上記の処理を再帰的に、処理し、どちらかの部分木が存在すればtrueを返却する.

```go
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
```