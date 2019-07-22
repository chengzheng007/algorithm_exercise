/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	// sort by asc
	sort.Sort(sort.IntSlice(nums))
	size := len(nums)
	var (
		low  = 1
		last int
	)
	for i := 0; i < size; i++ {
		if nums[i] <= 0 {
			continue
		}
		// if great than low, find the min,missing positive num
		if nums[i] > low {
			return low
		}
		// if num i less than and it doesn't exist previously(e.x. [1,1,2,2])
		if nums[i] <= low && nums[i] != last {
			low += 1
		}
		last = nums[i]
	}
	return low
}

