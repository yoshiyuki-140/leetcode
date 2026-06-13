package singlenumber

func singleNumber(nums []int) int {
	XOR := nums[0]
	for i := 1; i < len(nums); i++ {
		XOR ^= nums[i]
	}
	return XOR
}
