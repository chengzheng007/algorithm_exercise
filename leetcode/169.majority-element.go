/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */
func majorityElement(nums []int) int {
	tmp := make(map[int]int)
	for _, num := range nums {
		if times, ok := tmp[num]; ok {
			tmp[num] = times+1
	   	} else {
			tmp[num] = 1
		}
	}
	var majNum,majTimes int
	for num, times := range tmp {
		if times > majTimes {
			majTimes = times
			majNum = num
		}
	}
	return majNum
}

