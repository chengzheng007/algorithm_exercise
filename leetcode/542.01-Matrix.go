func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    if m == 0 {
        return nil
    }
    n := len(mat[0])
    if n == 0 {
        return nil
    }
    
    // 广搜bfs，从一个点四面发散搜索，复杂度O(n^2xm^2)非常高
    // 一种办法是用带备忘录的广搜，另一种是dp
    // 分别从坐上->右下和右下->坐上做两次遍历搜索，即可完成四个方向上的查找
    // 无论什么方法，一个点最终需要跟四周所有4个点比较，这个定义是递归的
    // 适用于矩阵中每一个位置，所以左上、右下分别扫描一次很巧妙
    
    dp := make([][]int, m)
    dist := m*n // 设置一个默认较大值
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = dist
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                dp[i][j] = 0                
            } else {
                if i > 0 && dp[i-1][j] < dp[i][j] {
                    dp[i][j] = dp[i-1][j] + 1
                }
                if j > 0 && dp[i][j-1] < dp[i][j] {
                    dp[i][j] = dp[i][j-1] + 1
                }
            }
        }
    }
    
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if mat[i][j] == 0 {
                continue
            }
            if i < m-1 && dp[i+1][j] < dp[i][j] {
                dp[i][j] = dp[i+1][j] + 1
            }
            if j < n-1 && dp[i][j+1] < dp[i][j] {
                dp[i][j] = dp[i][j+1] + 1
            }
        }
    }
    
    return dp
}
