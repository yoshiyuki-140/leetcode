package pascalsTriangle

/*
	constraint(制約)
	1 <= numRows <= 30
*/

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	resultArray := make([][]int, numRows)
	// この行が読み取られた時点で numRows >= 3 が確定する
	// resultArrayを[[1],[1,1]]の状態にする
	resultArray[0] = []int{1}
	resultArray[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		resultArray[i] = make([]int, i+1)
		resultArray[i][0] = 1
		resultArray[i][i] = 1
		for j := 1; j < i; j++ {
			resultArray[i][j] = resultArray[i-1][j-1] + resultArray[i-1][j]
		}
	}
	return resultArray
}
