package twosum

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int) // key:val -> target-n:numsIndex
	for numsIndex, num := range nums {
		if val, ok := prevMap[num]; ok {
			return []int{val, numsIndex}
		} else {
			prevMap[target-num] = numsIndex
		}
	}
	return []int{}
}
