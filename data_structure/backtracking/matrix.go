package main

import (
	"math"
	"fmt"
)

var (
	// 矩阵中包含一些权值
	matrix = [][]int32{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	matrixSize = 4
)

// problem：从左上角走到右下角，每次只能向右或向下走，到达终点时经过的权值之和最小为多少

var (
	minDist int32 = math.MaxInt32
)
// backtracking
func minDistBT(i, j int, curDist int32, matrix [][]int32, n int) {
	curDist += matrix[i][j]
	//fmt.Printf("i:%v j:%v curDist:%v\n", i, j, curDist)
	// 仅当i、j均等于n时，到达右下角
	if i == n-1 && j == n-1 {
		if curDist < minDist {
			minDist = curDist
		}
		return
	}
	// 往下走
	if i < n-1 {
		minDistBT(i+1, j, curDist, matrix, n)
	}
	// 往右走
	if j < n-1 {
		minDistBT(i, j+1, curDist, matrix, n)
	}
}

func minDistDP(matrix [][]int32, n int) {
	minDst := make([][]int32, n)
	for i := 0; i < n; i++ {
		minDst[i] = make([]int32, n)
	}

	// 初始化值-第1行
	var sum int32 = 0
	for j := 0; j < n; j++ {
		sum += matrix[0][j]
		minDst[0][j] = sum
	}

	// 初始化值-第1列
	sum = 0
	for i := 0; i < n; i++ {
		sum += matrix[i][0]
		minDst[i][0] = sum
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			val1 := minDst[i-1][j]+matrix[i][j]
			val2 := minDst[i][j-1]+matrix[i][j]
			if val1 < val2 {
				minDst[i][j] = val1
			} else {
				minDst[i][j] = val2
			}
		}
	}
	fmt.Printf("min dstination value:%d\n", minDst[n-1][n-1])
}