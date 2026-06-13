package missingnumber

func missingNumber(nums []int) int {
	// 0からnまでの操作
	s := len(nums) * (len(nums) + 1) / 2
	// 一つずつ引いていく
	for _, n := range nums {
		s -= n
	}
	return s
}
