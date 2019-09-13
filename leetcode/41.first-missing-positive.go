/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	// 凡是未明确说不能用排序的，都能用排序处理，只要利于解题
	sort.Ints(nums)
	var (
		// 缺失的最小正数
		low = 1
		// 上一个遍历的正数
		lastNum int
	)
	for _, num := range nums {
		// <=0的过滤
		if num <= 0 {
			continue
		}
		if num > low {
			return low
		}
		if num <= low && lastNum != num {
			low++
		}
		lastNum = num
	}
	return low
}

