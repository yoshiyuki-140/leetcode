package climbingstairs

func climbStairs(n int) int {
	// 1段と2段の場合は、そのまま段数が答えになる
	if n <= 2 {
		return n
	}

	// n-2段目までの通り数
	prev2 := 1
	// n-1段目までの通り数
	prev1 := 2

	// 3段目から計算を開始
	for i := 3; i <= n; i++ {

		// 現在の段への通り数は、(n-1段目の通り数) + (n-2段目の通り数)
		current := prev1 + prev2

		// 次のループのために値を更新
		prev2 = prev1
		prev1 = current
	}
	return prev1
}
