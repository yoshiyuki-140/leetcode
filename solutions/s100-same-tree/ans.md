# 解答

アルゴリズム言語化

1. どっちもnilなら同じ木構造なのでtrue
2. どっちか一方がnilでもう一方がnil出ないならfalse
3. 値が違えばfalse
4. 1,2,3の処理をp,qの左右のノードに対して再帰的に行う


```go
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```