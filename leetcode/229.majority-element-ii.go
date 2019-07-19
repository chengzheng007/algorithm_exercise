import "sort"

/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */
func majorityElement(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}
	min := size / 3
	// 首先排序（只要没主动说不能使用排序均可以使用）
	sort.Sort(sort.IntSlice(nums))
	list := make([]int, 0)
	pre := nums[0]
	number := 1
	for i := 1; i < size; i++ {
		if nums[i] == pre {
			number++
		} else {
			if number > min {
				list = append(list, pre)
			}
			number = 1
		}
		pre = nums[i]
	}
	// 数组最后一个元素可能也满足，否则会漏掉
	if number > min {
		list = append(list, pre)
	}
	return list
}

