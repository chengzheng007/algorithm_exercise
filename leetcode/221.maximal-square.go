/*
对于在矩阵内搜索正方形或长方形的题型，一种常见的做法是定义一个二维dp数组，其中
dp[i][j]表示满足题目条件的、以(i, j)为右下角的正方形或者长方形的属性。对于本题，则表示
以(i, j)为右下角的全由1构成的最大正方形面积。如果当前位置是0，那么dp[i][j]即为0；如果
当前位置是1，我们假设dp[i][j] = k2，其充分条件为dp[i-1][j-1]、dp[i][j-1]和dp[i-1][j]的值必须
都不小于(k-1)2，否则 (i, j) 位置不可以构成一个边长为k的正方形。同理，如果这三个值中的
的最小值为k-1，则(i, j) 位置一定且最大可以构成一个边长为k的正方形
*/
func maximalSquare(matrix [][]byte) int {
    // dp[i][j]表示以坐标点(i,j)为右下角的图形
    // 能形成最大正方形的变长，可得方程
    // dp(i,j) = min(dp[i-1,j-1],dp[i,j-1],dp[i-1,j]),matrix[i][j] = 1
    // dp[i][j] = 0, matrix[i][j] != 1
    
    rows := len(matrix)
    if rows == 0 {
        return 0
    }
    cols := len(matrix[0])
    if cols == 0 {
        return 0
    }
    
    dp := make([][]int, rows+1)
    for i := 0; i <= rows; i++ {
        dp[i] = make([]int, cols+1)
    }
    
    maxlen := 0
    for i := 1; i <= rows; i++ {
        for j := 1; j <= cols; j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = getMin(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
            if dp[i][j] > maxlen {
                maxlen = dp[i][j]
            }
        }
    }
    return maxlen*maxlen
}

func getMin(x, y, z int) int {
    min := x
    if y < min {
        min = y
    }
    if z < min {
        min = z
    }
    return min
}

/*
func maximalSquare(matrix [][]byte) int {
    // broute force
    rows := len(matrix)
    if rows == 0 {
        return 0
    }
    cols := len(matrix[0])
    if cols == 0 {
        return 0
    }
    
    maxlen := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '0' {
                continue
            }
            sqlen := 1
            // 对角线式前进
            loop:
            for sqlen+i < rows && sqlen+j < cols {
                // 检测行
                for p := j; p <= j+sqlen; p++ {
                    if matrix[i+sqlen][p] == '0' {
                        break loop
                    }
                }
                // 检测列
                for p := i; p <= i+sqlen; p++ {
                    if matrix[p][j+sqlen] == '0' {
                        break loop
                    }
                }
                sqlen++
            }

            if maxlen < sqlen {
                maxlen = sqlen
            }
        }
    }
    return maxlen * maxlen // square number
}
*/
