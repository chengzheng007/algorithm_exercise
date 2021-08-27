/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// 从上到下的备忘录+递归方式，不够直观
// bottom-up dp，minpath[k][i]表示第k层第i个元素的最小路径和
// minpath[k][i] = min(minpath[k+1][i], minpath[k+1][i+1]) + triangle[k][i]
// 由于k+1层只依赖第k层，可进一步压缩为，对于第k层
// minpath[i] = min(minpath[i],minpath[i+1])+triangle[k][i]
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    if n == 0 {
        return 0
    }
    
    minpath := make([]int, n)
    // 初始化，最后一行作为dp初始值
    for i := 0; i < n; i++ {
        minpath[i] = triangle[n-1][i]
    }
    
    // 从底向上
    for layer := n-2; layer >= 0; layer-- {
        for i := 0; i <= layer; i++ {
            if minpath[i] < minpath[i+1] {
                minpath[i] = minpath[i] + triangle[layer][i]
            } else {
                minpath[i] = minpath[i+1] + triangle[layer][i]
            }
        }
    }
    
    return minpath[0]
}

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

