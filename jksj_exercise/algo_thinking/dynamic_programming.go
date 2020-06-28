package main

import (
	"fmt"
	"math"
)

// 0-1 package problem
// dynamic programming
func knapsack(weight []int, n int, maxW int) int {
	state := make([][]bool, n)
	for i := range state {
		state[i] = make([]bool, maxW+1)
	}
	// 初始化
	state[0][0] = true
	if weight[0] <= maxW {
		state[0][weight[0]] = true
	}
	// 填充状态
	for i := 1; i < n; i++ {
		// 不放第i个物品
		for j := 0; j <= maxW; j++ {
			if state[i-1][j] {
				state[i][j] = true
			}
		}
		// 放入第i个物品
		for j := 0; j <= maxW-weight[i]; j++ {
			if state[i-1][j] {
				state[i][j+weight[i]] = true
			}
		}
	}
	// 结果在最后一行
	j := maxW
	for ; j >= 0; j-- {
		if state[n-1][j] {
			break
		}
	}

	fmt.Printf("选择了以下编号的商品:\n")
	p := j
	for i := n - 1; i >= 1; i-- {
		// 选择了第i件物品，当前状态(i,p)从(i-1,p-weight[i])转移过来
		if p-weight[i] >= 0 && state[i-1][p-weight[i]] {
			fmt.Printf("index:%d weight:%d\n", i, weight[i])
			p -= weight[i]
		} // else没有选择i，p不变，下一次循环检测(i-1,p)
	}
	if p > 0 {
		fmt.Printf("index:%d weight:%d\n", 0, weight[0])
	}

	// 以上只能求出一个解
	var (
		res  [][]int
		path []int
	)
	res = getSolution(weight, state, len(weight), len(weight)-1, j, path, res)
	fmt.Printf("solution\n")
	for _, sol := range res {
		fmt.Printf("%v\n", sol)
	}

	return j
}

// 回溯法求出所有可行的解
func getSolution(weight []int, state [][]bool, n, i, cw int, path []int, res [][]int) [][]int {
	if i == 0 && cw == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return res
	}

	if !state[i][cw] {
		return res
	}

	if i > 0 {
		if state[i-1][cw] { // 没有选第i件物品
			res = getSolution(weight, state, n, i-1, cw, path, res)
		}
		if cw-weight[i] >= 0 && state[i-1][cw-weight[i]] { // 选择了第i件物品
			path = append(path, i)
			res = getSolution(weight, state, n, i-1, cw-weight[i], path, res)
			path = path[:len(path)-1]
		}
	} else if i == 0 && cw > 0 { // state第0行例外检查下有没有可能添加weight中第0件物品
		if cw-weight[i] >= 0 && state[i][cw-weight[i]] {
			path = append(path, i)
			res = getSolution(weight, state, n, i, cw-weight[i], path, res)
			path = path[:len(path)-1]
		}
	}
	return res
}

func knapsack2(weight []int, n int, maxW int) int {
	state := make([]bool, maxW+1)
	// 初始化
	state[0] = true
	if weight[0] <= maxW {
		state[weight[0]] = true
	}
	for i := 1; i < n; i++ {
		// 放置第i个物品
		for j := maxW - weight[i]; j >= 0; j-- {
			if state[j] {
				state[j+weight[i]] = true
			}
		}
	}
	//fmt.Printf("knapsack2 state:%v\n", state)
	for i := maxW; i >= 0; i-- {
		if state[i] {
			return i
		}
	}
	return 0
}

func knapsack3(weight, value []int, n, maxW int) int {
	state := make([][]int, n)
	for i := range state {
		state[i] = make([]int, maxW+1)
		for j := range state[i] {
			state[i][j] = -1
		}
	}

	// 初始化
	state[0][0] = 0
	if weight[0] <= maxW {
		state[0][weight[0]] = value[0]
	}

	for i := 1; i < n; i++ {
		// 不选择第i个物品
		for j := 0; j <= maxW; j++ {
			if state[i-1][j] >= 0 {
				state[i][j] = state[i-1][j]
			}
		}

		// 选择第i个物品
		for j := 0; j <= maxW-weight[i]; j++ {
			v := state[i-1][j] + value[i]
			// 存储比原来大的，小于的丢弃
			if v > state[i][j+weight[i]] {
				state[i][j+weight[i]] = v
			}
		}
	}

	// 取出最大的值
	maxValue := 0
	for j := 0; j <= maxW; j++ {
		if state[n-1][j] > maxValue {
			maxValue = state[n-1][j]
		}
	}
	return maxValue
}

func backTrackMinDist(w [][]int, n, i, j, currDist int, minDist *int) {
	currDist += w[i][j]
	if i == n-1 && j == n-1 {
		if currDist < *minDist {
			*minDist = currDist
			return
		}
	}
	if i < n-1 {
		backTrackMinDist(w, n, i+1, j, currDist, minDist)
		// 如果以这种方式传递，最开始不 += w[i][j]，j=n-1时，最后一格（右下角）是没加上的
		//backTrackMinDist(w, n, i+1, j, currDist+w[i][j], minDist)
	}
	if j < n-1 {
		backTrackMinDist(w, n, i, j+1, currDist, minDist)
	}
}

func dpMinDist(w [][]int, n int) int {
	state := make([][]int, n)
	for i := range state {
		state[i] = make([]int, n)
	}
	// 初始化
	temp := 0
	for i := 0; i < n; i++ {
		state[i][0] = w[i][0] + temp
		temp = state[i][0]
	}
	temp = 0
	for j := 0; j < n; j++ {
		state[0][j] = w[0][j] + temp
		temp = state[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if state[i-1][j] < state[i][j-1] {
				state[i][j] = w[i][j] + state[i-1][j]
			} else {
				state[i][j] = w[i][j] + state[i][j-1]
			}
		}
	}

	return state[n-1][n-1]
}

/* 钱币找零 dp
 * 分别有1元、3元、5元若干张纸币，找9元，最少可以是多少张
 */

func findChangeDfs(m int, coins []int) int {
	min := math.MaxInt64
	findChangeBacktrack(m, coins, 0, 0, 0, &min)
	return min
}

// num：零钱的张数
// 首先探寻回溯类型的算法，可以发现是以 f(当前金额,当前张数) 作为递归函数的核心使用参数
// 以这两个变量画递归树，会出现重复节点
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
			// 不足时，先用当前面额的零钱凑齐
			// 当面额足够时，算算耗费的纸币数，比最小的纸币数要小时记录下来(前提是刚好凑齐额度)
			// 然后再递归用下一个
			findChangeBacktrack(m, coins, i, curr+coins[i], num+1, min)
		}
	}
}

// 钱币找零dp
// 因此在dp解法中，需要围绕 [当前金额，当前张数] 构造状态变量
// 可以用一个[]int整型数组，长度为金额+1
func findChangeDP(m int, coins []int) int {
	dp := make([]int, m+1) // 下标表示金额，值表示该金额对应最小找零纸币数
	// 初始化可以不用，状态转移已经包含零钱等值情况
	// for _, num := range coins {
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
		// for j := 1; j < i; j++ {
		// 	if _, ok := moneyMP[j]; ok {
		// 		num := dp[i-j] + 1
		// 		if num < min {
		// 			min = num
		// 		}
		// 	}
		// }
		for _, coin := range coins {
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
