/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// 不需要从根节点开始，意味着pathSum本身要从根节点
	// 递归函数套递归
	return countPathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func countPathSum(node *TreeNode, sum int) int {
	// sum本身可能为负数，节点值是负的
	if node == nil {
		return 0
	}

	cnt := 0
	if node.Val == sum {
		cnt = 1
		// 后面的节点也许和为0，所以仍需递归下去
	}
	// 从node节点递归
	return cnt + countPathSum(node.Left, sum-node.Val) + countPathSum(node.Right, sum-node.Val)
}