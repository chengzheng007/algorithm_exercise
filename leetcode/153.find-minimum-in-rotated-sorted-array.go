func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1

	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			// 偏大的数字在左边占多数
			low = mid + 1
		} else if nums[mid] < nums[high] {
			// 偏小的数字在右边并占多数
			// 该情况下mid可能是最小的值，所以high赋值为mid
			high = mid
		} else {
			// mid == high，如 [5, 0, 1, 1, 1, 1, 1]
			// 最小值出现在mid左边
			minVal := nums[low]
			for i := low + 1; i <= mid; i++ {
				if nums[i] > minVal {
					minVal = nums[i]
				} else {
					minVal = nums[i]
					break
				}
			}
			return minVal
		}
	}

	return nums[low]
}