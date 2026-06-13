package climbingstairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// n-1段目の初期値
	prev1 := 2
	// n-2段目の初期値
	prev2 := 1

	// 現在の値
	var curr int

	for i := 3; i <= n; i++ {
		// 計算
		curr = prev1 + prev2

		// 値を更新
		prev2 = prev1
		prev1 = curr
	}
	return curr
}
