package kthlargestelementinastream

import "container/heap"

// 最小ヒープ用の型を定義
type IntHeap []int

// sort.Interfaceの実装
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小ヒープ
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// heap.Interfaceの実装
// PushとPopはレシーバをポインタにする必要がある
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// LeetCodeの指定クラス (構造体)
type KthLargest struct {
	k    int
	heap *IntHeap
}

// コンストラクタ
func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	kl := KthLargest{
		k:    k,
		heap: h,
	}

	// 初期データを全てadd経由で追加
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

// 要素の追加と、k番目の値の取得
func (this *KthLargest) Add(val int) int {
	// ヒープに要素を追加
	heap.Push(this.heap, val)

	// サイズがkを超えたら、最小値をポップして捨てる
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}

	// 最小ヒープの先頭が、上位k個の中の最小値 = k番目に大きい値
	return (*this.heap)[0]
}
