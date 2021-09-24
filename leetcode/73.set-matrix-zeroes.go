func setZeroes(matrix [][]int)  {
    // O(1)的空间复杂度：利用第1行、第1列的矩阵作为标记
    // 检测到某元素为0时，将它所处的第1行、第1列赋为0
    // 然后再次扫描整个矩阵，只要该元素所处的第1行/第1列为0时即赋为0
    m := len(matrix)
    n := len(matrix[0])
    
    // 检测处理第1行、第1列的元素，有可能这里存在0
    rowHasZero := false
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            rowHasZero = true
            break
        }
    }
    
    colHasZero := false
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            colHasZero = true
            break
        }
    }
    // 注意：对原本第1行/列存在0的情况，必须在下面处理逻辑的前面
    // 因为下面更改了第1行/列的原始值
    
    // 更改第1行/列值为0，作为标记
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0 // 第1列
                matrix[0][j] = 0 // 第1行
            }
        }
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    // end:更改第1行/列值为0，作为标记

    // 处理原始第1行、第1列有0的情况
    if rowHasZero {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    if colHasZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
    // 单独拎出来，循环次数更少

}


func setZeroes(matrix [][]int)  {
    // O(m+n)空间复杂度：记录某行或某列是否存在0
    // 若存在直接将该行/列全部更新为0
    m := len(matrix)
    n := len(matrix[0])
    rowHasZero := make([]bool, m)
    colHasZero := make([]bool, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rowHasZero[i] = true
                colHasZero[j] = true
            }
        }
    }
    
    for i := 0; i < m; i++ {
        if !rowHasZero[i] {
            continue
        }
        // 将该行填充为0
        for j := 0; j < n; j++ {
            matrix[i][j] = 0
        }
    }
    for j := 0; j < n; j++ {
        if !colHasZero[j] {
            continue
        }
        // 将列填充0
        for i := 0; i < m; i++ {
            matrix[i][j] = 0
        }
    }
}
