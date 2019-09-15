/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
func jump(nums []int) int {
	size := len(nums)
	jumps := 0
	curEnd := 0
	curFarthest := 0
	for i := 0; i < size-1; i++ {
		if i+nums[i] > curFarthest {
			curFarthest = i + nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps
}

