/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
func myAtoi(str string) int {
	// clean space in the front
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	str = str[i:]

	// flag indicate it is a positive number or negitive
	flag := 1
	if len(str) != 0 {
		// negitive number
		if str[0] == '-' {
			flag = -1
			str = str[1:]
		} else if str[0] == '+' {
			str = str[1:]
		} else if str[0] < '0' || str[0] > '9' {
			// not a number
			return 0
		}
	} else {
		return 0
	}

	// get a numerical world
	var word []byte
	for j := 0; j < len(str); j++ {
		if str[j] < '0' || str[j] > '9' {
			if j == 0 {
				word = []byte{}
			}
			break
		}
		word = append(word, str[j])
	}

	// interpret word to a digit number

	var (
		maxInt = 1<<31 - 1
		minInt = -1 << 31
		num    = 0
		pos    = 1
	)

	for j := len(word) - 1; j >= 0; j-- {
		curData := int(word[j]-'0') * pos * flag
		if flag > 0 {
			// overflow check
			if curData >= maxInt-num {
				num = maxInt
				break
			} else {
				num += curData
			}
		} else {
			// overflow check
			if curData <= minInt-num {
				num = minInt
				break
			} else {
				num += curData
			}
		}
		// pos may be greater than maxInt ...
		if maxInt/pos < 10 {
			pos = maxInt
		} else {
			pos *= 10
		}
	}

	return num
}

