import "math"

/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */
func longestCommonSubsequence(text1 string, text2 string) int {
	size1 := len(text1)
	size2 := len(text2)
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
	}
	// 核心思想：i、j为text1、text2中的字符下标，当
	// dp[i][j] = dp[i-1][j-1]+1(if text1[i] == text2[j])
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1])(if text1[i] != text2[j])
	// i-1或者j-1表示删除i或j处的字符，或将i处的字符替换成j的
	// 或者j处字符替换成i的
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			// 注意这里是比较i-1与j-1的字符，更新i,j
			// 所以dp[size1][size2]才是表示最终字符的匹配长度
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[size1][size2]
}

// dp下标从1开始
func lcsDP2(text1, text2 string) int {
	size1 := len(text1)
	size2 := len(text2)
	if size1 == 0 || size2 == 0 {
		return 0
	}

	dp := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		dp[i] = make([]int, size2+1)
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[size1][size2]
}

func maxInt(args ...int) int {
	max := math.MinInt64
	for _, num := range args {
		if num > max {
			max = num
		}
	}
	return max
}
