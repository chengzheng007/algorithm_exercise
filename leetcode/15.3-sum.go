/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	size := len(nums)
	list := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))
	for i := 0; i < size-2; i++ {
		target := -nums[i]
		front := i + 1
		back := size - 1
		for front < back {
			sum := nums[front] + nums[back]
			switch {
			case sum < target:
				front++
			case sum > target:
				back--
			default:
				// find one solution
				list = append(list, []int{nums[i], nums[front], nums[back]})
				// Processing duplicates of Number 2
				for front+1 < back && nums[front+1] == nums[front] {
					front++
				}
				// Processing duplicates of Number 3
				for back-1 > front && nums[back] == nums[back-1] {
					back--
				}
				front++
				back--
			}
		}
		for i+1 < size && nums[i] == nums[i+1] {
			i++
		}
	}
	return list
}

