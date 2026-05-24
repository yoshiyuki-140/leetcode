package reversebits

func reverseBits(n int) int {
	N := uint32(n)
	result := uint32(0)
	// 32-1回ループを回す
	for i := 0; i < 31; i++ {
		// Nの最後の行が1であればresultの最後の行に1を足す
		if N&0b1 == 0b1 {
			result |= 0b1
		}
		// resultを1bit左シフトする
		// 次の処理結果を格納する隙間を作る
		result <<= 1
		// Nを1bit右シフトする
		// 処理対象のビットを末尾に持ってくる
		N >>= 1
	}
	return int(result)
}
