/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	lastStairSteps1 := 1
	lastStairSteps2 := 2
	var currSteps int
	for i := 3; i <= n; i++ {
		currSteps = lastStairSteps2 + lastStairSteps1
		lastStairSteps1 = lastStairSteps2
		lastStairSteps2 = currSteps
	}
	return currSteps
}

