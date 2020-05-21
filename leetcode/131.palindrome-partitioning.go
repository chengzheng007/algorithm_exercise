func partition(s string) [][]string {
	res := make([][]string, 0)
	var path []string
	dfs(s, path, &res)
	return res
}

func dfs(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := 1; i <= len(s); i++ {
		if isPalindromeString(s[:i]) {
			path = append(path, s[:i])
			dfs(s[i:], path, res)
			path = path[:len(path)-1]
		}
	}
}

func isPalindromeString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}