package validparentheses

func isValid(s string) bool {
	// 偶数長ならfalseを返却
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	// スタックを初期化
	var stack []rune

	// 各文字をパースする
	for _, c := range s {
		// cが閉じ括弧の時
		if open, ok := pairs[c]; ok {
			// スタックの長さが0であれば、false
			if len(stack) == 0 {
				return false
			}
			// stackのtopに含まれるのがopenであった場合にはfalse
			if stack[len(stack)-1] != open {
				return false
			}
			// この時点で正しくペアになったので、stackからpop
			stack = stack[:len(stack)-1]
		} else {
			// 開き括弧の時スタックにpushする
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
