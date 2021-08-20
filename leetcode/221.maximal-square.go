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
