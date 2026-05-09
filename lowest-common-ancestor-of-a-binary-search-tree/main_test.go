package lowestcommonancestorofabinarysearchtree

import (
	"math"
	"testing"
)

const null = math.MinInt

// buildTree builds a TreeNode tree from a LeetCode-style BFS int slice.
// Use null (math.MinInt) to represent absent nodes.
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// findNode returns the first node in the tree whose Val equals val.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantVal int
	}{
		{
			// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
			// Output: 6
			name: "p=2, q=8, LCA=6",
			args: func() args {
				root := buildTree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5})
				return args{root: root, p: findNode(root, 2), q: findNode(root, 8)}
			}(),
			wantVal: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q)
			if got == nil || got.Val != tt.wantVal {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.wantVal)
			}
		})
	}
}
