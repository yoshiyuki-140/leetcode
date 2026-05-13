# 解答

## 最適解

- 時間計算量(log(N))
- 空間計算量O(1)

アルゴリズム言語化

- n回計算する代わりに、計算結果を使いまわしてショートカットをしたい
- 指数を2進数としてとらえて、1, 2, 4, 8... と倍々になったパーツを作っていく
- 必要なパーツだけを掛け算して、たったlogN回で答えが得られる

```go
func myPow(x float64, n int) float64 {
	if n == 0 { // x^0 == 1.0
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var result float64 = 1.0
	currentProduct := x // 現在の掛け算の値を計算する

	for n > 0 {
		if n%2 == 1 {
			result *= currentProduct
		}

		currentProduct *= currentProduct // 自分自身を倍々にする

		// 1bit右シフト
		n >>= 1
	}

	return result
}
```
