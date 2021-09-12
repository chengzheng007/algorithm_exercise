// 看对应题解，找数字规律
func nextPermutation(nums []int)  {
    n := len(nums)
    if n <= 1 {
        return
    }
    
    // 假设从右往左为升序，找到第一个降序的数
    i := n-2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    
    // 从右往左，找到第一个比i所在数字大的数
    if i >= 0 {
        j := n-1
        for nums[j] <= nums[i] {
            j--
        }
        // 交换分区数和第一个比分区数大的数
        nums[i], nums[j] = nums[j], nums[i]
    }
    
    // reverse partiton index's right part 
    i = i+1
    j := n-1
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
    
}
