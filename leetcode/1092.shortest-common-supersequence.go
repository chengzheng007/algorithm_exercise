func shortestCommonSupersequence(str1 string, str2 string) string {
	// 获取最长公共子串
	lcs := lcsStr(str1, str2)
	if lcs == "" {
		return str1 + str2
	}
	res := make([]byte, 0)
	var i, j int
	// 遍历lcs，如果在str1中没有的，追加到res末尾，直到遇到相同的字符，如果在str2中没有的，追加到res末尾，知道遇到相同的
	for _, c := range lcs {
		for str1[i] != byte(c) {
			res = append(res, str1[i])
			i++
		}
		for str2[j] != byte(c) {
			res = append(res, str2[j])
			j++
		}
		// 将相同的追加到res末尾，子串下标加1
		res = append(res, byte(c))
		i++ // c为a、b公共字符
		j++
	}
	// 不要忘了子串可能有剩下字符
	return string(res) + str1[i:] + str2[j:]
}

// 动归求得最长公共子序列
func lcsStr(text1, text2 string) string {
	size1 := len(text1)
	size2 := len(text2)
	if size1 == 0 || size2 == 0 {
		return ""
	}

	dp := make([][]string, size1+1)
	for i := 0; i <= size1; i++ {
		dp[i] = make([]string, size2+1)
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(text1[i-1])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[size1][size2]
}