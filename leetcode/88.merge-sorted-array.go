/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	tempArr := make([]int, m+n)
	count := 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			tempArr[count] = nums1[i]
			i++
		} else {
			tempArr[count] = nums2[j]
			j++
		}
		count++
	}
	for i < m {
		tempArr[count] = nums1[i]
		i++
		count++
	}
	for j < n {
		tempArr[count] = nums2[j]
		j++
		count++
	}
	copy(nums1, tempArr)
}

