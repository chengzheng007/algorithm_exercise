/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var path []int
	backtrack(candidates, len(candidates), target, 0, path, &res)
	return res
}

func backtrack(nums []int, n, target, k int, path []int, res *[][]int) {
	sum := 0
	for i := 0; i < len(path); i++ {
		sum += path[i]
	}
	// 达到递归结束条件
	if sum >= target {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}
	for i := k; i < n; i++ {
		// 选择
		path := append(path, nums[i])
		backtrack(nums, n, target, i, path, res)
		// 撤销恢复
		path = path[:len(path)-1]
	}
}

