/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	if rows == 1 {
		return triangle[0][0]
	}

	elems := (1 + rows) * rows / 2
	dp := make([]int, elems)
	dp[0] = triangle[0][0]

	for i := 1; i < rows; i++ {
		for j := 0; j <= i; j++ {
			// 到上一行为止，总元素个数
			toLastRow := (1 + i) * i / 2
			pos := toLastRow + j
			if j == 0 { // 该行的第一个元素
				dp[pos] = dp[toLastRow-i] + triangle[i][j]
			} else if j == i { // 该行的最后一个元素
				dp[pos] = dp[toLastRow-1] + triangle[i][j]
			} else {
				// 中间元素，从上一行选择头顶元素相邻两个中较大的，前一阶段和较小的
				lastRowLeftElemIdx := toLastRow - i + j - 1
				if dp[lastRowLeftElemIdx] < dp[lastRowLeftElemIdx+1] {
					dp[pos] = dp[lastRowLeftElemIdx] + triangle[i][j]
				} else {
					dp[pos] = dp[lastRowLeftElemIdx+1] + triangle[i][j]
				}
			}
		}
	}

	// 在最后一行中寻找最小数
	start := rows * (rows - 1) / 2
	min := dp[start]
	for i := start + 1; i < elems; i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}

