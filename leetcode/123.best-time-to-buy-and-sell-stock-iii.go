import "math"

func maxProfit(prices []int) int {
	//参考discuss
	var oneByOneSell, twoBuyTwoSell int
	oneBuy := math.MaxInt64
	twoBuy := math.MaxInt64
	for _, price := range prices {
		oneBuy = min(oneBuy, price)                      // 寻找最小买价
		oneByOneSell = max(oneByOneSell, price-oneBuy)   // 最大化第一次卖出收入
		twoBuy = min(twoBuy, price-oneByOneSell)         // 寻找最小第二次买入，在第一次收益基础上
		twoBuyTwoSell = max(twoBuyTwoSell, price-twoBuy) // 最大化第二次卖出
	}
	// 如果不限交易次数，那么可以直接定义一个最小价格，每次检查比该价格大，且收益比当前最大
	// 收益还大时，做一次交易，借助这个思想，这里使用了四个变量，每次均寻找最小买入价格，最大收益
	return twoBuyTwoSell
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
//maxProfit := 0
//makeProfit(prices, len(prices), 0, 0, -1, 0, &maxProfit)
// 回溯
func makeProfit(prices []int, n, k, times, buy, curr int, maxProfit *int) {
    if curr > *maxProfit {
        *maxProfit = curr
    }
    if times >= 2 {
        return
    }
    for i := k; i < n; i++ {
        if buy == -1 {
            makeProfit(prices, n, i+1, times, i, curr, maxProfit)
        } else if buy != -1 && i != buy && prices[i] > prices[buy] {
            earn := prices[i] - prices[buy]
             makeProfit(prices, n, i+1, times+1, -1, curr+earn, maxProfit)
        }
    }
}
*/