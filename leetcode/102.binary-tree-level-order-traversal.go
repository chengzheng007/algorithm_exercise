/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	levelMaker(root, &res, 0)
	return res
}

// 参考leetcode107解法
func levelMaker(node *TreeNode, res *[][]int, depth int) {
	if node == nil {
		return
	}
	if len(*res) <= depth {
		*res = append(*res, make([]int, 0))
	}
	levelMaker(node.Left, res, depth+1)
	levelMaker(node.Right, res, depth+1)
	(*res)[depth] = append((*res)[depth], node.Val)
}