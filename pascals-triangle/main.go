package pascalsTriangle

/*
	constraint(制約)
	1 <= numRows <= 30
*/

func generate(numRows int) [][]int {
	/*
		1. numRowsが1であれば、[[1]]を返却する
		1. numRowsが2であれば、[[1],[1,1]]を返却する
		1. numRows >= 3の時、resultArray:[][]intに[[1],[1,1]]を格納する.
		1. for文でnumRows-2分ループをかける -> i:index の初期値は2
			1. ひとつ前の配列からループをかける
	*/
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	// この行が読み取られた時点で numRows >= 3 が確定する
	resultArray := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		appendArray := make([]int, i+1)
		appendArray[0] = 1
		appendArray[i] = 1
		for j := 1; j < i; j++ {
			appendArray[j] = resultArray[i-1][j-1] + resultArray[i-1][j]
			// fmt.Printf("i:%d, j:%d, appendArray[j]:%d = resultArray[i-1][j-1]:%d + resultArray[i-1][j]:%d\n", i, j, appendArray[j], resultArray[i-1][j-1], resultArray[i-1][j])
		}
		resultArray = append(resultArray, appendArray)
	}
	return resultArray
}
