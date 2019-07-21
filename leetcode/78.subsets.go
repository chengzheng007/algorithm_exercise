/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
func subsets(nums []int) [][]int {
	size := len(nums)
	subSets := make([][]int, 0)
	currSet := make([]int, 0)
	// backtrack(0, size, nums, currSet, &subSets)
	backtrack2(nums, size, 0, currSet, &subSets)
	return subSets
}

// 回溯算法版本1-双递归实现
// func backtrack(i, size int, nums, currSet []int, sets *[][]int) {
// 	if i == size {
// 		*sets = append(*sets, currSet)
// 		return
// 	}
// 	// 不选第i个元素
// 	backtrack(i+1, size, nums, currSet, sets)

// 	// 选第i个元素
// 	// 不能共用一个currSet变量，下面的append可能对上面的有影响，因go语言的实现机制导致
// 	currSet2 := make([]int, len(currSet)+1)
// 	copy(currSet2, currSet)
// 	currSet2[len(currSet)] = nums[i]
// 	backtrack(i+1, size, nums, currSet2, sets)
// }

// 回溯算法版本2-加一个循环实现
func backtrack2(nums []int, size, start int, currSet []int, sets *[][]int) {
	*sets = append(*sets, currSet)
	for i := start; i < size; i++ {
		currSet2 := make([]int, len(currSet)+1)
		copy(currSet2, currSet)
		currSet2[len(currSet)] = nums[i]
		backtrack2(nums, size, i+1, currSet2, sets)
		currSet = currSet2[0 : len(currSet2)-1]
	}
}

