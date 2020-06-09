func findLength(A []int, B []int) int {
	sizeA := len(A)
	sizeB := len(B)
	if sizeA == 0 || sizeB == 0 {
		return 0
	}
	// dp[i][j]表示在A[0...i)和B[0...j)的重复子数组长度
	dp := make([][]int, sizeA+1)
	for i := 0; i <= sizeA; i++ {
		dp[i] = make([]int, sizeB+1)
	}
	maxLen := 0
	for i := 1; i <= sizeA; i++ {
		for j := 1; j <= sizeB; j++ {
			if A[i-1] == B[j-1] {
				if i > 1 && j > 1 && A[i-2] == B[j-2] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				dp[i-1][j-1] = maxInt(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}

	return maxLen
}

func maxInt(x, y, z int) int {
	max := 0
	if x > max {
		max = x
	}
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}