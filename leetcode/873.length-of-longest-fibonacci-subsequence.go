/*
 * @lc app=leetcode id=873 lang=golang
 *
 * [873] Length of Longest Fibonacci Subsequence
 */
func lenLongestFibSubseq(A []int) int {
	size := len(A)
	// 将所有的值和索引存储于一个map，额外的空间O(n)
	mp := make(map[int]int, size)
	for i, v := range A {
		mp[v] = i
	}
	// 动态规划思想, i,j,k均为数组中的索引，假设A[i]+A[j]=A[k](A[k]存在)
	// 可以看成是从的[i, j]转移至[j, k]
	// 假设longest[i, j]表示数组中到i,j出下边最长的类斐波拉契子序列长度
	// 因为longest是一个map，要同时记录i,j和长度值三个变量
	// 在实现上，这里将i*size+j整体作为map的key，他的长度值作为val
	// 时间复杂度：O(n^2)
	longest := make(map[int]int, 0)
	ans := 0
	for k := 0; k < size; k++ {
		for j := 0; j < k; j++ {
			if A[k]-A[j] >= A[j] {
				continue
			}
			i, ok := mp[A[k]-A[j]]
			if !ok {
				// 两数之差不存在
				continue
			}
			longest[j*size+k] = longest[i*size+j] + 1
			// +2是因为，斐波拉契数列至少要3个数
			if longest[j*size+k]+2 > ans {
				ans = longest[j*size+k] + 2
			}
		}
	}
	if ans >= 3 {
		return ans
	}
	return 0
}

