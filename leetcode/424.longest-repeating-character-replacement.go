func characterReplacement(s string, k int) int {
	var (
		slen = len(s)
		// 保存遍历的每个字符的个数
		chars      [26]int
		maxSameCnt int // 相同字符数最多的字符个数
		maxLen     int
		l, r       int
	)

	for r = 0; r < slen; r++ {
		idx := byte(s[r]) - 'A'
		chars[idx]++
		if chars[idx] > maxSameCnt {
			maxSameCnt = chars[idx]
		}
		// 随着r右移，会遇到不同的字符，当总长度减去最多
		// 相同字符个数多于k时，表示了k+1个不同的字符
		// 窗口缩小，左边界右移，左边界的字符数减1
		for r-l+1-maxSameCnt > k {
			if chars[byte(s[l])-'A'] > 0 {
				chars[byte(s[l])-'A']--
			}
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}