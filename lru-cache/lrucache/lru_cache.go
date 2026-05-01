package lrucache

// NOTE: 双方向リストの定義
// Node
// - Prev *Node
// - Next *Node

// List
// - Head *Node
// - Tail *Node

// 今回必要になる処理
// -

type LRUCache struct {
}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key, value int) int {
	return 0
}
