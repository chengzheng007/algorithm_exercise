/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	var path []int
	genPath(root, sum, path, &res)
	return res
}

func genPath(node *TreeNode, sum int, path []int, res *[][]int) {
	if node == nil {
		return
	}

	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil && sum == node.Val {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	genPath(node.Left, sum-node.Val, path, res)
	genPath(node.Right, sum-node.Val, path, res)

	path = path[:len(path)-1]
}