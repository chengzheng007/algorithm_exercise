/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */
func rotate(nums []int, k int) {
	size := len(nums)
	// k可以超过size，这时相当于翻转到头，再翻转一次，k对size取模
	k %= size
	// 首先翻转整个数组
	reverse(nums, 0, size-1)
	// 翻转前k个元素
	reverse(nums, 0, k-1)
	// 翻转最后的size-k个元素
	reverse(nums, k, size-1)
}

func reverse(list []int, start, end int) {
	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}
}

