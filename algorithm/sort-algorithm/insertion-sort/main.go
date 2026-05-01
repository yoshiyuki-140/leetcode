package main

import "fmt"

func main() {
	fmt.Println("挿入ソート(InsertionSort)")
}

// 挿入ソート
func InsertionSort(nums []int) []int {
	// NOTE: 外側のループは,基準となる要素を一つずつ取り出すためのループ
	for i := 1; i < len(nums); i++ {
		// 基準となる要素を決める
		key := nums[i]

		j := i - 1

		// NOTE: 以下のfor文は挿入位置を決める
		// 終了条件は j < 0 && keyよりnums[j]が大きかったら
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		// 内側のループが終わり次第挿入する
		nums[j+1] = key // これはなぜここでj+1にするの？
	}
	return nums
}
