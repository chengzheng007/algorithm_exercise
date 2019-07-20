/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 垂直长度、水平长度
	verLen := len(grid)
	horLen := len(grid[0])
	area := 0
	for i := 0; i < verLen; i++ {
		for j := 0; j < horLen; j++ {
			if grid[i][j] == 0 {
				continue
			}
			tempArea := getArea(i, j, verLen, horLen, grid)
			if tempArea > area {
				area = tempArea
			}
		}
	}
	return area
}

func getArea(i, j, verLen, horLen int, grid [][]int) int {
	movePos := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	area := 1 // 当前位置grid[i][j]已经是1
	// 将当前位置置为0，很重要一步
	grid[i][j] = 0
	for _, pos := range movePos {
		newI, newJ := i+pos[0], j+pos[1]
		if newI < 0 || newJ < 0 || newI >= verLen || newJ >= horLen {
			continue
		}
		if grid[newI][newJ] == 0 {
			continue
		}
		// 深度优先搜索
		area += getArea(newI, newJ, verLen, horLen, grid)
	}
	return area
}

