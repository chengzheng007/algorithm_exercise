/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	sum := 0
	left := 0
	ans := n + 1
	// 双指针：right右移sum增加，一旦发现sum>=s时，检测
	// 长度是否更小，并且减去最左边的数，最左边的数由left
	// 标记，减去最左边数后left加1，可分析出最对每个元素
	// 最多被访问两边，因此总体上说时间复杂度是O(n)

	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= s {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if ans == n+1 {
		return 0
	}

	return ans
}