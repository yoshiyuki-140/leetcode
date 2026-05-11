# 解説


## アルゴリズム1

```go
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
```

## アルゴリズム2

アルゴリズム言語化
1. 長さが偶数であれば、括弧は閉じられないので、falseを返却する
2. スタックを定義
3. 開き括弧の場合は、スタックに格納
4. 閉じ括弧の場合は、スタックの領域が空であればfalseを返却．

```go
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	// 括弧の対応表をマップで定義(閉じ括弧をキーにすると判定が楽)
	pairs := map[rune]rune{
		// Key: Value
		'}': '{',
		']': '[',
		')': '(',
	}
	// スタックとして使うスライス
	var stack []rune

	for _, c := range s {
		// cが閉じ括弧かどうかをチェック
		if open, ok := pairs[c]; ok {
			// 1. スタックが空 : 対応する開き括弧がない
			// 2. スタックのトップが、期待される開き括弧ではない -> 最後の要素が
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// 正しくペアになったので、スタックからポップして元のスタックを更新
			stack = stack[:len(stack)-1]
		} else {
			// 開き括弧の場合はスタックにPush
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
```