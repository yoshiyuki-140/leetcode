package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	hashMap := make(map[int]int, len(nums))
	for i, n := range nums {
		complement := target - n
		if _, ok := hashMap[n]; ok {
			result[0] = hashMap[n]
			result[1] = i
			return result
		} else {
			hashMap[complement] = i
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 3}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
