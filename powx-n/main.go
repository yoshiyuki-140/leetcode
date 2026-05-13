package powxn

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	// NOTE: int64への変換はしなくてもいい気がするので変換はやめておく
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

		currentProduct *= currentProduct

		// 1bit右シフト
		n >>= 1
	}

	return result
}
