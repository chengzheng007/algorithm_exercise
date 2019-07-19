/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
	size := len(nums)
	var i, j int
	// 由于中间有j到末尾的移动，所以时间复杂度是O(n^2)
	// 参考solution，两个下标可以做到O(n)
	for i = 0; i < size; i++ {
		// 寻找与起始一段元素不同的索引位置
		for j = i + 1; j < size; j++ {
			if nums[j] != nums[i] {
				break
			}
		}
		// 将j开始的元素覆盖到i+1的位置
		if j > i+1 {
			for p := j; p < size; p++ {
				nums[p-(j-i)+1] = nums[p]
			}
			size -= j - i - 1
		}
	}
	return size
}

