package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	ans := removeDuplicates(nums)
	fmt.Println(ans)
	nums = []int{1,1,2}
	ans = removeDuplicates(nums)
	fmt.Println(ans)
}

func removeDuplicates(nums []int) int {
	count := 0
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		} else {
			m[nums[i]] = true
			nums[count] = nums[i]
			count++
		}
	}
	nums = nums[:len(m)]
	return len(m)
}
