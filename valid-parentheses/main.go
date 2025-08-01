package validParentheses

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var stack []rune
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v) // スタックに詰める
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]   // スタックから取り出す
			stack = stack[:len(stack)-1] //　後処理
			if !((top == '(' && v == ')') || (top == '{' && v == '}') || (top == '[' && v == ']')) {
				return false
			}
		}
	}
	// forループ終了時点でスタックが空であれば、括弧はきちんと閉じられていると判断
	return len(stack) <= 0
}
