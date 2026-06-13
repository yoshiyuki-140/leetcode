# 解答

```go
func missingNumber(nums []int) int {
	hashMap := make(map[int]struct{}, len(nums)+1)
	for _, n := range nums {
		hashMap[n] = struct{}{} // データ型は何でもいい
	}
	// 探索
	for i := 0; i <= len(nums); i++ {
		if _, ok := hashMap[i]; !ok {
			return i
		}
	}
	return 0
}
```

## 別解(best)

https://leetcode.com/problems/missing-number/solutions/6728491/simple-mathematical-simplipication-beats-zxq6

```go
func missingNumber(nums []int) int {
	// 0からnまでの操作
	s := len(nums) * (len(nums) + 1) / 2
	// 一つずつ引いていく
	for _, n := range nums {
		s -= n
	}
	return s
}
```