/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 0 {   // 长度为0的单独处理
        return 0
    }
    
    index := 0  // index记录上一个相同元素的起始位置
    for i := 1; i < n; i++ {
        // 思想：一旦发现下一个不同的元素，覆盖到index后面的位置
        if nums[i] != nums[index] { 
            index++
            nums[index] = nums[i]
        }
    }
    // 返回唯一元素的个数
    return index+1
}
