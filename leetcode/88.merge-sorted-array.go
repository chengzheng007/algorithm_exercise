/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// 最优解法，一个下标指向最终合并的地方：pos
// i, j分别指向nums1、nums2末尾，从后往前遍历，边扫边合并
// 注意有可能，一个数组中的数比另一个大，所以单独处理循环后
// i/j各自不为0的情况
func merge(nums1 []int, m int, nums2 []int, n int)  {
    pos := m+n-1
    i := m-1
    j := n-1
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[pos] = nums1[i]
            i--
        } else {
            nums1[pos] = nums2[j]
            j--
        }
        pos--
    }
    for i >= 0 {
        nums1[pos] = nums1[i]
        i--
        pos--
    }
    for j >= 0 {
        nums1[pos] = nums2[j]
        j--
        pos--
    }
}

// 耗费了额外O(n)空间
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

