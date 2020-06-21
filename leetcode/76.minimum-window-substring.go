func minWindow(s string, t string) string {
	slen := len(s)
	tlen := len(t)
	if slen == 0 || tlen == 0 {
		return ""
	}

	// value表示相同字符需要出现的个数
	setT := make(map[byte]int)
	for _, char := range t {
		setT[byte(char)]++
	}

	// 独立字符的个数
	required := len(setT)

	window := make(map[byte]int)
	minLen := -1
	minL := 0

	currChar := 0
	l := 0
	r := 0

	// sliding window:通常是两个变量l/r维护一个窗口，r扩展窗口
	// l收缩窗口，当一端扩张或收缩时，另一端固定
	for r < slen {
		window[s[r]] = window[s[r]] + 1

		if setT[s[r]] > 0 && window[s[r]] == setT[s[r]] {
			currChar++
		}

		for currChar == required && l <= r {
			if minLen == -1 || minLen > r-l+1 {
				minLen = r - l + 1
				minL = l
			}
			window[s[l]]--
			if setT[s[l]] > 0 && window[s[l]] < setT[s[l]] {
				currChar--
			}
			l++
		}
		r++
	}

	if minLen == -1 {
		return ""
	}

	return s[minL : minL+minLen]
}