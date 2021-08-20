/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

func lengthOfLIS(nums []int) int {
    size := len(nums)
    if size <= 1 {
        return size
    }
    
    dp := make([]int, size)    
    for i := 0; i < size; i++ {
        dp[i] = 1
    }
    
    max := dp[0]
    for i := 1; i < size; i++ {
        for j := 0; j < i; j++ {
	    // nums[j]<nums[j]，意味着下标为j的数可能会贡献递增的序列
	    // 但是这里要扫描完所有比nums[i]小的数取最大的贡献值
	    // 而不是取第一个比nums[i]小的子问题解+1就break
	    // 因为可能i的前一个数就已经是当前最后一个递增数，也可能前面几个数才是针对当前数一直递增的
	    // 比如3, 1, 6, 4, 10，10比4大，但是最长的是1,6,10或1,4,10或3,6,10	
            if nums[i] > nums[j] && dp[i] < dp[j]+1 {
                dp[i] = dp[j]+1
            }
        }
        if dp[i] > max {
            max = dp[i]
        }
    }

    return max
}

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}

	dp := make([]int, size)

	// 初始化，每个位置自增都是1（只看自己）
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			// nums[j]<nums[j]，意味着下标为j的数可能会贡献递增的序列
			// 但是这里要扫描完所有比nums[i]小的数取最大的贡献值
			// 而不是取第一个比nums[i]小的子问题解+1就break
			// 因为可能i的前一个数就已经是当前最后一个递增数，也可能前面几个数才是针对当前数一直递增的
			// 比如3, 1, 6, 4, 10，10比4大，但是最长的是1,6,10或1,4,10或3,6,10
			if nums[j] < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := 1
	for i := 0; i < size; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// 另一种写法
func lengthOfLIS2(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}

	dp := make([]int, size)
	max := 0
	for i := 1; i < size; i++ {
		localMax := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && localMax < dp[j] {
				localMax = dp[j]
			}
		}
		dp[i] = localMax + 1
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// 递归方式计算最长增长子序列，思想：从已有的数字中，选中或者不选中某个数，作为备选的LIS
// 如果比前一个数大，那么选中并且得到+1后的长度值，递归使用当前数，位置进行到选一个位置，
// 如果不选中，那么长度不加1，递归继续下一个位置，重复上述操作
// prev：前一个元素值，currPos：当前做出选择的下标
// 调用：lengthOfLISDFS(nums, math.MinInt64, 0)
func lengthOfLISDFS(nums []int, prev, currPos int) int {
	if currPos == len(nums) {
		return 0
	}
	take := 0
	// 当前位置数字与前一数字比较，如果比它大则选中，prev变为当前位置数据，当前位置前进一位，重复
	if nums[currPos] > prev {
		take = 1 + lengthOfLISDFS(nums, nums[currPos], currPos+1)
	}
	// 不选中，依然使用prev作为比较数字，位置前进1位
	nottake := lengthOfLISDFS(nums, prev, currPos+1)
	if take > nottake {
		return take
	}
	return nottake
}
