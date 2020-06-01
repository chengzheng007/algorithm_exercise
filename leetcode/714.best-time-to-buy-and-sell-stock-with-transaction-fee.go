import "math"

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := math.MaxInt64
	profit := 0
	for _, price := range prices {
		if price < minPrice { // 寻找交易低价
			minPrice = price
		} else if price > (minPrice + fee) { // 首先只有price-minPrice-fee 才会产生收益，else进入到卖出操作
			profit += price - minPrice - fee
			// 但是每次卖出都会产生交易费用fee，如果
			// 没有交易费用可以放心无数次买卖，交易费
			// 用就是成本，意味着不能随便卖出，那么如
			// 果本次交易的成本，在下一次卖出操作时能
			// 收回来，那么下一次交易就是划算的，也就
			// 是说这样的交易才会使得总利润越来越多
			// 的。为了找到这样的价格，那么下一次交易
			// 的买卖价格差需要足够大，如果下一交易时
			// 低价比 当前价格-交易费用(price -
			// fee)还低，那么这一笔交易就能挣回上次
			// 交易产生的交易费，就能使得总利润和是在
			// 增长的
			minPrice = price - fee
		}
	}

	return profit
}