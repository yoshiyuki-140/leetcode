# 解答

hashMapを使った解法

hashMapのキーにtarget-nを格納し、hashMapのバリューにその時点のnumsのインデックスを格納して探索する。


```go
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int) // key:value -> index:target-n
	for i, n := range nums {
		if indexVal, ok := hashMap[n]; ok {
			return []int{indexVal, i}
		}
		hashMap[target-n] = i
	}
	return []int{}
}
```