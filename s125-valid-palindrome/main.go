package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	// 1. 小文字化
	s = strings.ToLower(s)
	// 2. twoポインタを使って比較
	start, end := 0, len(s)-1

	// startがend以上になったら終了
	for start < end {
		startIsValud := isValidValue(s[start])
		endIsValud := isValidValue(s[end])
		if !startIsValud { // startがascii文字列じゃなかったらポインタを進めてスキップ
			start++
			continue
		}
		if !endIsValud { // endがascii文字列じゃなかったらポインタを進めてスキップ
			end--
			continue
		}
		if startIsValud && endIsValud { // startとendのポインタが示す値がどちらもascii文字列であれば比較を行う
			// 比較を行って値が違えば即falseを返却
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
	}
	// 全てが適切であれば回文判定としてtrueを返却
	return true
}

// A-Z, a-z, 0-9であればtrueを返却する
func isValidValue(c byte) bool {
	if (48 <= c && c <= 57) || // 0-0
		(65 <= c && c <= 90) || // A-Z
		(97 <= c && c <= 122) { // a-z
		return true
	}
	return false
}
