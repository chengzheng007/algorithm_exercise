func rob(nums []int) int {
    /*
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    dp := make([]int, len(nums)+1)
    // 初始化状态，dp[1]收益最大是抢第一间房子
    dp[1] = nums[0]
    // dp[i]表示抢第i间房子累计最大金额
    // 两种可能：1.不抢第i间，dp[i]=dp[i-1]
    // 2.抢第i间，dp[i] = dp[i-2]+nums[i]
    // 转移方程dp[i] = max{dp[i-1], dp[i-2]+nums[i-1]}
    max := dp[0]
    for i := 2; i <= len(nums); i++ {        
        dp[i] = nums[i-1]+dp[i-2]
        if dp[i-1] > dp[i] {
            dp[i] = dp[i-1]
        }
        if dp[i] > max {
            max = dp[i]
        }
    }
    return max
    */
    
    // 继续优化为无空间（空间压缩）
    size := len(nums)
    if size == 0 {
        return 0
    }
    if size == 1 {
        return nums[0]
    }
    
    var pprev,prev,curr int
    for i := 0; i < size; i++ {
        if pprev+nums[i] > prev {
            curr = pprev+nums[i]
        } else {
            curr = prev
        }
        pprev = prev
        prev = curr
    }
    return curr
}

/*
// dfs
func rob(nums []int) int {
    max := 0
    helper(nums, 0, 0, &max)
    return max
}

func helper(nums []int, k, curr int, max *int) {
    if curr > *max {
        *max = curr
    }
    if k >= len(nums) {
        return
    }
    for i := k; i < len(nums); i++ {
        helper(nums, i+2, curr+nums[i], max)
    }
}
*/
