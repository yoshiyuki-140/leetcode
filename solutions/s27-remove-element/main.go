package removeelement

func removeElement(nums []int, val int) int {

	var i, l uint8 = 0, uint8(len(nums))
	for ; i < l; i++ {
		if nums[i] == val {
			nums = append(nums[:(i)], nums[(i+1):]...)
			i--
		}
	}
	return len(nums)
}
