package main

import "fmt"

// 最长递增子序列
// longest increasing sequence

func lisDFS(digit []int) int {
	if len(digit) == 0 {
		return 0
	}
	max := 0
	lisBacktrack(digit, 1, 1, &max)
	return max
}

func lisBacktrack(digit []int, k, currLIS int, max *int) {
	if currLIS > *max {
		*max = currLIS
	}
	if k == len(digit) {
		return
	}
	for i := k; i < len(digit); i++ {
		fmt.Printf("i:%d, k:%d, currLIS:%d, max:%d\n", i, k, currLIS, *max)
		if i > 0 && digit[i] >= digit[i-1] {
			lisBacktrack(digit, i+1, currLIS+1, max)
		}
	}
}

// 递归方式计算最长增长子序列，思想：从已有的数字中，选中或者不选中某个数，作为备选的LIS
// 如果比前一个数大，那么选中并且得到+1后的长度值，递归使用当前数作为下一次递归的前一个数用
// 如果不选中，那么长度不加1，for循环本层次的下一个位置，重复上述操作
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
	dp[0] = 1
	// for i := 0; i < size; i++ {
	// 	dp[i] = 1
	// }
	max := 0
	for i := 1; i < size; i++ {
		maxLocal := 0
		for j := 0; j < i; j++ {
			if digit[j] < digit[i] && maxLocal < dp[j] {
				// dp[i] = dp[j] + 1
				maxLocal = dp[j]
			}
		}
		dp[i] = maxLocal + 1
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Printf("dp:%v\n", dp)
	// max := 1
	// for i := 0; i < size; i++ {
	// 	if max < dp[i] {
	// 		max = dp[i]
	// 	}
	// }
	return max
}

// 最长自增序列长度另一种解法：https://blog.csdn.net/wqtltm/article/details/81253935
// 时间复杂度：O(n*logn)
func lis2(list []int) int {
	n := len(list)
	if n <= 1 {
		return n
	}
	dp := []int{list[0]}
	for i := 1; i < n; i++ {
		if list[i] > dp[len(dp)-1] {
			dp = append(dp, list[i])
		} else {
			left := 0
			right := len(dp) - 1
			for left < right {
				mid := (left + right) / 2
				if dp[mid] == list[i] {
					break
				} else if list[i] > dp[mid] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			dp[right] = list[i]
		}
	}

	fmt.Printf("lis2 dp:%v\n", dp)
	return len(dp)
}
