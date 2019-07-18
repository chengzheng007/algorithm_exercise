/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	/*
		Time Complexity: O(n^2)
		size := len(height)
		maxContain := 0
		for i := 1; i < size; i++ {
			for j := 0; j < i; j++ {
				currContain := 0
				if height[i] > height[j] {
					currContain = height[j] * (i - j)
				} else {
					currContain = height[i] * (i - j)
				}
				if currContain > maxContain {
					maxContain = currContain
				}
			}
		}
		return maxContain
	*/

	// 两个计步器，O(n)可以解决
	start := 0
	end := len(height) - 1
	maxContain := 0
	for start < end {
		currContain := 0
		if height[start] > height[end] {
			currContain = height[end] * (end - start)
			// 哪一个高度较小，则移动高度较小的一端
			// 原理是两根线中，高度越短容量就越小，线的高度限制了容量大小
			// 那么同样高度的线，如果两根线相隔越远，容量就会越大
			end--
		} else {
			currContain = height[start] * (end - start)
			start++
		}
		if currContain > maxContain {
			maxContain = currContain
		}
	}
	return maxContain
}



