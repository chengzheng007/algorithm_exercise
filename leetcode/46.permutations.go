/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	backtrack(nums, n, 0, &result)
	return result
}

func backtrack(nums []int, n, start int, result *[][]int) {
	if start == n {
		// 这里必须临时复制一份，因为各个函数会对nums造成影响
		// 递归返回后，将数据换回来，导致result中
		// 的排列都是换回来的结果
		temp := make([]int, n)
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	// 外层：i从start循环到n，将这之间的数与下标为start的数交换
	// 每换一次递归处理本次换完后的情况
	for i := start; i < n; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		backtrack(nums, n, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// 另一种解法，使用辅助的变量，记录数字是否被访问过，找出当前未访问的放在路径变量中
func permute2(nums []int) [][]int {
	var res [][]int
	var path []int
	visited := make([]bool, len(nums))
	res = backtrack(nums, visited, path, res)
	return res
}

// 回溯算法一般都是沿着一棵树遍历（决策树），从根节点往下是一条条路径，每个节点提供一个选择列表
// 没经过一个节点，就填充路径变量，回退到上一个节点时，撤销对这个路径的填充，也称为选择和撤销选择
// 遍历到叶子节点，相当于到了结束条件，这时将路径变量收集到结果中
// 回溯继续下一次遍历
func backtrack2(nums []int, visited []bool, path []int, res [][]int) [][]int {
	if len(path) == len(nums) {
		copyOfNums := make([]int, len(nums))
		copy(copyOfNums, path)
		res = append(res, copyOfNums)
		return res
	}
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			// 选择
			visited[i] = true
			path = append(path, nums[i])
			res = backtrack(nums, visited, path, res)
			// 撤销选择
			visited[i] = false
			path = path[:len(path)-1] // 底层数组指针复用
		}
	}
	return res
}