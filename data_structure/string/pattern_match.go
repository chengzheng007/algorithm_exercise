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
			i = i - j + 1
			j = 0
		}
	}

	if j >= sizeOfP {
		return i - sizeOfP
	}
	
	return -1
}

// kmp，求模式串的最长前后缀子串
// 下标表示前缀子串的结尾下标，值为最长可匹配前缀的长度
func getNext(p string, size int) []int {
	next := make([]int, size)
	next[0] = 0
	var i, k int
	for i = 1; i < size; i++ {
		// p[0...k-1] p[i-k...i-1]已经匹配（p[0...k-1]是p[0...i-1]的最长前缀子串）
		// 但p[k] != p[i]，因此要在查找p[0, i-1]中次长的前缀子串
		// 这个次长前(后)缀子串，字符肯定也是以p[0]开头，以p[i-1]结尾，单个字符来说
		// 假设次长的前缀子串是p[0...r]（如果存在的话），对应着后缀子串是p[i-r-2...i-1]
		// 接下来就又可以比较p[r+1]与p[i]
		// 如何去确定这个次长前(后)缀子串？注意上面是p[k]与p[i]不匹配，即使用的是p[0...i-1]的最长前(后)缀子串去匹配
		// 匹配的前缀是p[0...k-1]
		// 如果我们要找另外一个子串，就需要第一个字符是p[0]，最后一个字符是p[k-1]
		// 实际就是在找p[0...k-1]的最长前(后)缀子串（有点回到了这个算法本身处理模式的感觉，从0...i-1的最长前缀子串回到这里）
		// 只有它的最长可匹配前缀子串才符合「第一个字符是p[0]，最后一个字符是p[k-1]」这个要求
		// p[0...k-1]的最长前(后)缀子串的长度就是next[k-1]，也就是这里 k = next[i-1] 的意思
		// 如果存在就继续匹配，否则继续往前找可匹配子串
		for k > 0 && p[i] != p[k] {
			k = next[k-1]
		}
		
		if p[i] == p[k] {
			k++
		}
		next[i] = k
	}
	return next
}

func kmp(mains, pattern string) int {
	sizeOfM := len(mains)
	sizeOfP := len(pattern)
	next := getNext(pattern, sizeOfP)
	var i, j int
	for ; i < sizeOfM; i++ {
		for mains[i] != pattern[j] && j > 0 {
			j = next[j-1]
		}
		if mains[i] == pattern[j] {
			j++
		}
		if j == sizeOfP {
			return i-j+1
		}
	}
	return  -1
}

