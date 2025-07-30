package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := ""
	baseList := strings.Split(strs[0], "") // 先頭文字列基準にしてからそれを一文字ずつのスライスに変換する
label:
	// base	には一文字ずつstrs[0]を分割した文字列が入る
	for i, base := range baseList {
		// 以下のループをすべて抜けたら `result += base` できる
		for j := 1; j < len(strs); j++ {
			str := strings.Split(strs[j], "")
			if len(str) <= i { // strの中の文字にアクセスできるかチェック
				break label
			}
			// 以下がメインとなる条件式で、すべてfalseの場合はresult += base
			if string(str[i]) != base {
				break label
			}
		}
		result += base
	}
	return result
}

func main() {
	// Test case 1
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"
	output := longestCommonPrefix(strs)
	if expected == output {
		fmt.Println("SUCCESS => expected:", expected, ",output:", output)
	} else {
		fmt.Println("FAILED => expected:", expected, ",output:", output)
	}

	// Test case 2
	strs = []string{"dog", "racecar", "car"}
	expected = ""
	output = longestCommonPrefix(strs)
	if expected == output {
		fmt.Println("SUCCESS => expected:", expected, ",output:", output)
	} else {
		fmt.Println("FAILED => expected:", expected, ",output:", output)
	}
	// Test case 3
	strs = []string{"dog", "dogs"}
	expected = "dog"
	output = longestCommonPrefix(strs)
	if expected == output {
		fmt.Println("SUCCESS => expected:", expected, ",output:", output)
	} else {
		fmt.Println("FAILED => expected:", expected, ",output:", output)
	}
	// Test case 4
	strs = []string{""}
	expected = ""
	output = longestCommonPrefix(strs)
	if expected == output {
		fmt.Println("SUCCESS => expected:", expected, ",output:", output)
	} else {
		fmt.Println("FAILED => expected:", expected, ",output:", output)
	}
	// Test case 5
	strs = []string{"a"}
	expected = "a"
	output = longestCommonPrefix(strs)
	if expected == output {
		fmt.Println("SUCCESS => expected:", expected, ",output:", output)
	} else {
		fmt.Println("FAILED => expected:", expected, ",output:", output)
	}
	// Test case 6
	strs = []string{"ab", "a"}
	expected = "a"
	output = longestCommonPrefix(strs)
	if expected == output {
		fmt.Println("SUCCESS => expected:", expected, ",output:", output)
	} else {
		fmt.Println("FAILED => expected:", expected, ",output:", output)
	}
}
