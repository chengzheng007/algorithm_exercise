func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// 环形数组分为两段：从0开始、n-2结束的路线（n-1不抢）
	// 从1开始、n-1结束的路线(0不抢)
	// 取两者最大的作为返回
	from0 := helper(nums, 0, n-2)
	from1 := helper(nums, 1, n-1)

	if from0 > from1 {
		return from0
	}

	return from1
}

func helper(nums []int, lo, hi int) int {
	pre := 0
	cur := 0
	for i := lo; i <= hi; i++ {
		// cur：前一个位置的抢劫最大金额(i-1)
		// pre：cur的上一个位置的最大金额(i-2)
		// temp：当前位置的最大金额
		temp := pre + nums[i]
		if temp < cur {
			temp = cur
		}
		pre = cur
		cur = temp
	}
	return cur
}