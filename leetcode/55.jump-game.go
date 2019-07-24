/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
func canJump(nums []int) bool {
	size := len(nums)
	lastPos := size - 1
	// 从最后一个位置往前找
	for i := size - 1; i >= 0; i-- {
		// 如果当前位置下标+加上当前位置可跳跃步数
		// 大于等于当前位置值，说明在i处可达lastPos
		// 注意当前下标一定可以到达当前下标，哪怕其对应的步数为0
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	// 如果最后的位置可达，lastPos一定会被减小到0
	return lastPos == 0
}
