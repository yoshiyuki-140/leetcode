package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// 未ソート部分から最小(または最大)の要素を探して、
	// 先頭(または末尾)と交換することを繰り返して配列をソートするアルゴリズム。
	// 簡単で直感的な方法だが、一般には非効率な方法とされている

	// 実装
	const length = 100
	nums := generateDemoData(length)
	fmt.Println(nums)
	sorted_nums := SelectionSort(nums, length)
	fmt.Println(sorted_nums)
}

// ダミーデータを用意
func generateDemoData(length int) []int {
	nums := make([]int, length)
	for i := 0; i < 100; i++ {
		nums[i] = rand.IntN(length)
	}
	return nums
}

// 選択ソート
func SelectionSort(nums []int, length int) []int {
	// 外側のループ
	// このループが終わると、i番目の要素はソート済みになる
	for i := 0; i < length-1; i++ {
		minIndex := i

		// 内側のループ : 最小の位置を探す
		for j := i; j < length; j++ {
			// より小さい値を見つけた場合、minIndexを更新
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		// 最小の要素と交換する
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}
