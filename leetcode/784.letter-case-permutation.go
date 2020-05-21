func letterCasePermutation(S string) []string {
	var res []string
	for i := 0; i < len(S); i++ {
		letterType := isLetter(S[i])
		size := len(res)
		switch letterType {
		case 1: // 小写字母
			var temp []string
			if size == 0 {
				temp = make([]string, 2)
				temp[0] = string(S[i])
				temp[1] = string(S[i] - 32)
			} else {
				temp = make([]string, 2*size)
				for j := 0; j < size; j++ {
					temp[j] = res[j] + string(S[i])
					temp[j+size] = res[j] + string(S[i]-32)
				}
			}
			res = temp
		case 2: // 大写字母
			var temp []string
			if size == 0 {
				temp = make([]string, 2)
				temp[0] = string(S[i])
				temp[1] = string(S[i] + 32)
			} else {
				temp = make([]string, 2*size)
				for j := 0; j < size; j++ {
					temp[j] = res[j] + string(S[i])
					temp[j+size] = res[j] + string(S[i]+32)
				}
			}
			res = temp
		default: // 非字母
			if size == 0 {
				res = []string{string(S[i])}
			} else {
				for j := range res {
					res[j] = res[j] + string(S[i])
				}
			}
		}
	}
	return res
}

// 0不是字母 1小写字母 2大写字母
func isLetter(s byte) int {
	if s >= 'a' && s <= 'z' {
		return 1
	} else if s >= 'A' && s <= 'Z' {
		return 2
	}
	return 0
}