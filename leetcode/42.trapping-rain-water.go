/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
func trap(height []int) int {
	size := len(height)
	if size <= 0 {
		return 0
	}
	// leftMax[i]记录从左往右数时，截止下标i处的最大水柱高度值
	leftMax := make([]int, size)
	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	// rightMax[i]记录从右往左数时，截止下标i处的最大水柱高度值
	rightMax := make([]int, size)
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
	// 核心思想：每个下标i处的能装的水，等于截止到i处，左边最大高度值（包括height[i]）
	// 和右边最大高度值（包括height[i]）中较小的那一个
	// 减去自身高度height[i]所得结果
	sum := 0
	for i := 0; i < size; i++ {
		min := leftMax[i]
		if min > rightMax[i] {
			min = rightMax[i]
		}
		sum += min - height[i]
	}
	return sum
}

