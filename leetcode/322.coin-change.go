import "math"

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */
func coinChange(coins []int, amount int) int {
	n := len(coins)
	if n == 0 || amount <= 0 {
		return 0
	}
	// dp[i]表示金额i的最小找零数量
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		minCost := math.MaxInt32
		for _, coin := range coins {
			// dp[i-coin]是已经得到的最小找零数
			// dp[i-coin]+1表示，金额i需要的在最小找零数上+1
			// 如果他比当前需要找零的数小，就赋为这个较小的数
			if i-coin >= 0 && dp[i-coin]+1 < minCost {
				minCost = dp[i-coin] + 1
			}
		}
		// 循环所有金额，找到可找零数中最小的
		dp[i] = minCost
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

