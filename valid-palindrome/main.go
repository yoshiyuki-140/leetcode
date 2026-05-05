package validpalindrome

// 2つのポインタを使う方法
func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	s = toLower(s) // 小文字化

	// start == end か end < startの時に終了
	for start < end {
		if isAsciiValue(s[start]) && isAsciiValue(s[end]) { // どっちもasciiの時に比較を行う
			// 違っていたら即falseを返却
			if s[start] != s[end] {
				return false
			}
			// 進める
			start++
			end--
		} else {
			// asciiValueじゃなかったらポインタを進める
			if isAsciiValue(s[start]) == false {
				start++
			}
			if isAsciiValue(s[end]) == false {
				end--
			}
		}
	}
	return true
}

// 文字列に含まれるアルファベット大文字をアルファベット小文字に変換する関数
func toLower(s string) string {
	result := ""
	for _, c := range s {
		if 65 <= c && c <= 90 { // A-Z検知
			c += 32 // 小文字化
		}
		result += string(c)
	}
	return result
}

// a-z,A-Z,0-9だったらtrueを返却する
func isAsciiValue(byteS byte) bool {
	if (48 <= byteS && byteS <= 57) || // 0-9検知
		(65 <= byteS && byteS <= 90) || // A-Z検知
		(97 <= byteS && byteS <= 122) { // a-z検知
		return true
	}
	return false
}
