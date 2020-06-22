/**
	Given a string S, you are allowed to convert it to a palindrome by adding
	characters in front of it. Find and return the shortest palindrome you
	can find by performing this transformation.
For example:
Given “aacecaaa”, return “aaacecaaa”.
Given “abcd”, return “dcbabcd”.
*/
func convertPalindrome(S string) string {
	// 思路：从第一个字母开始，寻找最长能形成的回文子串，然后将剩余的字母copy，倒序，放在S前面
	// 比如abacd，从开头第一个字符开始最长回文串是aba，剩下的是cd，那么把dc加到整个串前面就是最短的

	n := len(S)
	if n < 2 {
		return S
	}

	end := n
	for i := n; i > 0; i-- {
		if isPalindrome(S[0:i]) {
			end = i
			break
		}
	}
	// 原始串已经是回文串
	if end == n {
		return S
	}

	// end之后的子串进行reverse操作
	restr := ""
	for i := n - 1; i >= end; i-- {
		restr += string(S[i])
	}

	return restr + S
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}