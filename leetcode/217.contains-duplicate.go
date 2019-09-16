/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
func containsDuplicate(nums []int) bool {
	temp := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := temp[num]; ok {
			return true
		} else {
			temp[num] = struct{}{}
		}
	}
	return false
}

