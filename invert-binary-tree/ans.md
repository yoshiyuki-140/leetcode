# 解答

アルゴリズム言語化

1. もし、rootノードがnilであれば、nilを返却
2. 左右のノードを交換する
3. 左右の木に対して、1,2の処理を再帰的に行う
4. 反転済みの木の根を返す

```go
func invertTree(root *TreeNode) *TreeNode {
    // rootがnilであればnilを返却する
	if root == nil {
		return nil
	}

    // 左右のノードを交換する
	tmpNode := root.Left
	root.Left = root.Right
	root.Right = tmpNode

	// 左右の木に上記の処理を再帰的に行う
	invertTree(root.Left)
	invertTree(root.Right)

	// 反転済みの木の根を返す
	return root
}
```