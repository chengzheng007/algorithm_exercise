/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */
func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	// 定义状态：state[i]表示索引从0到i处，LIS长度
	state := make([]int, size)
	// 初始化
	state[0] = 1
	maxlen := 1
	for i := 1; i < size; i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && state[j] > maxval {
				maxval = state[j]
			}
		}
		state[i] = maxval + 1
		if state[i] > maxlen {
			maxlen = state[i]
		}
	}
	return maxlen
}

