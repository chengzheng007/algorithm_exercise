import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	size := len(nums)
	ret := nums[0] + nums[1] + nums[size-1]
	for i := 0; i < size-2; i++ {
		l := i + 1
		r := size - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			} else if sum > target {
				// 能这样比较大小，序列一定要是有序递增的，往左移才会使结果减小
				// 只要题目没说不能用排序的，便可以用排序，以此方便解答
				r--
			} else {
				l++
			}
			temp1 := sum - target
			if temp1 < 0 {
				temp1 *= -1
			}
			temp2 := ret - target
			if temp2 < 0 {
				temp2 *= -1
			}
			if temp1 < temp2 {
				ret = sum
			}
		}
	}
	return ret
}

