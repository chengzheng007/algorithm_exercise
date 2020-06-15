func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// 状态转移
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
			} else if i == 0 || j == 0 {
				// 在边界的一行或一列上，前面有障碍，后面的都是0
				if j == 0 && obstacleGrid[i][j] == 0 && obstacleGrid[i-1][j] == 1 {
					obstacleGrid[i][j] = 1
				} else if i == 0 && obstacleGrid[i][j] == 0 && obstacleGrid[i][j-1] == 1 {
					obstacleGrid[i][j] = 1
				} else {
					obstacleGrid[i][j] = 0
				}
			} else if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}