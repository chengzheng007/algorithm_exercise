import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	profit := 0
	// 可以不限制交易次数，那么每次只要是高价就出售
	// 比如[1,3,5]，第1天买第3天卖收益：5-1=4
	// 第1天买第2天卖，接着低天又买第3天卖，收益：3-1 + 5-3 = 4，结果一样！
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price > minPrice {
			profit += (price - minPrice)
			minPrice = price
		}
	}
	return profit
}