# 解説

```go
func isValid(s string) bool {
	// sの長さが奇数長だったらfalseを返却する
	if len(s)%2 == 1 {
		return false
	}

	// pairを定義
	pair := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	// スタックの定義
	var stack []rune

	for _, c := range s {
		// 対応する括弧を呼び出す
		if open, ok := pair[c]; ok {
			// 1. スタックが空ならfalse
			if len(stack) == 0 {
				return false
			}
			// 2. スタックのトップが、期待される開き括弧ではないならfalse
			if stack[len(stack)-1] != open {
				return false
			}
			// 正しいペアになったので、スタックを更新
			stack = stack[:len(stack)-1]
		} else {
			// 開き括弧の場合はスタックにPush
			stack = append(stack, c)
		}
	}

	// 最終的にスタックが空ならOK
	return len(stack) == 0
}
```