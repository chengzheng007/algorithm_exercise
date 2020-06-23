package main

import (
	"fmt"
	"math"
)

// 最长公共子序列（不连续，注意不是最长公共子串，子串要求连续）
// Longest Common Sequence
// 在两个字符不同时，只允许做插入或删除
func lcsDFS(word1, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	max := 0
	lcsBacktrack(word1, word2, 0, 0, 0, &max)
	return max
}

func lcsBacktrack(a, b string, i, j, currLcs int, max *int) {
	if i == len(a) || j == len(b) {
		if currLcs > *max {
			*max = currLcs
		}
		return
	}
	if a[i] == b[j] {
		lcsBacktrack(a, b, i+1, j+1, currLcs+1, max)
	} else {
		lcsBacktrack(a, b, i+1, j, currLcs, max) // 删除a[i]或在b[j]前插入a[i]
		lcsBacktrack(a, b, i, j+1, currLcs, max) // 删除b[j]或在a[i]前插入b[j]
	}
}

/**
lcs dp

假设有字符串a、b和 dp([][]int)，dp[i][j]表示到a[i]、b[j]字符位置的最长公共子序列
当a[i] != b[j]时，dp[i][j] = max(
	dp[i-1][j-1],
	dp[i][j-1],
	dp[i-1][j],
)
当a[i] == b[j]时，dp[i][j] = max(
	dp[i-1][j-1]+1,
	dp[i][j-1],
	dp[i-1][j],
)
注意从[i][j-1]或[i-1][j]到[i][j]显然是做过了一些操作的，插入或者删除，因为这两种情形
到[i][j]做了某个操作，所以lcs不能加1
*/

func lcsDP(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)

	if len1 == 0 || len2 == 0 {
		return 0
	}

	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}

	// 初始化text[0]和text2的状态值
	for j := 0; j < len2; j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else if j != 0 {
			dp[0][j] = dp[0][j-1]
		} else {
			// j==0且text1[0] != text2[0]
			dp[0][j] = 0
		}
	}
	// 初始化text1和text2[0]
	for i := 0; i < len1; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else if i != 0 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = 0
		}
	}

	// 动归状态转移
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = maxInt(dp[i-1][j-1]+1, dp[i][j-1], dp[i-1][j])
			} else {
				dp[i][j] = maxInt(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Printf("a:%s, b:%s\n", text1, text2)
	fmt.Println("dp state:")
	for i := 0; i < len1; i++ {
		fmt.Printf("%v\n", dp[i])
	}

	return dp[len1-1][len2-1]
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

func shortestCommonSupersequence(a, b string) string {
	lcs := lcsStr(a, b)
	fmt.Printf("lcsStr:%s\n", lcs)
	if lcs == "" {
		return a + b
	}
	res := make([]byte, 0)
	var i, j int
	for _, c := range lcs {
		for a[i] != byte(c) {
			res = append(res, a[i])
			i++
		}
		for b[j] != byte(c) {
			res = append(res, b[j])
			j++
		}
		res = append(res, byte(c))
		i++ // c为a、b公共字符
		j++
	}
	return string(res) + a[i:] + b[j:]
}

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
