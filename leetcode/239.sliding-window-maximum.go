/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow(nums []int, k int) []int {    
    if len(nums) == 0 {
        return []int{}
    }
    res := make([]int, len(nums)-k+1) // 可以提前分配好空间
    // dq为double queue，存储的是元素下标（非元素值）
    dq := make([]int, 0)    
    
    for i, num := range nums {
        // 保持窗口单调递减
        // 移除队尾比当前小的元素
        length := len(dq)
        for length > 0 && nums[dq[length-1]] < num {
            dq = dq[:length-1]
            length--
        }
        
        dq = append(dq, i)
        
        // 移除多出窗口数目的元素，即双端队列的队头元素
        length = len(dq)
        // 或判断下标：i-dq[0] == k
        if length > 0 && i-dq[0]+1 > k {
            dq = dq[1:]
        }
        
        if i >= k-1 {
            res[i-k+1] = nums[dq[0]]
            // res = append(res, nums[dq[0]])
        }
    }

    return res
}
