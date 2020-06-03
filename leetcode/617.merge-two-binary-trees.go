/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 思路：沿着一棵树遍历（如t1），如果其子节点为空，直接返回
	// 另一棵树的节点，如果两个节点都不为空，则将t2的值加到t1上
	// 在递归遍历t1、t2的左子节点，作为当前t1节点的左子树
	// 右子树同理
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}