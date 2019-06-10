/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
func reverseString(s []byte) {
	size := len(s)
	for i := 0; i < (size >> 1); i++ {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}
}

