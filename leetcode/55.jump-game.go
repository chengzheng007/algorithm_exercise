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

// 回溯解法（会超时）
func canJumpDFS(nums []int) bool {
	return jumpJudge(nums, 0)
}

func jumpJudge(nums []int, pos int) bool {
	if pos == len(nums)-1 {
		return true
	}
	nextMaxPos := pos + nums[pos]
	if nextMaxPos > len(nums)-1 {
		nextMaxPos = len(nums) - 1
	}
	for nextPos := pos + 1; nextPos <= nextMaxPos; nextPos++ {
		if jumpJudge(nums, nextPos) {
			return true
		}
	}
	return false
}

// dp解法
func canJumpDP(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	// dp[i]表示i坐标是否可达末尾
	// -1:未知  0:不可达  1:可达
	dp := make([]int, n)
	for i := 0; i < n-1; i++ {
		dp[i] = -1
	}
	// 最后一个元素永远可达，可以自己跳跃到自己
	dp[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		maxPos := nums[i]
		if maxPos > n-1 {
			maxPos = n - 1
		}
		// 从i处，看看跳到的下一个位置是否是一个
		// 好位置，如果是退出本次循环
		// 不是的话沿着i+1处继续找，当然找的最远处
		// 是max(nums[i], n-1)
		for j := i + 1; j <= maxPos; j++ {
			if dp[j] == 1 {
				dp[i] = 1
				break
			}
		}
	}

	if dp[0] == 1 {
		return true
	}

	return false
}