package main

import (
	"math"
	"fmt"
)

// dp exercise
// https://www.zhihu.com/question/23995189
// 每个阶段只有一个状态->递推；
// 每个阶段的最优状态都是由上一个阶段的最优状态得到的->贪心；
// 每个阶段的最优状态是由之前所有阶段的状态的组合得到的->搜索；
// 每个阶段的最优状态可以从之前某个阶段的某个或某些状态直接得到而不管之前这个状态是如何得到的->动态规划。
//
// 每个阶段的最优状态可以从之前某个阶段的某个或某些状态直接得到
// 这个性质叫做最优子结构；

// 而不管之前这个状态是如何得到的
// 这个性质叫做无后效性。


// fibonacci number
// call: fibonacciNumDP(num)
func fibonacciNumDP(n int) int {
	// 假设n>=0
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	state := make([]int, n+1)

	// 初始化状态
	state[0] = 0
	state[1] = 1

	for i := 2; i <= n; i++ {  // 状态转移
		// 转移方程：f(i) = f(i-1)+f(i-2)
		state[i] = state[i-1]+state[i-2]
	}
	return state[n]
}

func fibonacciNumBT(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	n2 := fibonacciNumBT(n-2)
	n1 := fibonacciNumBT(n-1)
	return n1+n2
}

// 钱币找零
// 找回所选的金额所需要最少的个数
var (
	rmbs = []int{1, 5, 11}
	currNeedNum = math.MaxInt32
)
func findChangeBT(i, cw, num int, rmbs []int, n int, w int) {
	if i == n || cw == w {
		//fmt.Printf("i(%d) == n || cw(%d) == w\n", i, cw)
		if num < currNeedNum {
			currNeedNum = num
		}
		return
	}
	for i := 0; i < n; i++ {
		// 如果当前选择的零钱加上已有零钱小于等于需找回的零钱，继续选择用当前零钱进行找回
		if cw+rmbs[i] <= w {
			//fmt.Printf("choose %d rmb:%d num:%d cw:%d cw+rmbs[i]:%d\n", i, rmbs[i], num+1, cw, cw+rmbs[i])
			findChangeBT(i, cw+rmbs[i], num+1, rmbs, n, w)
		}
	}
}

func findChangeDP(rmbs []int, n, m int) int {
	state :=  make([]int, m+1)
	// 初始化
	state[0] = 0
	
	for i := 1; i <= m; i++ {
		minCost := math.MaxInt32
		for j := 0; j < n; j++ {
			if i-rmbs[j] >= 0 && minCost > state[i-rmbs[j]]+1 {
				minCost = state[i-rmbs[j]]+1
			}
		}
		state[i] = minCost
	}
	fmt.Printf("state:%v\n", state)
	return state[m]
}

// 杨辉三角

// longest increasing substring/最长增长子序列
var (
	series = []int{1,5,3,4,6,9,7,8}
)
func lisDP(list []int) int {
	size := len(list)
	if size <= 1 {
		return size
	}
	state := make([]int, size)
	// 初始化
	for i := range state {
		state[i] = 1
	}

	fmt.Printf("state[0]=%d\n", state[0])
	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if list[j] < list[i] {
				state[i] = int(math.Max(float64(state[i]), float64(state[j]+1)))
			}
		}
		fmt.Printf("state[%d]=%d\n", i, state[i])
	}
	max := state[0]
	for i := 1; i < size; i++ {
		if max < state[i] {
			max = state[i]
		}
	}
	return max
}
