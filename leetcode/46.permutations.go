/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	backtrack(nums, n, 0, &result)
	return result
}

func backtrack(nums []int, n, start int, result *[][]int) {
	if start == n {
		// 这里必须临时复制一份，因为各个函数会对nums造成影响
		// 递归返回后，将数据换回来，导致result中
		// 的排列都是换回来的结果
		temp := make([]int, n)
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	// 外层：i从start循环到n，将这之间的数与下标为start的数交换
	// 每换一次递归处理本次换完后的情况
	for i := start; i < n; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		backtrack(nums, n, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
