package main

import "fmt"

// 递归方式计算最长增长子序列，思想：从已有的数字中，选中或者不选中某个数，作为备选的LIS
// 如果比前一个数大，那么选中并且得到+1后的长度值，递归使用当前数，位置进行到选一个位置，
// 如果不选中，那么长度不加1，递归继续下一个位置，重复上述操作
// prev：前一个元素值，currPos：当前做出选择的下标
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

// longestIncreasingSequence
func lisDP(digit []int) int {
	size := len(digit)
	if size <= 1 {
		return size
	}

	dp := make([]int, size)

	// 初始化
	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if digit[j] < digit[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	fmt.Printf("dp:%v\n", dp)
	max := 1
	for i := 0; i < size; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
