import "math"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	// dp[i]表示到i下标的最大连续数的和
	// dp := make([]int, size)
	// base case
	// 可以不用数组记录备忘录的值，用临时变量
	// dp[0] = nums[0]
	// max := dp[0]

	prev := nums[0]
	max := nums[0]
	for i := 1; i < size; i++ {
		var temp int
		// 如果dp[i-1]加上当nums[i]，使得比nums[i]还小，说明[i-1]是一个比较小的负数，加上dp[i-1]整体只会是减小的趋势，这种情况直接取nums[i]作为dp数组的结果
		// 如果dp[i-1]+nums[i]比nums[i]大，说明在朝着增大方向发展，加上nums[i]会越来越大
		// 这么做的另一个原因是题目要求连续数据
		if prev+nums[i] <= nums[i] {
			temp = nums[i]
		} else {
			temp = prev + nums[i]
		}
		if max < temp {
			max = temp
		}
		prev = temp

		// dp版本
		// if dp[i-1]+nums[i] <= nums[i] {
		//     dp[i] = nums[i]
		// } else {
		//     dp[i] = dp[i-1]+nums[i]
		// }
		// if max < dp[i] {
		//     max = dp[i]
		// }
	}
	return max
}

// 另一种解法
// https://en.wikipedia.org/wiki/Maximum_subarray_problem
func traverse(nums []int) int {
	smax := math.MinInt64
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > smax {
			smax = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return smax
}