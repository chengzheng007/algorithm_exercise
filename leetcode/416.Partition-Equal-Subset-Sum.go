func canPartition(nums []int) bool {
    // 看成背包问题，如何转换？
    // 题设是选取一些数，剩下的数的和等于选取的数
    // 如果所有元素和为sum的话，相当于在数组中选取一部分元素
    // 使他们的和为sum/2，相当于每个元素价值（数字大小即
    // 重量）相等背包问题
    sum := 0
    for _, num := range nums {
        sum += num
    }
    // 和为奇数分为两半不可能相等
    if sum&1 == 1 { // 等同于sum%2 == 1
        return false
    }
    size := len(nums)
    sum /= 2
    
    dp := make([][]bool, size+1)
    for i := 0; i <= size; i++ {
        dp[i] = make([]bool, sum+1)
        // 初始化，后面计算需要
        dp[i][0] = true
    }
    
    
    // dp[i][j]:数组中下标为n，且和为j时的状态
    // true表示可以满足，false不行
    for i := 1; i <= size; i++ {
        for j := 1; j <= sum; j++ {
            dp[i][j] = dp[i-1][j]
            // dp[i-1][j]: 不选第i个数，结果与dp[i-1][j]相同
            // dp[i-1][j-nums[i-1]]，选第i个数，结果与j-nums[i-1]
            // 相同，取二者的或运算结果
            // 不能直接dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]] 
            // 会漏算
            if j >= nums[i-1] {
                dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]] 
                
            }
           
        }
    }
        
    return dp[size][sum]
}
