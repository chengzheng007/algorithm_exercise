func removeDuplicates(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }
    
    // index指向上一次最多重复两次的元素下标
    // 如果赋值，需先给index+1
    index := 0 
    curr := 1
    for i := 1; i < n; i++ {
        if nums[i] == nums[index] {
            if curr < 2 {
                index++
                nums[index] = nums[i]
                curr++
            }
        } else {
            index++
            nums[index] = nums[i]
            curr = 1
        }
    }
    return index+1
}
