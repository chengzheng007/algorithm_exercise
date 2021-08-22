import "math"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
func maxSubArray(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    } else if n == 1 {
        return nums[0]
    }
    // 注意题设是子数组，数字必须连续，除非从头开始。
    // 什么情况下第i个数加入前面已经生成的子数组有意义呢？
    // 答案是前面子数组的和是大于0，假设前面子数组的和小于0(负数)，加上
    // 当前数字就是没意义的，因为加上前面的子数组，还不如直接
    // 以当前数字作为一个子数组，当前数字作为子数组和会更大
    // 前面子数组和为0也没意义，跟当前数单独作为子数组比起来没起到作用
    // 相当于带了一堆多余的数。总结就是前面的子数组的和+当前数字
    // 比当前数字大才有效
    // 所以当从头到尾遍历这个数组的时候，对于数组里的一个整数，它有
    // 几种选择呢？只有两种：1.加入之前的SubArray
    // 2.自己另起一个SubArray。什么时候加入，什么时候另起
    // 就是上面分析的问题
    // 设dp[i]表示以第i个元素为结尾的子数组最大的和
    // dp[i] = dp[i-1] > 0? dp[i-1]+nums[i] : nums[i],
    // 0<=i<=len(nums)-1
    // 最终求的结果是max(dp[i])，dp数组中最大的
    
    dp := make([]int, n)
    dp[0] = nums[0]
    max := dp[0]
    for i := 1; i < n; i++ {
        if dp[i-1] > 0 {
            dp[i] = dp[i-1]+nums[i]
        } else {
            dp[i] = nums[i]
        }
        if max < dp[i] {
            max = dp[i]
        }
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
