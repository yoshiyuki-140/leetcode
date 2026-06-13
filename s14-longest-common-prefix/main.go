package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := ""
label:
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
				break label
			}
		}
		result += string(strs[0][i])
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
