func rob(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		// dp[i-2]+nums[i]表示不抢第i-1间屋子，抢劫第i间
		// dp[i-1]表示抢第i-1间屋子的总金额，不能再抢第i间
		// 比较二者谁大，取大的即可
		// 进阶：最多保留i-2的数即可，不用全部使用
		m := nums[i]
		if i >= 2 {
			m += dp[i-2]
		}
		if dp[i-1] > m {
			dp[i] = dp[i-1]
		} else {
			dp[i] = m
		}
	}

	return dp[n-1]
}