/*
 * @lc app=leetcode id=670 lang=golang
 *
 * [670] Maximum Swap
 */
func maximumSwap(num int) int {
	list := make([]int, 0)
	temp := num
	for temp > 0 {
		n := temp % 10
		list = append(list, n)
		temp /= 10
	}
	size := len(list)
	maxNum := num
	// 题目是要求最终结果是最多交换一次的结果，并不是我们交换一次就完事了，因为
	// 直叫唤一次后的结果不一定是最大的，审题时要注意
	// 暴力枚举，O(n^3)，注意还要从list获取
	// 一个整数，需要遍历整个list，所以是n的3次方
	// 参考leetcode solution Approach #2，非常有特点的解法，复杂度能到O(n)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			list[i], list[j] = list[j], list[i]
			curr := getNumFromList(list)
			if curr > maxNum {
				maxNum = curr
			}
			// 交换后再交换回来，保证实际只发生一次交换
			list[i], list[j] = list[j], list[i]
		}
	}
	return maxNum
}

func getNumFromList(list []int) int {
	unit := 1
	num := 0
	for i := 0; i < len(list); i++ {
		num += list[i] * unit
		unit *= 10
	}
	return num
}

