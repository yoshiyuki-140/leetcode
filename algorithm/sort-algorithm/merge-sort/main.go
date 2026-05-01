package main

import "fmt"

func main() {
	fmt.Println("マージソート(MergeSort)")
	nums := []int{64, 25, 12, 2, 11}
	MergeSort(nums)
}

// マージソート
func MergeSort(nums []int) []int {
	// マージソートは配列を二つに分割し、それぞれをソートした後で結合(マージ)
	// することで配列全体をソートする.

	// 配列の要素数が1以下の場合、そのまま返す
	if len(nums) <= 1 {
		return nums
	}

	// 配列を中央で2つに分割
	// 奇数の場合はrightが1つ大きくなる
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	// 左右の配列をそれぞれ再帰的にソート
	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

// 2つのソート済み配列をお互いの要素を1津ずつ比較してマージする関数
func merge(left, right []int) []int {
	// 結果格納スライス
	result := make([]int, 0, len(left)+len(right))

	// 左右のソート済み配列のループ用のイテレーター
	l, r := 0, 0

	// leftとrightのどちらも要素が残っている間は、小さい方の要素をresultに追加
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	/*
		各left, rightともにスライスは既にソート済みであるため、
		先に処理が完了した後に残っている配列は最後尾に追加できる
	*/

	// leftに残った要素があればresultに追加
	result = append(result, left[l:]...)
	// rightにのこった要素があれば、resultに追加
	result = append(result, right[r:]...)

	return result
}
