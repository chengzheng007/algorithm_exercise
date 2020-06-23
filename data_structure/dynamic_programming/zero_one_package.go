package main

import "fmt"

// 0-1背包问题
// 不超过w的最大重量
// 动归将求解过程视为一个多阶段的状态，这个多阶段本质上由回溯算法(dfs)的每一层递归进化而来，将每一层重复的节点合并
// 合并后可得到一个数量上少于(当递归层数很多时就远少于)递归节点数的状态值
// 下一层的的状态由上一层推演（或者说填充）而来
func zeroOnePackageDP(w int, weight []int) int {
	size := len(weight)
	state := make([][]bool, size)
	for i := 0; i < size; i++ {
		state[i] = make([]bool, w+1)
	}

	// 初始化第一层状态
	state[0][0] = true // 不选第1个物品
	if weight[0] <= w {
		state[0][weight[0]] = true // 选择第1个物品
	}
	for i := 1; i < size; i++ {
		// 不选第i件物品，直接copy第i-1层的状态
		for j := 0; j <= w; j++ {
			if state[i-1][j] {
				state[i][j] = true
			}
		}
		// 选第i件物品，在第i-1层已选项的基础上，加上第i件物品（注意重量不能超标）
		for j := 0; j <= w-weight[i]; j++ {
			if state[i-1][j] {
				state[i][j+weight[i]] = true
			}
		}
	}
	// 最大重量出现在最后一层
	maxw := 0
	for i := w; i >= 0; i-- {
		// 靠后的大
		if state[size-1][i] {
			maxw = i
			break
		}
	}
	var (
		res  [][]int
		path []int
	)
	res = getSolution(weight, state, len(weight), len(weight)-1, maxw, path, res)
	fmt.Printf("solution\n")
	for _, sol := range res {
		fmt.Printf("%v\n", sol)
	}

	fmt.Printf("选择了以下编号的商品:\n")
	p := maxw
	for i := len(weight) - 1; i >= 1; i-- {
		// 选择了第i件物品，当前状态(i,p)从(i-1,p-weight[i])转移过来
		if p-weight[i] >= 0 && state[i-1][p-weight[i]] {
			fmt.Printf("index:%d weight:%d\n", i, weight[i])
			p -= weight[i]
		} // else没有选择i，p不变，下一次循环检测(i-1,p)
	}
	if p > 0 {
		fmt.Printf("index:%d weight:%d\n", 0, weight[0])
	}

	return maxw
}

func getSolution(weight []int, state [][]bool, n, i, cw int, path []int, res [][]int) [][]int {
	if i == 0 && cw == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, path)
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

func zeroOnePackageDP2(w int, weight []int) int {
	state := make([]bool, w+1)

	state[0] = true
	if weight[0] <= w {
		state[weight[0]] = true
	}

	for i := 1; i < len(weight); i++ {
		for j := w - weight[i]; j >= 0; j-- {
			// 不选第i件时的状态的转移，实际已包含在这里
			// 比如state[0]=true，该循环会重置state[0+weight[1]]，因为j>=0
			// 至于state[weight[0]]=true，如果weight[0]+weight[1]<=w，那么循环
			// state[j+weight[i]]也能循环到，除非weight[0]>w，这也不符合要求了
			if state[j] {
				state[j+weight[i]] = true
			}
		}
	}
	maxw := 0
	for i := w; i >= 0; i-- {
		if state[i] {
			maxw = i
			break
		}
	}
	return maxw
}

func zeroOnePackageDP3(w int, weight []int, value []int) int {
	if len(weight) != len(value) {
		panic("weight num not equal value num")
	}
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, w+1)
	}

	dp[0][0] = 0        // 不选第1件物品
	if weight[0] <= w { // 选第1件物品
		dp[0][weight[0]] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		// 不选第i件物品，i+1件的价值与其一致
		for j := 0; j <= w; j++ {
			if dp[i-1][j] > 0 {
				dp[i][j] = dp[i-1][j]
			}
		}
		// 选第i件物品，在第i-1件的j上+第i件的物品重量，并比较加完后的价值是否比当前大
		// 如果大则
		for j := 0; j <= w-weight[i]; j++ {
			if dp[i-1][j]+value[i] > dp[i][j+weight[i]] {
				dp[i][j+weight[i]] = dp[i-1][j] + value[i]
			}
		}
	}
	maxv := 0
	for i := w; i >= 0; i-- {
		if dp[len(weight)-1][i] > maxv {
			maxv = dp[len(weight)-1][i]
		}
	}
	return maxv
}
