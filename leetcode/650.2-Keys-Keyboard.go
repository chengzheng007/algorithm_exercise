// 不同于以往通过加减实现的动态规划，这里需要乘除法来计算
// 位置，因为粘贴操作是倍数增加的。我们使用一个一维数组dp，
// 其中位置i表示延展到长度i的最少操作次数。对于每个位置
// j，如果j可以被i整除，那么长度i就可以由长度j操作得到，其
// 操作次数等价于把一个长度为1的A延展到长度为i/j。因此我们
// 可以得到递推公式dp[i] = dp[j] + dp[i/j]
func minSteps(n int) int {
    if n <= 1 {
        return 0
    }
    
    dp := make([]int, n+1)
    h := int(math.Sqrt(float64(n)))
    
    for i := 2; i <= n; i++ {
        dp[i] = i
        for j := 2; j <= h; j++ {
            if i%j == 0 { // i能整除j，
                dp[i] = dp[j]+dp[i/j]
                break
            }
        }
    }
    
    return dp[n]
}
