func singleNonDuplicate(nums []int) int {
    return find(nums, len(nums), 0, len(nums)-1)
}

// 只能用递归，因没法确定在哪一边
func find(nums []int, n, low, high int) int {
    if low > high { // 找不到
        return math.MinInt64
    }
    mid := (low+high)/2
    if mid > 0 && mid < n-1 && 
    nums[mid] != nums[mid+1] && 
    nums[mid] != nums[mid-1] {
        return nums[mid]
    } else if mid == 0 { 
        if n == 1 || nums[mid] != nums[mid+1] {
            // 只有一个元素直接返回，或多个元素与mid+1不等
            return nums[mid]
        }
    } else if mid == n-1 {
        if n == 1 || nums[mid] != nums[mid-1] {
            return nums[mid]
        }
    }
    // 寻找左边
    left := find(nums, n, low, mid-1)
    if left != math.MinInt64 {
        return left
    }
    // 继续寻找右边
    return find(nums, n, mid+1, high)
}
