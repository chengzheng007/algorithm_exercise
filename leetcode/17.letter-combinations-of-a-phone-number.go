func letterCombinations(digits string) []string {
	size := len(digits)
	if size == 0 {
		return nil
	}
	digit2letter := make([][]byte, size)
	for i := 0; i < size; i++ {
		_, letters := getLettersByDigit(digits[i])
		digit2letter[i] = letters
	}
	res := make([]string, 0)
	var path []byte
	backtrack(digit2letter, size, 0, 0, path, &res)
	return res
}

func backtrack(letters [][]byte, n, p, q int, path []byte, res *[]string) {
	if len(path) == n {
		str := string(path)
		*res = append(*res, str)
		return
	}
	for i := p; i < n; i++ {
		for j := q; j < len(letters[i]); j++ {
			path = append(path, letters[i][j])
			// 注意，p传i+1表示递归从下一个数字字符开始
			// q传0表示始终从下一个数字字符对应的字母列表的第一个字母开始组装
			backtrack(letters, n, i+1, 0, path, res)
			path = path[:len(path)-1]
		}
	}
}

func getLettersByDigit(digit byte) (int, []byte) {
	switch digit {
	case '2':
		return 3, []byte{'a', 'b', 'c'}
	case '3':
		return 3, []byte{'d', 'e', 'f'}
	case '4':
		return 3, []byte{'g', 'h', 'i'}
	case '5':
		return 3, []byte{'j', 'k', 'l'}
	case '6':
		return 3, []byte{'m', 'n', 'o'}
	case '7':
		return 4, []byte{'p', 'q', 'r', 's'}
	case '8':
		return 3, []byte{'t', 'u', 'v'}
	case '9':
		return 4, []byte{'w', 'x', 'y', 'z'}
	}
	return 0, nil
}