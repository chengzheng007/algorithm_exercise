func numTrees(n int) int {
	if n < 0 {
		return 0
	} else if n <= 1 {
		return 1
	}
	// 卡特兰数
	// https://leetcode.com/problems/unique-binary-search-trees/discuss/31666/DP-Solution-in-6-lines-with-explanation.-F(i-n)-G(i-1)-*-G(n-i)
	// G(n) = G(0) * G(n-1) + G(1) * G(n-2) + … + G(n-1) * G(0)
	G := make([]int, n+1)
	// base case
	G[0] = 1
	G[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}

	return G[n]
}