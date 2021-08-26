func isValidSudoku(board [][]byte) bool {
    // 三个数组bitmap，行、列、每个九宫格的数据是否重复
    n := len(board)
    rows := make([]int, n) // 9行
    cols := make([]int, n) // 9列
    boxes := make([]int, n) // 9个小九宫格
    
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if board[r][c] == '.' {
                continue
            }
            
            val := 1 << (board[r][c]-'0'-1)
            if rows[r] & val > 0 { // 该位置已出现过
                return false
            }
            rows[r] |= val
            
            if cols[c] & val > 0 {
                return false
            }
            cols[c] |= val
            
            boxIdx := (r/3)*3 + c/3
            if boxes[boxIdx] & val > 0 {
                return false
            }
            boxes[boxIdx] |= val
        }
    }
    return true
}
