package powxn

func myPow(x float64, n int) float64 {
	// 負の最小値対策としてint64にキャスト
	N := int64(n)

	// nが負の場合は、xを逆数にしてnを正の数として扱う
	if N < 0 {
		x = 1 / x
		N = -N
	}

	// 結果を格納する変数を定義する
	var result float64 = 1.0
	currentProduct := x

	// Nが0になるまでループ(ここがO(logN)回まわる)
	for N > 0 {
		if N%2 == 1 {
			result *= currentProduct
		}

		// currentProductを自乗して、次のビットに備える
		currentProduct *= currentProduct

		// 右に1ビットシフト
		N /= 2
	}
	return result
}
