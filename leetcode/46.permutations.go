/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	size := len(nums)
	allSubSet := make([][]int, 0)
	backtrack(nums, size, 0, &allSubSet)
	return allSubSet
}

// 参考：http://wuchong.me/blog/2014/07/28/permutation-and-combination-realize/
func backtrack(nums []int, size, k int, allSubSet *[][]int) {
	if k == size-1 {
		list := make([]int, size)
		copy(list, nums)
		*allSubSet = append(*allSubSet, list)
		return
	} else {
		// 将下标为k的数开始，分别与它及后面的数交换
		for i := k; i < size; i++ {
			nums[i], nums[k] = nums[k], nums[i]
			backtrack(nums, size, k+1, allSubSet)
			// 交换回来
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
}

