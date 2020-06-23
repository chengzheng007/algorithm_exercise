package main

// 钢条切割

// 一组钢条长度、售卖价格对应关系，如何对长度为n寸的钢条进行切割，是收益最大（只能裁成整数）
// price []int，下标表示尺寸，值表示售卖价格（price[0]=0表示长度为0售价为0）
// 从1到n的每一个尺寸都有对应价格
// 思路：先切割成1寸，剩余的长度递归继续先切1寸，一直切到剩余的n'存，在切割完的基础上回退尝试2寸，直到尝试所有情况
func cutRodDFS(price []int, n, dep int) int {
	// 由于切割尺寸从1开始，价格尺寸对应关系下标也需从1开始
	// 从1到n的每一个尺寸都有对应价格
	if len(price) < n+1 {
		return 0
	}
	// 剩余切割长度为0
	if n == 0 {
		return 0
	}
	q := 0
	for i := 1; i <= n; i++ {
		// fmt.Printf("cutRodDFS(price, n:%d, n-i:%d) i:%d dep:%d\n", n, n-i, i, dep+1)
		// 逐个尺寸切割，找出该尺寸最大的切割方案利润
		profit := price[i] + cutRodDFS(price, n-i, dep+1)
		if q < profit {
			q = profit
		}
	}
	// fmt.Printf("get max profit, n:%d, profit:%d, dep:%d\n", n, q, dep)
	return q
}

// 定义数组dp，dp[i](i>=1)表示尺寸为i的钢条切割获得的最大利润
func cutRodDP(price []int, n int) int {
	if len(price) < n+1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		maxp := 0
		// 每个子问题需从子问题中寻找一个最优解
		for j := 1; j <= i; j++ {
			p := price[j] + dp[i-j]
			if maxp < p {
				maxp = p
			}
		}
		dp[i] = maxp
	}
	return dp[n]
}
