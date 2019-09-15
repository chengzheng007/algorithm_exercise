/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */
func findLengthOfLCIS(nums []int) int {
	size := len(nums)
	if size <= 0 {
		return 0
	}
	// 可以用dp解法，但空间浪费较大
	preVal := 1
	var currVal int
	max := preVal
	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			currVal = preVal + 1
		} else {
			currVal = 1
		}
		preVal = currVal
		if currVal > max {
			max = currVal
		}
	}
	return max
}

