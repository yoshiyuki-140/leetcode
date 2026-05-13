package powxn

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
