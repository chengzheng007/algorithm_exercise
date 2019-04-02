package main

// pattern 模式串
// mains 主串

// brute force朴素匹配
func bf(mains, pattern string) int {
	sizeOfM := len(mains)
	sizeOfP := len(pattern)

	var i, j int

	for i < sizeOfM && j < sizeOfP {
		if pattern[j] == mains[i] {
			i++
			j++
		} else {
			j = 0
			i = i - j + 1
		}
	}

	if j >= sizeOfP {
		return i - sizeOfP
	}
	
	return -1
}
