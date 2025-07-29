package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	array := strings.Split(s, "")
	length := len(s)
	for i := 0; i < length/2; i++ {
		if array[i] != array[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	x := 1211
	result := isPalindrome(x)
	fmt.Println(result)
}
