func isPalindrome(x int) bool {
	// 思路1:转换为字符串，判断字符串，需消耗额外O(n)空间
	// 思路2:如果是回文数，数字的前一半与后一半相等，如1221
	// 问题：如果判断是获取到后一半数字

	// 负数都不是回文数末尾是0的且不为0的数也不是
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	lastHalf := 0
	for x > lastHalf {
		lastHalf = lastHalf*10 + x%10
		x /= 10
	}
	// 偶数位(1221) || 基数位(121)
	return lastHalf == x || x == lastHalf/10

}