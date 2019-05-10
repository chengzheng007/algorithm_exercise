package main

import "fmt"

var (
	n = 5
	weight = []int{2,2,4,6,3}
	w = 9
	maxW = 0
)

func ksnapsack(i, cw int) {
	if cw == w || i == n {
		if cw == w {
			fmt.Printf("cw == w, i:%d cw:%d\n", i, cw)
		}
		if cw > maxW {
			maxW = cw
		}
		return
	}
	fmt.Printf("not put i(%d) cw(%d)\n", i, cw)
	ksnapsack(i+1, cw)
	if weight[i]+cw <= w {
		fmt.Printf("put i(%d) cw(%d)+weight[%d](%d)=%d\n", i, cw, i, weight[i], cw+weight[i])
		ksnapsack(i+1, weight[i]+cw)
	} else {
		fmt.Printf(">w(9) not put i(%d) cw(%d)+weight[%d](%d)=%d\n", i, cw, i, weight[i], cw+weight[i])
	}
}

func ksnapsack2(weight []int, n, w int) {
	state := make([][]bool, n)
	for i := 0; i < n; i++ {
		state[i] = make([]bool, w+1)
	}

	state[0][0] = true
	state[0][weight[0]] = true
	fmt.Printf("0 level: %v\n", state[0])

	for i := 1; i < n; i++ { // 动态规划转移
		for j := 0; j <= w; j++ {
			if state[i-1][j] { // 不把第i个物品放入背包
				state[i][j] = true
			}
		}
		fmt.Printf("[1] %d level: %v weight[%d]:%d\n", i, state[i], i, weight[i])
		for j := 0; j <= w-weight[i]; j++ { // 把第i个物品放入背包
			if state[i-1][j] {
				state[i][j+weight[i]] = true
			}
		}
		fmt.Printf("[2] %d level: %v weight[%d]:%d\n", i, state[i], i, weight[i])
	}

	// 输出结果
	for i := w; i >= 0; i-- {
		if state[n-1][i] {
			fmt.Printf("weight:%v\n", i)
		}
	}
}

func ksnapsack3(weight []int, n, w int) {
	state := make([]bool, w+1)
	// 第一行特殊处理，即第一个物品放与不放先计算出来
	state[0] = true
	state[weight[0]] = true

	for i := 1; i < n; i++ {
		// 从后往前计算，只能否则j为较小的值的时候，计算后面的值为true的话，j自增循环到后面
		// 该值又为true，又会接着计算state，造成重复计算
		for j := w-weight[i]; j >= 0; j-- {
			if state[j] {
				state[j+weight[i]] = true
			}
		}
	}
	// 输出结果
	fmt.Println("ksnapsack3 cal w:")
	for i := w; i >= 0; i-- {
		if state[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// dp 动态规划
var (
	value = []int{3,4,8,9,6}
)
// 满足背包中重量限制的前提下，价值最大
func ksnapsack4(weight, value []int, n, w int) int {
	state := make([][]int, n)
	// 初始化
	for i := 0; i < n; i++ {
		state[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			state[i][j] = -1
		}
	}

	// 初始化边界值，第一层值（由底层往上推）
	state[0][0] = 0
	state[0][weight[0]] = value[0]

	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {  // 不放第i个物品，也就是将第i-1层摆放的拷贝过来
			if state[i-1][j] >= 0 {
				state[i][j] = state[i-1][j]
			}
		}

		for j := 0; j <= w-weight[i]; j++ { // 放第i个物品
			if state[i-1][j] >= 0 {
				v := state[i-1][j]+value[i]
				if v > state[i][j+weight[i]] {  // 如果第i个物品的放入后的价值比原来的大，则放入，否则略过
					state[i][j+weight[i]] = v
				}
			}
		}
	}

	// 打印最后一层所有组合出的价值，其中最大的便是所求的
	fmt.Println("ksnapsack4 cal value:")
	max := state[n-1][0]
	fmt.Printf("%d ", max)
	for i := 1; i <= w; i++ {
		fmt.Printf("%d ", state[n-1][i])
		if max < state[n-1][i] {
			max = state[n-1][i]
		}
	}
	fmt.Println()
	return max
}