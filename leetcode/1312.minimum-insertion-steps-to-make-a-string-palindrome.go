func minInsertions(s string) int {
	// 假设字符串为s[0, n-1]
	// 如果s[0]==s[n-1]，问题转化为s[1, n-2]
	// 如果s[0]!=s[n-1]，有两种选择：插入s[0]或s[n-1]
	// 如果插入s[0]，问题变为求s[1..n-1]，如果插入s[n-1]
	// 问题变为s[0...n-2]，当然这两者都增加一次插入字符的操作
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for gap := 1; gap < n; gap++ {
		low := 0
		high := gap + low
		for high < n {
			if s[low] == s[high] {
				dp[low][high] = dp[low+1][high-1]
			} else {
				if dp[low][high-1] > dp[low+1][high] {
					dp[low][high] = dp[low+1][high] + 1
				} else {
					dp[low][high] = dp[low][high-1] + 1
				}
			}
			low++
			high++
		}
	}

	return dp[0][n-1]
}