# 解答

XOR演算を使えば、重複する要素のみ取り除かれて、最後に残るのは重複していない要素のみになる

## Best

https://leetcode.com/problems/single-number/solutions/7410577/xor-by-drobyshevv-qklf/

```go
func singleNumber(nums []int) int {
	XOR := nums[0]
	for i := 1; i < len(nums); i++ {
		XOR ^= nums[i]
	}
	return XOR
}
```