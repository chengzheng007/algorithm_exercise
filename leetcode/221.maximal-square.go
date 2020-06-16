import "math"

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 最大位置为1的个数
	maxSideLen := 0

	// dp[i][j]表示以maxtrix[i-1][j-1]作为右下角的点
	// 所形成的最大的正方形的【边的长度】
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = minInt(dp[i][j-1], dp[i-1][j-1], dp[i-1][j]) + 1
				if dp[i][j] > maxSideLen {
					maxSideLen = dp[i][j]
				}
			}
		}
	}

	return maxSideLen * maxSideLen
}

func minInt(x, y, z int) int {
	min := math.MaxInt64
	if x < min {
		min = x
	}
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}