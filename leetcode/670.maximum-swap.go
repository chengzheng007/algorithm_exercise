/*
 * @lc app=leetcode id=670 lang=golang
 *
 * [670] Maximum Swap
 */
func maximumSwap(num int) int {
	// 参考 solution Approach #2
	list := make([]int, 0)
	temp := num
	for temp > 0 {
		list = append(list, temp%10)
		temp /= 10
	} // 注意:list为数字的逆序排列
	// 翻转后变为顺序
	size := len(list)
	for i := 0; i < size/2; i++ {
		list[i], list[size-i-1] = list[size-i-1], list[i]
	}

	// last数组: key为num数字中单个数，value为该数位于num中的索引位置
	last := make([]int, 10)
	for i := 0; i < 10; i++ {
		last[i] = -1
	}
	// 核心
	for i := 0; i < size; i++ {
		last[list[i]] = i
	}

	for i := 0; i < size; i++ {
		for d := 9; d > list[i]; d-- {
			if last[d] > i {
				list[i], list[last[d]] = list[last[d]], list[i]
				return getNumFromList(list)
			}
		}
	}
	return num
}

func getNumFromList(list []int) int {
	unit := 1
	num := 0
	for i := len(list) - 1; i >= 0; i-- {
		num += list[i] * unit
		unit *= 10
	}
	return num
}

