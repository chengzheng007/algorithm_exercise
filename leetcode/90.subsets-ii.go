import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
func subsetsWithDup(nums []int) [][]int {
	size := len(nums)
	// 因为数组中有重复元素，需要先排序
	sort.Sort(sort.IntSlice(nums))
	subSets := make([][]int, 0)
	currSet := make([]int, 0)
	backtrack2(nums, size, 0, currSet, &subSets)
	return subSets
}

// 此时这种方式not correct，if判断是否加入时
// 求得的是不重复的元素的集合
// func backtrack1(nums []int, size, i int, currSet []int, sets *[][]int) {
// 	if i == size {
// 		*sets = append(*sets, currSet)
// 		return
// 	}

// 	// 不选第i个元素
// 	backtrack1(nums, size, i+1, currSet, sets)

// 	// 当第i个元素不等于第i-1个元素时才被入选
// 	if i < 1 || (i >= 1 && nums[i] != nums[i-1]) {
// 		currSet2 := make([]int, len(currSet)+1)
// 		copy(currSet2, currSet)
// 		currSet2[len(currSet)] = nums[i]
// 		backtrack1(nums, size, i+1, currSet2, sets)
// 	}
// }

func backtrack2(nums []int, size, start int, currSet []int, sets *[][]int) {
	*sets = append(*sets, currSet)
	for i := start; i < size; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		currSet2 := make([]int, len(currSet)+1)
		copy(currSet2, currSet)
		currSet2[len(currSet)] = nums[i]
		backtrack2(nums, size, i+1, currSet2, sets)
		// 删除之前添加的，其实可以不要
		currSet = currSet2[0 : len(currSet2)-1]
	}
}

