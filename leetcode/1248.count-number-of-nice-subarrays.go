func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k <= 0 {
		return 0
	}
	// mp[sum]表示截止到某下标，总奇数个数为sum的这样的下标总数
	// 令sum[i]表示数组第0个数到第i个数中奇数的个数，因此区间[l, r]符合题意
	// 当且仅当下式成立：sum[r]-sum[l-1]=k，由此我们可以令mp[x]表示有
	// 多少个节点i满足sum[i]=x，然后从左向右枚举，当求得第i个点的sum值后，
	// 更新mp[sum[i]]数组，并计算有多少个l满足区间[l, i]符合题意。
	// 累加答案即可得到最终结果
	mp := make([]int, n+1)
	mp[0] = 1

	sum := 0
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i]%2 == 1 {
			sum++
		}

		mp[sum]++

		if sum >= k {
			// 有mp[num-k]个区间符合条件，累加
			ans += mp[sum-k]
		}
	}

	return ans
}