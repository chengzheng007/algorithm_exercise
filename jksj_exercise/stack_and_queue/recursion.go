package main

import "fmt"

func fibonacci(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1)+fibonacci(n-2)
}

// 全排列
// 假设给定参数list中为唯一整数
func fullPermutation(list []int, cursor, n int) {
	if cursor == n {
		fmt.Printf("permutation:%v\n", list)
	}
	for i := cursor; i < n; i++ {
		list[i], list[cursor] = list[cursor], list[i]
		fullPermutation(list, cursor+1, n)
		list[cursor], list[i] = list[i], list[cursor]
	}
}
