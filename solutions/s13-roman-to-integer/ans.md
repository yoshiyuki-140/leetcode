# 解答


アルゴリズム言語化

- 基本的に足す
- もし、過去のやつより今のやつが大きかったら、過去に足した数を2倍したものを引く
  - これをすると、結局普通に計算したときと同じになる(紙面で確かめてほしい)

```go
func romanToInt(s string) int {
	// valueMap
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
		result += valMap[rune(v)]
		if i != 0 {
			if valMap[rune(s[i-1])] < valMap[rune(v)] {
				result -= 2 * valMap[rune(s[i-1])]
			}
		}
	}
	return result
}
```

## 初めて書いたコード


```go
/*
	---

	table1
	---
	CM 900
	CD 400
	XC 90
	XL 40
	IX 9
	IV 4
	---

	table2
	---
	I 1
	V 5
	X 10
	L 50
	C 100
	D 500
	M 1000
	---

	LVIII
	---
	1. Lに分解
		1. L -> 50
		2. V -> 5
		4. I -> 1
		4. I -> 1
		5. I -> 1
		= 58
	---

*/

func romanToInt(s string) int {
	symbols := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	result := 0
	S := strings.Split(s, "")
	l := len(S)
	for i := 0; i < l-1; i++ {
		if symbols[S[i]] < symbols[S[i+1]] {
			result -= symbols[S[i]]
		} else {
			result += symbols[S[i]]
		}
	}
	return result + symbols[S[l-1]]
}
```