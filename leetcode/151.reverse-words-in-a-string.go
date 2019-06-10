/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */
func reverseWords(s string) string {
	// reverse whole string
	s = reverseStr(s)
	// reverse each word
	i := 0
	storeIdx := 0
	for i < len(s) {
		if s[i] != ' ' {
			// find and reverse the word
			j := i
			for j < len(s) && s[j] != ' ' {
				j++
			}
			word := reverseStr(s[i:j])

			if storeIdx > 0 {
				s = s[0:storeIdx] + " " + word + s[j:]
				storeIdx += j - i + 1
			} else {
				s = word + s[j:]
				storeIdx += j - i
			}

			i = storeIdx

		} else {
			i++
		}
	}

	// clean space at the end of s
	j := len(s) - 1
	for j > 0 && s[j] == ' ' {
		j--
	}

	// s is full of space string, like "  "
	if j == 0 && s[j] == ' ' {
		s = ""
	} else {
		s = s[0 : j+1]
	}

	return s
}

func reverseStr(s string) string {
	byteData := []byte(s)
	size := len(s)
	for i := 0; i < size/2; i++ {
		byteData[i], byteData[size-1-i] = byteData[size-1-i], byteData[i]
	}
	return string(byteData)
}

