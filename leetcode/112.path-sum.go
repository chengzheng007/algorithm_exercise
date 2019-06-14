/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	res := false
	AddNode(&res, root, 0, sum)
	return res
}

func AddNode(res *bool, node *TreeNode, curSum, sum int) {
	if *res || node == nil {
		return
	}
	if node.Val+curSum == sum && node.Left == nil && node.Right == nil {
		*res = true
		return
	}
	AddNode(res, node.Left, node.Val+curSum, sum)
	AddNode(res, node.Right, node.Val+curSum, sum)
}