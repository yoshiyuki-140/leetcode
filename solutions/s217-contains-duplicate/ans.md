# 解答

## 1. ハッシュマップを使う方法

- 時間計算量: O(n)
- 空間計算量: O(n)

```go
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}
```

## 2. ソートを使う方法

空間計算量をO(1)にする解法を面接官に聞かれたときは、これを答えるといい。

アルゴリズムのアイデアとして、配列に同じ数字が含まれていたら、ソートしたときに隣り合う数字が同じになることを応用するものです。

- 時間計算量: O(nlogn)
- 空間計算量: O(1)

```go
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
```