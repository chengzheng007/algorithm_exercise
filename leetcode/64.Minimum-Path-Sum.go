func minPathSum(grid [][]int) int {
    rows := len(grid)
    if rows == 0 {
        return 0
    }
    cols := len(grid[0])
    if cols == 0 {
        return 0
    }
    
    dp := make([][]int, rows)
    for i := 0; i < rows; i++ {
        dp[i] = make([]int, cols)
    }
    
    // 初始化
    dp[0][0] = grid[0][0]
    for j := 1; j < cols; j++ {
        dp[0][j] = dp[0][j-1] + grid[0][j]
    }
    for i := 1; i < rows; i++ {
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }
    
    
    // dp[i][j] = min{dp[i-1][j],dp[i][j-1]}+grid[i][j]
    // 从1开始是为了方便在for循环统一判断，但是if条件相对复杂
    for i := 1; i < rows; i++ {
        for j := 1; j < cols; j++ {
            if dp[i-1][j] < dp[i][j-1] {
                dp[i][j] = grid[i][j]+dp[i-1][j]
            } else {
                dp[i][j] = grid[i][j]+dp[i][j-1]
            }
        }
    }
    
    return dp[rows-1][cols-1]
}

/*
func minPathSum(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    min := math.MaxInt64
    dfs(grid, len(grid), len(grid[0]), 0, 0, 0, &min)
    return min
}

// 不一定是正方形
func dfs(grid [][]int, row, col, i, j, curr int, min *int) {
    curr += grid[i][j]
    if i == row-1 && j == col-1 {
        if curr < *min {
            *min = curr
        }
        return
    }
    if i < row-1 {
        dfs(grid, row, col, i+1, j, curr, min)
    }
    if j < col-1 {
        dfs(grid, row, col, i, j+1, curr, min)
    }
}
*/
