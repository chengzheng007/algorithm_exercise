package main

import "fmt"

// n皇后
// pos 表示存棋子放位置, 下标表示行，值表示列
func nQueen(pos[]int, n, row int) {
	if len(pos) != n {
		return
	}
	if row == n {
		// 多种摆放方式
		printQueen(pos, n)
		return
	}
	for column := 0; column < n; column++ {
		if isOK(pos, n, row, column) {
			pos[row] = column
			nQueen(pos, n, row+1)
		}
	}
}

// 判断(row,column)位置能否放置棋子
func isOK(pos []int, n, row, column int) bool {
	leftUp := column-1
	rightUp := column+1
	for i := row-1; i >= 0; i-- {
		// 同一列(上半部分)已有棋子
		if pos[i] == column {
			return false
		}
		// 左上角已有棋子
		if leftUp >= 0 && pos[i] == leftUp {
			return false
		}
		// 右上角已有棋子
		if rightUp < n && pos[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}

func printQueen(pos []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if pos[i] == j {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// 0-1背包问题
// 选择合适的物品装入背包，总重量不超过needW，每个物品重量不等，不可分割，如何装包使得装入重量最大
// i：处理第i个物品，currW当前重量
// maxW：背包最大可承载重量，inW：装入的重量
func packagesProblem(weight []int, n, i, currW, maxW int, inW *int) {
	if currW == maxW ||  i == n {
		if currW > *inW {
			*inW = currW
		}
		return
	}
	// 选择不装第i个物品
	packagesProblem(weight, n, i+1, currW, maxW, inW)
	// 剪枝处理，如果当前重量+第i个物品重量<=最大重量时，装入，否则不装，避免过多无谓的递归
	if currW+weight[i] <= maxW {
		packagesProblem(weight, n, i+1, currW+weight[i], maxW, inW)
	}
}