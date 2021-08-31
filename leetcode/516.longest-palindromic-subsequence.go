func longestPalindromeSubseq(s string) int {
    // dp[i][j]表示s[i]到s[j]之间（包含i,j下标）最长的回文子序列
    // i为左边索引，j为右边索引
    // dp[i][j] = dp[i+1][j-1]+2, if s[i]==s[j]
    // dp[i][j] = max{dp[i+1][j], dp[i][j-1]}
    // dp[i+1][j]：左边删除一个字符，dp[i][j-1]:右边删除一个字符
    dp := make([][]int, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]int, len(s))
    }
    
    for i := len(s)-1; i >= 0; i-- {
        dp[i][i] = 1 // 单独一个字符，视为回文
        for j := i+1; j < len(s); j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]+2
            } else {
                if dp[i+1][j] > dp[i][j-1] {
                    dp[i][j] = dp[i+1][j]
                } else {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }
    return dp[0][len(s)-1] // s[0]-s[len-1]
}
