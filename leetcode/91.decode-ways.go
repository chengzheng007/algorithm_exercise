
func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 && i+1 <= n && validStr(s[i:i+1]) {
			dp[i] = 1
		} else {
			if i+1 <= n && validStr(s[i:i+1]) {
				dp[i] += dp[i-1]
			}
			if i >= 1 && i+1 <= n && validStr(s[i-1:i+1]) {
				if i >= 2 {
					dp[i] += dp[i-2]
				} else {
					dp[i] += 1
				}
			}
		}
	}

	return dp[n-1]
	// return dfs(s, len(s), 0, 0)
}

// 回溯法
func dfs(s string, n, pos, curr int) int {
	// 当检测到最后一个字符时，才算一次有效解码
	if pos >= n {
		curr++
		return curr
	}

	if pos+1 <= n && validStr(s[pos:pos+1]) {
		curr = dfs(s, n, pos+1, curr)
	}

	if pos+2 <= n && validStr(s[pos:pos+2]) {
		curr = dfs(s, n, pos+2, curr)
	}

	return curr
}

func validStr(str string) bool {
	n := len(str)

	if n == 1 && str[0] >= '1' && str[0] <= '9' {
		return true
	} else if n == 2 {
		switch str[0] {
		case '1':
			if str[1] >= '0' && str[1] <= '9' {
				return true
			}
		case '2':
			if str[1] >= '0' && str[1] <= '6' {
				return true
			}
		}
	}
	return false
}

