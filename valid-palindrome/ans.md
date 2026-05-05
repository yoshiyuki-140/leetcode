# 解答

## 解答案1

1. 不要な文字を削除する
2. 逆文字列を作成する
3. 比較する -> 逆にした文字列と元の文字列を比較して一致すれば回文

leetcodeでは、stringsパッケージが使えるので、使っていきます。

```go
func isPalindrome(s string) bool {
	newS := ""
	reversedS := ""
	s = strings.ToLower(s) // 小文字化
	for _, c := range s {
		if (97 <= byte(c) && byte(c) <= 122) || (48 <= byte(c) && byte(c) <= 57) { // a-zのチェックと0-9のチェック
			newS += string(c)
		} // それ以外は飛ばす
	}
	// 逆文字列作成
	for _, c := range newS {
		reversedS = string(c) + reversedS
	}
	if reversedS == newS {
		return true
	}
	return false
}
```

- 空間計算量 : O(n)
- 時間計算量 : O(n)


## 解答案2

二つのポインタ(start,end)を使って、文字列を前と後ろから挟み撃ちにする作戦です。

こっちはstringsパッケージを使わずに純粋にASCIIコード表を参照して小文字変換してみました。

```go
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
```

- 空間計算量 : O(1)
- 時間計算量 : O(n)

こっちは時間計算量が線形のままで、空間計算量がO(1)で済みます