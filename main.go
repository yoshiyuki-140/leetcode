package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}
	reversedHalf := 0
	for x > reversedHalf {
		// append the end of 'x' to the end of 'reversedHalf'
		reversedHalf = reversedHalf*10 + x%10
		// remove the end of 'x'
		x /= 10
	}
	return ((x == reversedHalf) || (x == reversedHalf/10))
}

func main() {
	x := 121
	result := isPalindrome(x)
	fmt.Println(result)
}
