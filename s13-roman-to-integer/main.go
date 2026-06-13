package romantointeger

func romanToInt(s string) int {
	valMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i, v := range s {
		result += valMap[v]
		// 前のやつより今のやつが大きい場合
		if i != 0 {
			if valMap[rune(s[i-1])] < valMap[v] {
				result -= 2 * valMap[rune(s[i-1])]
			}
		}
	}
	return result
}
