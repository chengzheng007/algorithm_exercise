import (
	"strconv"
	"strings"
)

// 非回溯方法
func restoreIpAddresses(s string) []string {
	var res []string
	n := len(s)
	for i := 1; i < 4 && i <= n-3; i++ {
		for j := i + 1; j < i+4 && j <= n-2; j++ {
			for k := j + 1; k < j+4 && k <= n-1; k++ {
				if !isValid(s[0:i]) {
					continue
				}
				if !isValid(s[i:j]) {
					continue
				}
				if !isValid(s[j:k]) {
					continue
				}
				if !isValid(s[k:]) {
					continue
				}
				str := s[0:i] + "." + s[i:j] + "." + s[j:k] + "." + s[k:]
				res = append(res, str)
			}
		}
	}
	return res
}

// 回溯方法
// 思路：分成4块，每一次取1-3个字符组成一个ip数字，劈分后如果长度不够，如不足4个字符的跳过
// 每一次劈分成功后递归进行下一次检测，当到第4块时，劈分后的s应该为空
func backtrack(s string, index int, path []string, res *[]string) {
	if index == 4 {
		if s == "" {
			*res = append(*res, strings.Join(path, "."))
		}
		return
	}
	for i := 1; i < 4; i++ {
		if i > len(s) {
			return
		}
		if isValid(s[:i]) {
			path = append(path, s[:i])
			backtrack(s[i:], index+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isValid(str string) bool {
	if len(str) == 0 || len(str) > 3 {
		return false
	}
	if str[0] == '0' && len(str) > 1 {
		return false
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}