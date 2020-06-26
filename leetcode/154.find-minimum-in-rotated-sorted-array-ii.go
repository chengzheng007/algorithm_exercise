func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] <= nums[high] {
			// 如果mid同时小于low和high，mid不变
			// 缩小low、high范围
			// 否则去[low,mid]里边找最小数
			if nums[mid] <= nums[low] {
				high--
				low++
			} else {
				high = mid
			}
		}
	}
	return nums[mid]
}