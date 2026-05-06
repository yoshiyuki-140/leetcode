package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	buy, sell := 0, 1
	maxP := 0
	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			profit := prices[sell] - prices[buy]
			maxP = max(maxP, profit)
		} else {
			buy = sell
		}
		sell++
	}
	return maxP
}
