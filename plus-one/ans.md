# 解答

アルゴリズム言語化

1. 後ろから見ていく
2. 下一桁が9より小さい値であれば、その値に1加算してからdigitsを返却する
3. 下一桁が9であった場合は、その桁を0にしてから、上の桁の処理に移る。
4. ... これらを反復処理する。
5. 最終的に、もしこの反復処理が全て終わった場合は、digitsは(e.g. [0,0,0,0])のようになるため、前に1を付けてからdigitsを返却する

```go
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
```