// 参考 leetcode 33
func findMin(nums []int) int {
    // 不存在相同元素的写法
    low := 0
    high := len(nums)-1
    for low < high {
        mid := (low+high)/2
        // 中点数值大于高点，说明左边已经有序，且左边数值偏大，往右移
        // 反之，如果中点数值小于右边，说明右边是排好序的
        if nums[mid] > nums[high] {
            low = mid+1
        } else {
            high = mid
        }
    }
    return nums[low]
}
