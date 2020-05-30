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
	val := root.Val
	return dfs(root.Left, nil, &val) && dfs(root.Right, &val, nil)
}

func dfs(node *TreeNode, lower, upper *int) bool {
	if node == nil {
		return true
	}
	// 右子节点需要比lower大，比他小不合理，lower规定了右分支的下界
	if lower != nil && node.Val <= *lower {
		return false
	}
	// 左子节点需要比upper小，比他大不合理，upper规定了左分支的上界
	if upper != nil && node.Val >= *upper {
		return false
	}

	val := node.Val
	return dfs(node.Left, lower, &val) && dfs(node.Right, &val, upper)
}