/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size <= 1 {
		return size
	}
	res := 0
	// l,r控制窗口
	l := 0
	r := -1
	temp := make(map[byte]*struct{}, 0)
	for l < size {
		if r+1 < size && temp[s[r+1]] == nil {
			r++ // 扩大窗口右边界
			temp[s[r]] = &struct{}{}
		} else {
			// r到达有边界，或者遇到重复字符，窗口缩小
			delete(temp, s[l])
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

