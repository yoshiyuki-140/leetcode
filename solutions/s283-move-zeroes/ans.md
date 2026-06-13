# 解答

問題は0を動かすように書いてあるから、ゼロを動かすのだと思ってしまいがちだ。
しかし、この問題はゼロ以外の値(NonZeroValue)を左に詰めるというアイデアで解いていく。

```go
func moveZeroes(nums []int) {
	// IDEA: ゼロ以外の値を左側に動かす(ゼロを動かすのではない！)
	// 左側から探索用のポインタを動かす
	lastNoneZeroPlace := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNoneZeroPlace], nums[i] = nums[i], nums[lastNoneZeroPlace] // 交換
			lastNoneZeroPlace++ // ポインタを一つ動かす
		}
	}
}
```