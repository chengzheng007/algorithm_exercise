/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(lnode, rnode *TreeNode) bool {
	if lnode == nil && rnode == nil {
		return true
	}
	if lnode == nil || rnode == nil {
		return false
	}
	if lnode.Val != rnode.Val {
		return false
	}
	// 注意是镜像，画图看，左子树的左子节点需要与右子树的右子节点比较
	// 左子树的右子节点需要与右子树的左子节点比较
	return check(lnode.Left, rnode.Right) && check(lnode.Right, rnode.Left)
}