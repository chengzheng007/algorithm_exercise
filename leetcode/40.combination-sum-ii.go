import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	var res [][]int
	var path []int
	dfs(candidates, len(candidates), target, 0, path, &res)
	return res
}

func dfs(nums []int, n, t, k int, path []int, res *[][]int) {
	var sum int
	for i := 0; i < len(path); i++ {
		sum += path[i]
	}
	if sum >= t {
		if sum == t {
			temp := make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}
	for i := k; i < n; i++ {
		if i > k && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, n, t, i+1, path, res)
		path = path[:len(path)-1]
	}
}