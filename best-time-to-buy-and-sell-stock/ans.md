# 解答

二つのポインタを利用する方式を採用する。
なぜなら、株取引を扱う問題であるため、単純にpricesの最大値から最小値を引けば良いわけではないからだ。
株取引において、過去の株がどれだけ高かったとしても、過去の株の値段で現在の株を売却することはできない。
今回使用するTwoポインタアルゴリズムはC言語で出てくるポインタの概念を使用するものではない。
今回の問題ではpricesが与えられるが、このpricesの場所を示すindexを二つの変数で管理する方法だ。

```go

func maxProfit(prices []int) int {
	l, r := 0, 1
	maxP := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if maxP < profit {
				maxP = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxP
}
```