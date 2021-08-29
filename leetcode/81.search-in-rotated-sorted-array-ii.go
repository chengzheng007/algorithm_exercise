// 参考33题题解，注意注释中else if细微差别，因为元素可能相同
// 即使数组被旋转过，我们仍然可以利用这个数组的递增性，使用二分查找。对于当前的中点，
// 如果它指向的值小于等于右端，那么说明右区间是排好序的；反之，那么说明左区间是排好序的。
// 如果目标值位于排好序的区间内，我们可以对这个区间继续二分查找；反之，我们对于另一半区
// 间继续二分查找。
// 注意，因为数组存在重复数字，如果中点和左端的数字相同，如Input: nums=[2,5,6,0,0,1,2], target=0
// 我们并不能确定是左区间全部相同，还是右区间完全相同。在这种情况下，我们可以简单地将左端点右移一位
// 然后继续进行二分查找
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
