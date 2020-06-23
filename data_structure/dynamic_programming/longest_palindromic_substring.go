package main

import "fmt"

// 最长回文子串

func longestPalindromeDFS(s string) string {
	if len(s) <= 1 {
		return s
	}
	var l, r int
	lpsDFS(s, 0, 0, &l, &r)
	fmt.Printf("l:%v, r:%v\n", l, r)
	return s[l : r+1]
}

func lpsDFS(s string, i, j int, l, r *int) {
	if i >= len(s) || j >= len(s) {
		return
	}
	if isPalindrome(s, i, j) {
		if j-i+1 > (*r - *l + 1) {
			*l = i
			*r = j
		}
	}
	lpsDFS(s, i, j+1, l, r)
	lpsDFS(s, i+1, j, l, r)
	lpsDFS(s, i+1, j+1, l, r)
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindromeDP(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([]int, len(s))
	dp[0] = 1

	for i := 0; i < len(s); i++ {
		for j := dp[i] - 1; j < len(s); j++ {
			if isPalindrome(s, i, j) {
				dp[i] = j
			}
		}
	}
	fmt.Printf("dp:%v\n", dp)
	var l, r int
	for i, j := range dp {
		if r-l < j-i {
			l = i
			r = j
		}
	}
	fmt.Printf("l:%v, r:%v\n", l, r)
	return s[l : r+1]
}
