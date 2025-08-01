package validParentheses

import (
	"fmt"
	"strings"
)

// スタック構造定義
type Stack struct {
	sData []string
}

func (pStack *Stack) Push(s string) {
	pStack.sData = append(pStack.sData, s)
}

func (pStack *Stack) Pop() (string, error) {

	sDataLength := len(pStack.sData)
	if sDataLength <= 0 {
		return "", fmt.Errorf("%s", "スタックにデータがありません.")
	}

	lastS := pStack.sData[sDataLength-1]
	pStack.sData = pStack.sData[:sDataLength-1]

	return lastS, nil
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	sList := strings.Split(s, "")
	stack := Stack{}
	// 括弧数チェック
	for _, str := range sList {
		switch str {
		case "(", "{", "[":
			stack.Push(str)
		case ")":
			popedStr, err := stack.Pop()
			if err != nil || popedStr != "(" {
				return false
			}
		case "}":
			popedStr, err := stack.Pop()
			if err != nil || popedStr != "{" {
				return false
			}
		case "]":
			popedStr, err := stack.Pop()
			if err != nil || popedStr != "[" {
				return false
			}
		}
	}
	return len(stack.sData) == 0
}
