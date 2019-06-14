/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res, val := true, root.Val
	doValidJudge(&res, root.Left, nil, &val)
	doValidJudge(&res, root.Right, &val, nil)
	return res
}

func doValidJudge(res *bool, node *TreeNode, lower, upper *int) {
	if node == nil || !*res {
		return
	}
	if (lower != nil && node.Val <= *lower) ||
		(upper != nil && node.Val >= *upper) {
		*res = false
		return
	}

	val := node.Val
	doValidJudge(res, node.Left, lower, &val)
	doValidJudge(res, node.Right, &val, upper)
}