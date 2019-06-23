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