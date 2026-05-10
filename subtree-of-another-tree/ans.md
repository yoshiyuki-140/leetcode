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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    if subRoot == nil {
        return true
    }
    if isIdent(root, subRoot) {
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// 木の形が同じであればtrueを返却する関数
func isIdent(a, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    if a.Val != b.Val {
        return false
    }
    return isIdent(a.Left, b.Left) && isIdent(a.Right, b.Right)
}
```