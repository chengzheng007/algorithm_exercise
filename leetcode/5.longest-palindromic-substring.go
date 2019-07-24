/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
func longestPalindrome(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
	}
	// 初始化包含在循环填充值的过程中
	var start, maxLen int
	for i := size - 1; i >= 0; i-- {
		for j := i; j < size; j++ {
			// 当前首尾的字母相同
			// j-i<3：当子串长度小于3时，不需要检测他的子串（不检测边界）
			// dp[i+1][j-1]，当长度大于等于3时，[i,j]的子串也必须是回文串
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			// 如果当前位置[i,j]是回文串，比较当前回文串长度与历史长度
			// 如果比历史记录长，则更新新的回文串的起点下标和长度值
			if dp[i][j] && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return s[start : start+maxLen]
}

