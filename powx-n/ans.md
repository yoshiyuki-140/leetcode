# 解答

## 最適解

- 時間計算量(log(N))
- 空間計算量O(1)

```go
func myPow(x float64, n int) float64 {
	// 負の最小値対策としてint64にキャスト
	N := int64(n)
	// nが負の場合は、xを逆数にしてnを正の数として扱う
	if N < 0 {
		x = 1 / x
		N = -N
	}

	var result float64 = 1.0
	currentProduct := x

	// Nが0になるまでループ(ここがO(logN)回まわる)
	for N > 0 {
		// Nの最下位ビットが1(つまり奇数)の場合
		// 今まで育ててきたcurrentProductを結果にかける
		if N%2 == 1 {
			result *= currentProduct
		}

		// currentProductを自乗して、次のビット(x^1 -> x^2 -> x^4...)に備える
		currentProduct *= currentProduct

		// Nを右に1ビットシフト (2で割るのと同義)
		N /= 2
	}
	return result
}
```

## スタックオーバーフローする解

```go
func myPow(x float64, n int) float64 {
	// 0乗は常に1
	// nが負の数の場合、nを正の数に転換して計算する
	if n < 0 {
		return 1.0 / myPow(x, -(n))
	}

	// NOTE: 指数を半分にした値を1回だけ計算してキープしておく
	half := myPow(x, n/2)

	// nが偶数なら、半分にしたものを2かいかけるだけ
	// e.g. x^10 = x^5 * x^5
	// nが奇数なら、xが1個余るので、それも掛ける
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
```