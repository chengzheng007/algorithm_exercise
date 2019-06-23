package main

import "fmt"

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
	for i := n-1; i >= 1; i-- {
		// 选择了第i件物品，当前状态(i,p)从(i-1,p-weight[i])转移过来
		if p-weight[i] >= 0 && state[i-1][p-weight[i]] {
			fmt.Printf("index:%d weight:%d\n", i, weight[i])
			p -= weight[i]
		} // else没有选择i，p不变，下一次循环检测(i-1,p)
	}
	if p > 0 {
		fmt.Printf("index:%d weight:%d\n", 0, weight[0])
	}
	
	return j
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
		for j := maxW-weight[i]; j >= 0; j-- {
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
			v := state[i-1][j]+value[i]
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
		state[i][0] = w[i][0]+temp
		temp = state[i][0]
	}
	temp = 0
	for j := 0; j < n; j++ {
		state[0][j] = w[0][j]+temp
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