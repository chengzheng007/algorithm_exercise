func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	if n <= 0 {
		return res
	}

	// 下标表示行数，值表示皇后在第几列
	path := make([]int, n)
	backtrack(n, 0, path, &res)
	return res
}

func backtrack(n, row int, path []int, res *[][]string) {
	if row == n {
		// 将数据格式化为所需的一行行字符串
		pos := make([]string, n)
		r := 0
		for _, c := range path {
			str := ""
			for i := 0; i < n; i++ {
				if i == c {
					str += "Q"
				} else {
					str += "."
				}
			}
			pos[r] = str
			r++
		}
		*res = append(*res, pos)
		return
	}
	for col := 0; col < n; col++ {
		if isOK(n, row, col, path) {
			path[row] = col
			backtrack(n, row+1, path, res)
			path[row] = 0
		}
	}
}

func isOK(n, row, col int, path []int) bool {
	leftUp := col - 1
	rightUp := col + 1
	for r := row - 1; r >= 0; r-- {
		// 同一列上有皇后
		if path[r] == col {
			return false
		}
		// 检测左上角
		if leftUp >= 0 && path[r] == leftUp {
			return false
		}
		// 检测右上角
		if rightUp < n && path[r] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}