package main

import (
	"fmt"
	"strings"
)

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

func main() {
	// Test case 1
	s := "III"
	expected_output := 3
	output := romanToInt(s)
	if expected_output == output {
		fmt.Println("SUCCESS : Test case 1", expected_output, output)
	} else {
		fmt.Println("FAILD : Test case 1", expected_output, output)
	}
	// Test case 2
	s = "LVIII"
	expected_output = 58
	output = romanToInt(s)
	if expected_output == output {
		fmt.Println("SUCCESS : Test case 2", expected_output, output)
	} else {
		fmt.Println("FAILD : Test case 2", expected_output, output)
	}
	// Test case 3
	s = "MCMXCIV"
	expected_output = 1994
	output = romanToInt(s)
	if expected_output == output {
		fmt.Println("SUCCESS : Test case 3", expected_output, output)
	} else {
		fmt.Println("FAILD : Test case 3", expected_output, output)
	}
}
