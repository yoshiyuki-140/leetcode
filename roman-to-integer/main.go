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
	result := 0
	array := strings.Split(s, "")
	length := len(array)
	c := ""
	for i := 0; i < length; {
		if length-i >= 2 {
			c = array[i] + array[i+1]
			switch c {
			case "CM":
				result += 900
				i += 2
				continue
			case "CD":
				result += 400
				i += 2
				continue
			case "XC":
				result += 90
				i += 2
				continue
			case "XL":
				result += 40
				i += 2
				continue
			case "IX":
				result += 9
				i += 2
				continue
			case "IV":
				result += 4
				i += 2
				continue
			default:
				c = array[i]
				i++
				switch c {
				case "M":
					result += 1000
				case "D":
					result += 500
				case "C":
					result += 100
				case "L":
					result += 50
				case "X":
					result += 10
				case "V":
					result += 5
				case "I":
					result += 1
				}

			}
		} else { // when length == 1
			c = array[i]
			i++
			switch c {
			case "M":
				result += 1000
			case "D":
				result += 500
			case "C":
				result += 100
			case "L":
				result += 50
			case "X":
				result += 10
			case "V":
				result += 5
			case "I":
				result += 1
			}
		}
	}
	return result
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
