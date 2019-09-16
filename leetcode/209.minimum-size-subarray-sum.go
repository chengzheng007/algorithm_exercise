/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */
func minSubArrayLen(s int, nums []int) int {
	size := len(nums)
	if size <= 0 {
		return 0
	}
	sum := 0
	// 使用两个指针，类似滑动窗口
	// 当窗口内的和小于s时，右指针向右移动扩大窗口
	// 当sum>=s时，左指针向右移动，缩小窗口
	// 根据下标差值可知数据长度
	l := 0
	r := -1
	arrLen := size + 1
	for l < size {
		if r+1 < size && sum < s {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			if r-l+1 < arrLen {
				arrLen = r - l + 1
			}
		}
	}
	// 没有找到和等于s的数
	if arrLen == size+1 {
		arrLen = 0
	}
	return arrLen
}

