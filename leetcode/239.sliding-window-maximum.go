/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow(nums []int, k int) []int {
	// q:双端队列，存储数组元素的下标
	// 维持q内下标的元素单调递减，每次取队头比较
	q := make([]int, 0)
	res := make([]int, 0)
	for idx, num := range nums {
		// 如果当前位置比队列头部元素多出k个元素，删除多余元素
		if len(q) > 0 && idx-q[0] >= k {
			q = q[1:]
		}

		// 删除尾部比当前元素小的
		for len(q) > 0 {
			l := len(q)
			if nums[q[l-1]] < num {
				q = q[:l-1]
			} else {
				break
			}
		}

		q = append(q, idx)

		if idx >= k-1 {
			// 追加最大的数字到结果集
			res = append(res, nums[q[0]])
		}
	}

	return res
}