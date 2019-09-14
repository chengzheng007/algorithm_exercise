/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	// dp[i]表示nums[0:i]之间和最大的连续序列
	// 该序列必须以nums[i]结尾
	// 状态转移方程：dp[i] = nums[i] + (dp[i-1] > 0 ? dp[i-1] : 0)
	size := len(nums)
	dp := make([]int, size)
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < size; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

