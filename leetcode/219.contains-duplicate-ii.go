/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	size := len(nums)
	if size <= 1 || k <= 0 {
		return false
	}
	// 以下为暴力方法，时间复杂度为O(n*k)
	// for i := 0; i < size; i++ {
	// 	maxJ := i + k
	// 	if maxJ >= size {
	// 		maxJ = size - 1
	// 	}
	// 	for j := i + 1; j <= maxJ; j++ {
	// 		if nums[i] == nums[j] {
	// 			return true
	// 		}
	// 	}
	// }

	// 一个map记录访问过的元素，最多有k个
	// 相当于一个记录表
	temp := make(map[int]struct{}, 0)
	for i := 0; i < size; i++ {
		if _, ok := temp[nums[i]]; ok {
			return true
		}
		temp[nums[i]] = struct{}{}
		// 保持map中永远只有k个元素
		// 当元素个数大于等于k+1时，随着遍历右移
		// 我们删除距离当前数字k+1的值
		// 这个map(记录表)相当于一个窗口滑动窗口
		if len(temp) == k+1 {
			delete(temp, nums[i-k])
		}
	}
	return false
}

