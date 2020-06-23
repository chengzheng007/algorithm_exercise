package main

import (
	"fmt"
	"math"
)

// 钱币找零dp
// minNum[currMoney] = min(minNum[currMoney-coin[0]]+1, minNum[currMoney-coin[1]]+1 ... minNum[currMoney-coin[n-1]]+1)
func findChangeDP(m int, money []int) int {
	dp := make([]int, m+1) // 下标表示金额，值表示该金额对应最小找零纸币数
	// 初始化可以不用
	// for _, num := range money {
	// 	// 初始化状态
	// 	if num <= m {
	// 		dp[num] = 1
	// 	}
	// }

	// 填充剩下状态
	for i := 1; i <= m; i++ {
		min := math.MaxInt64
		// 不用找所有，只需寻找i-money中的金额的下标，因为i-money中金额在money中一定会出现
		// 那么dp[i]一定等于dp[i-money]+1，并且最小的也诞生在这些里面
		// 反过来讲，如果不是跟当前金额之差刚好为零钱金额的，找零只会更多
		// for j := 1; j < i; j++ {
		// 	if _, ok := moneyMP[j]; ok {
		// 		num := dp[i-j] + 1
		// 		if num < min {
		// 			min = num
		// 		}
		// 	}
		// }
		for _, coin := range money {
			if i-coin >= 0 && dp[i-coin]+1 < min {
				min = dp[i-coin] + 1
			}
		}
		dp[i] = min
	}

	if dp[m] == math.MaxInt64 {
		return -1 // 无解
	}

	return dp[m]
}

// 一组零钱[1,3,5]，个数不限，找零n元最多有多少种组合方法（DP）
func findChangeMaxDP(coins []int, n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= n; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	fmt.Printf("findChangeMaxDP dp:%v\n", dp)
	return dp[n]
}

// 一组零钱[1,3,5]，个数不限，找零n元最多有多少种组合方法（回溯）
func findChangeMaxDFS(coins []int, n int) int {
	return findChangeMaxBacktrack(coins, n, 0, 0, 0)
}

func findChangeMaxBacktrack(coins []int, n, k, curr int, num int) int {
	if curr >= n {
		// 统计最大的组合次数，只要等凑出的面额等于金额时，最多次数就加1
		// 和找零的硬币个数不一样，硬币个数是累加
		if curr == n {
			num++
		}
	}
	for i := k; i < len(coins); i++ {
		if curr+coins[i] <= n {
			// 如果当前面额零钱不够，继续使用当前面额找下去
			// 直到拼凑出的额度大于n或者零钱种类用完
			num = findChangeMaxBacktrack(coins, n, i, curr+coins[i], num)
		}
	}
	return num
}

func findChangeDfs(m int, coins []int) int {
	min := math.MaxInt64
	findChangeBacktrack(m, coins, 0, 0, 0, &min)
	return min
}

// num：零钱的张数
func findChangeBacktrack(m int, coins []int, k, curr, num int, min *int) {
	// 零钱找到了最后一种或已找到符合数量的零钱
	if curr >= m {
		if curr == m && num < *min {
			*min = num
		}
		return
	}
	for i := k; i < len(coins); i++ {
		if curr+coins[i] <= m {
			findChangeBacktrack(m, coins, i, curr+coins[i], num+1, min)
		}
	}
}
