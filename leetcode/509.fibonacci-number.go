/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */
func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	var (
		n1  = 0
		n2  = 1
		ret int
	)
	for i := 2; i <= N; i++ {
		ret = n1 + n2
		n1 = n2
		n2 = ret
	}
	return ret
}

