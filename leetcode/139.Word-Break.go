// 类似于完全平方数分割问题，这道题的分割条件由集合内的字符串决定，因此在考虑每个分
// 割位置时，需要遍历字符串集合，以确定当前位置是否可以成功分割。注意对于位置 0，需要初
// 始化值为真。
func wordBreak(s string, wordDict []string) bool {
    if len(s) == 0 || len(wordDict) == 0 {
        return false
    }
        
    dp := make([]bool, len(s)+1)
    dp[0] = true // 注意dp[0]，当第一个单词完成匹配时需依赖它的值，必须为true
    
    for i := 1; i <= len(s); i++ {
        for _, word := range wordDict {
            length := len(word)
            if i >= length && s[i-length:i] == word {
                // 注意：如果i-length处不符合，即便此处相等也是false
                // 因为前面已经不符了
                dp[i] = dp[i] || dp[i-length]
            }
        }
    }
    return dp[len(s)]
}
