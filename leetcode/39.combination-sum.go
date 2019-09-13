/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	nums := make([]int, 0)
	bt(candidates, target, 0, nums, &result)
	return result
}

func bt(candidates []int, target, start int, nums []int, result *[][]int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum == target {
		*result = append(*result, nums)
		nums = []int{}
	}
	for i := start; i < len(candidates); i++ {
		nums1 := make([]int, len(nums)+1)
		copy(nums1, nums)
		nums1[len(nums)] = candidates[i]
		// 如果和比目标值小，继续递归用当前值计算
		if sum < target {
			bt(candidates, target, i, nums1, result)
		}
	}
}

