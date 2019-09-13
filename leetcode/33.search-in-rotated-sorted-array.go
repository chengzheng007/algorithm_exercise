/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
func search(nums []int, target int) int {
	// 先找出翻转后的分界点索引
	size := len(nums)
	lo := 0
	hi := size - 1
	for lo < hi {
		mid := (lo + hi) / 2
		// 当中间数比最后一个值大时，想象
		// 应该是左边的翻转的数比较的多，且都是偏大的数
		// 比分界点右边的大(3,4,5,6,1,2)，也就是说分界点在
		// 偏向右侧的地方，因此lo右移
		// 否则hi左移，左移时hi要等于mid
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	// 分界点为lo的地方
	pivotIdx := lo
	lo = 0
	hi = size - 1
	// 下面开始用二分法定位元素位置
	for lo <= hi {
		mid := (lo + hi) / 2
		// 注意，因为实际的分界点不在中间（当没有翻转时的mid）
		// 因此这里的mid只是索引大小上的中点（非值排列的中点）
		// 那么偏移了中间多少呢，就是这个pivotIdx个位置
		// 所以要用计算的中间点+偏移量，因为可能大于整个数组长度
		// 所以要对长度取模，也就是realMid
		// realMid从数值上来说，才是真正的中点
		realMid := (mid + pivotIdx) % size
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

