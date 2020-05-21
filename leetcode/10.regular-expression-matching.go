func isMatch(s string, p string) bool {
	// 参考solution回溯解法
	slen := len(s)
	plen := len(p)
	if plen == 0 {
		return slen == 0
	}
	// 先检查第一个字是否匹配
	firstMatch := slen > 0 && (p[0] == s[0] || p[0] == '.')
	// 检测第二个正则是否为*
	if plen >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
