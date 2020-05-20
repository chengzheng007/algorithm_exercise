import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	var res [][]int
	var path []int
	// used表示下标为i的数是否已被使用过
	used := make([]bool, len(nums))
	res = dfs(nums, len(nums), used, path, res)
	return res
}

func dfs(nums []int, n int, used []bool, path []int, res [][]int) [][]int {
	if len(path) == n {
		temp := make([]int, n)
		copy(temp, path)
		res = append(res, temp)
		return res
	}

	for i := 0; i < n; i++ {
		// 已使用过的剪枝
		if used[i] {
			continue
		}
		// 重复元素的剪枝
		// 元素从前一个兄弟节点回退，到当前节点的下推时，前一个相同（假定相同）的元素因回退被标记为未使用状态
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}

		// 做出选择
		path = append(path, nums[i])
		used[i] = true

		res = dfs(nums, n, used, path, res)

		// 递归回退时撤回选择（恢复）
		path = path[:len(path)-1]
		used[i] = false
	}
	return res
}