package countingbits

func countBits(n int) []int {
	// n = 10 -> i = 0 ~ 9
	results := make([]int, n+1)
	for i := 0; i <= n; i++ {

		// バイナリ形式にした時の1の数を返却する
		result := 0
		j := i
		for j >= 1 {
			result += j % 2
			j /= 2
		}
		results[i] = result
	}
	return results
}
