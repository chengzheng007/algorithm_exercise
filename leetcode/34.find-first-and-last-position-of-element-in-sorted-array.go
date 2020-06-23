func searchRange(nums []int, target int) []int {
	firstIdx := 0
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				firstIdx = mid
				break
			} else {
				high = mid - 1
			}
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// 无结果
	if low > high {
		return []int{-1, -1}
	}

	// 寻找最后一个相等的下标
	lastIdx := firstIdx
	for ; lastIdx+1 < len(nums); lastIdx++ {
		if nums[lastIdx+1] != target {
			break
		}
	}

	return []int{firstIdx, lastIdx}
}