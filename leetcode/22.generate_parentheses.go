func generateParenthesis(n int) []string {
	// 回溯暴力枚举
	if n <= 0 {
		return nil
	}
	res := make([]string, 0)
	path := make([]byte, 2*n)
	genAll(2*n, 0, path, &res)
	// genAll2(n, 0, 0, "", &res)
	return res
}

func genAll(pathLen, pos int, path []byte, res *[]string) {
	if pos == pathLen {
		if isValid(path) {
			temp := string(path)
			*res = append(*res, temp)
		}
		return
	}
	// 先尝试左括号
	path[pos] = '('
	genAll(pathLen, pos+1, path, res)
	// 再尝试右括号
	path[pos] = ')'
	genAll(pathLen, pos+1, path, res)
	// 因为每个位置都会尝试左右括号，所以时间复杂度为2的2n次方乘n
}

func isValid(str string) bool {
	balance := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

// 追踪左右括号是否平衡，当发现左括号未超过n时，尝试加左括号
// 当发现右括号比左括号少时，尝试加右括号
func genAll2(n, opened, closed int, str string, res *[]string) {
	if len(str) == 2*n {
		*res = append(*res, str)
		return
	}
	// 如果左括号未超过一半，尝试加左括号
	if opened < n {
		genAll2(n, opened+1, closed, str+"(", res)
	}
	if closed < opened {
		genAll2(n, opened, closed+1, str+")", res)
	}
}