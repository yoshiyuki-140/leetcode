package linkedlistcycle

import "testing"

// テスト用の循環しているリストを作成するヘルパー関数
func getCycledList() *ListNode {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node0 := &ListNode{Val: 0}
	node4 := &ListNode{Val: 4}

	// 3 -> 2 -> 0 -> 4
	node3.Next = node2
	node2.Next = node0
	node0.Next = node4

	// 4 -> 2 (ここでサイクルをつなぐ)
	node4.Next = node2

	return node3
}

// テスト用の循環していないリストを作成するヘルパー関数
func getNotCycledList() *ListNode {
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node0 := &ListNode{Val: 0}
	node4 := &ListNode{Val: 4}

	// 3 -> 2 -> 0 -> 4
	node3.Next = node2
	node2.Next = node0
	node0.Next = node4

	return node3
}

func Test_hasCycle(t *testing.T) {

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[正常系] サイクルがあるListはtrueを返却する",
			args: args{
				head: getCycledList(),
			},
			want: true,
		},
		{
			name: "[正常系] サイクルがないListはfalseを返却する",
			args: args{
				head: getNotCycledList(),
			},
			want: false,
		},
		{
			name: "[正常系] リストの長さが0の場合はfalseを返却する",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "[正常系] リストの長さが1の場合はtrueを返却する",
			args: args{
				head: &ListNode{Val: 0, Next: nil},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
