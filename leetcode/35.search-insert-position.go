func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// 落在第1个元素位置且比它小，则返回0
	if high == 0 && target < nums[high] {
		return high
	}

	return high + 1
}