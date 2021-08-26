// 参考33题题解，注意注释中else if细微差别，因为元素可能相同
func search(nums []int, target int) bool {
    length := len(nums)
    low := 0
    high := length-1
    
    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            return true
        }
        
        if nums[low] == nums[mid] {
            low++
	} else if nums[mid] <= nums[high] { // !!注意:此处须<=，否则报错
            if target > nums[mid] && target <= nums[high] {
                low = mid+1
            } else {
                high = mid-1
            }
        } else {
            if target >= nums[low] && target < nums[mid] {
                high = mid-1
            } else {
                low = mid+1
            }
        }
    }
    
    return false
}
