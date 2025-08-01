package pascalstriangle

/*
	constraint(制約)
	1 <= numRows <= 30
*/

func generate(numRows int) [][]int {
	resultArray := make([][]int, numRows)
	resultArray[0] = []int{1}

	for i := 1; i < numRows; i++ {
		resultArray[i] = make([]int, i+1)
		resultArray[i][0] = 1
		resultArray[i][i] = 1
		for j := 1; j < i; j++ {
			resultArray[i][j] = resultArray[i-1][j-1] + resultArray[i-1][j]
		}
	}
	return resultArray
}
