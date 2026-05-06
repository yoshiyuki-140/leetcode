package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	buy, sell := 0, 1
	maxP := 0
	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			// profit計算
			profit := prices[sell] - prices[buy]
			// maxPの更新可能判断と必要に応じた更新
			maxP = max(maxP, profit)
		} else {
			// 利益が出なかった場合
			// sellの場所にbuyを持っていく
			buy = sell
		}
		// sellの更新
		sell++
	}
	return maxP
}
