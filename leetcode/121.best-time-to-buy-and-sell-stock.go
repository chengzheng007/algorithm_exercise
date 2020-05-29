package leetcode

import "math"

func maxProfit(prices []int) int {
	// O(n^2)暴力搜索每一种差价
	// size := len(prices)
	// max := 0
	// for i := 0; i < size-1; i++ {
	//     for j := i+1; j < size; j++ {
	//         p := prices[j]-prices[i]
	//         if p > max {
	//             max = p
	//         }
	//     }
	// }

	var (
		minPrice  = math.MaxInt64
		maxProfit = 0
	)
	for i := 0; i < len(prices); i++ {
		// 遍历时先找最小价格，如果出现比当前记录的售卖价格还小时，先持有，循环继续
		// 一旦发现比当前记录售卖价格大的价格，检测售出收益是否大于当前记录的收益值，如果比它大则卖出，否则仍持有，寻找下一个更大的卖价，注意此时最小买价并未改变，所以后面如果有更大价格是能得到卖出更高价格的收益
		// 思路就是：寻找【最小】的售价，当遍历到价格大于它并且收益更多时卖出
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
