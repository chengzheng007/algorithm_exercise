func search(nums []int, target int) bool {
	return dfs(nums, 0, len(nums)-1, target)
}

// 直接递归搜索两边，而不是一边
// 因为两边各自为递增序列，并且两边都可能有数字
func dfs(nums []int, l, r, target int) bool {
	if l > r {
		return false
	}

	mid := (l + r) / 2
	if nums[mid] == target {
		return true
	}

	return dfs(nums, mid+1, r, target) || dfs(nums, l, mid-1, target)
}