/*
// 方法一：由外到内一层层翻转
func rotate(matrix [][]int) {
    if len(matrix) == 0 {
        return
    }
    n := len(matrix[0])
    if n <= 1 {
        return
    }
    
    for i := 0; i < (n+1)/2; i++ {
        for j := 0; j < n/2; j++ {
            temp := matrix[i][j]
            matrix[i][j] = matrix[n-1-j][i]
            matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
            matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
            matrix[j][n-1-i] = temp
        }
    }
}
*/
// 方法二：先(反斜)对角线翻转，再沿竖线翻转
func rotate(matrix [][]int) {
    if len(matrix) == 0 {
        return
    }
    n := len(matrix[0])
    if n <= 1 {
        return
    }
    // 对角线翻转
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // 竖线翻转
    for i := 0; i < n; i++ {
        for j := 0; j < (n+1)/2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }
}
