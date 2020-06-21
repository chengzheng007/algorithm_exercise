/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	set := make(map[byte]bool)
	i, j := 0, 0
	max := 0
	for i < n && j < n {
		if !set[s[j]] {
			if max < j-i+1 {
				max = j - i + 1
			}
			set[s[j]] = true
			j++
		} else {
			if _, ok := set[s[i]]; ok {
				set[s[i]] = false
			}
			i++
		}
	}
	return max
}