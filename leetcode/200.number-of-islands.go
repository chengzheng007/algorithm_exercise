/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
// 需重点回顾
func numIslands(grid [][]byte) int {
	count := 0
	m := len(grid)
	if m == 0 {
		return count
	}
	n := len(grid[0])
	if m == 0 && n == 0 {
		return count
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, m, n, grid)
				count += 1
			}
		}
	}
	return count
}

// dfs函数就是将一个值为1的点周围所有的块，变成0的过程
// 所有的变为0后，就返回了。所以在主函数中count直接+1
// 此处很巧妙，数组的dfs，加上直接变原grid数据（由1变0），无需多申请一个记忆化的数组
func dfs(i, j, m, n int, grid[][]byte) {
	posDirs := [][]int{
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
		[]int{0, 1},
	}
	grid[i][j] = '0'
	for _, dir := range posDirs {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || y < 0 || x >= m || y >= n {
			continue
		}
		if grid[x][y] == '0' {
			continue
		}
		dfs(x, y, m, n, grid)
	}
}