package main

// 最小距离值
func minDistDfs(n, i, j, curr int, min *int, dist [][]int) {
	curr += dist[i][j]
	if i == n-1 && j == n-1 {
		if curr < *min {
			*min = curr
		}
		return
	}
	if i < n-1 {
		minDistDfs(n, i+1, j, curr, min, dist)
	}
	if j < n-1 {
		minDistDfs(n, i, j+1, curr, min, dist)
	}
}

// 状态转移方程：dp[i][j] := dist[i][j]+min(dp[i-1][j], dp[i][j-1])
// 当前点之可能从左边或上边过来，所以比较左边和上边值谁小，取小的加上当前距离
// 即可作为左上角顶点到当前点的最小距离，下一个顶点以此类推
// 状态转移过程走完，右下角顶点的值即为左上角到该顶点的最小值
func minDistDP(dist [][]int) int {
	size := len(dist)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}

	// 初始化第1行
	sum := 0
	for j := 0; j < size; j++ {
		sum += dist[0][j]
		dp[0][j] = sum
	}
	// 初始化第1列
	sum = 0
	for i := 0; i < size; i++ {
		sum += dist[i][0]
		dp[i][0] = sum
	}

	// 状态转移填充
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			if dist[i-1][j] < dist[i][j-1] {
				dp[i][j] = dist[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = dist[i][j] + dp[i][j-1]
			}
		}
	}

	return dp[size-1][size-1]
}
