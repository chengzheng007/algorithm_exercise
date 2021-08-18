func numberOfArithmeticSlices(nums []int) int {
    size := len(nums)
    if size < 3 {
        return 0
    }
    // 注意题目要求：1.至少3个数; 2.subarray需要是连续的，不能断开
    // 比如[1,2,3,4,5]中，[1,3,5]不能算，因为在整个数组中他们是
    // 非【连续子序列】，因此只需考虑前后的数字。
    // 另外对于用dp解法来说，dp[i]是以第i个数结尾的子数组数量
    // 而题设要求的等差数列可以在任意地方终结，需把dp中所有数求和
    // 才是所有的符合条件子数组总数量
    
    dp := make([]int, size)
    res := 0
    for i := 2; i < size; i++ {
        if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
            dp[i] = dp[i-1]+1
        }
        res += dp[i]
    }
    return res
    /*
    // 无需额外space的写法
    var curr, res int
    for i := 2; i < size; i++ {
        if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
            curr++
            res += curr
        } else {
            curr = 0
        }
    }
    return res
    */
}
