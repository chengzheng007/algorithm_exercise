func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	j := 0
	for i := 0; i < tLen; i++ {
		if j < sLen && t[i] == s[j] {
			j++
		}
		if j == sLen {
			break
		}
	}

	return j == sLen
}