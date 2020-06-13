/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	helper(root, 0, &res)
	return res
}

func helper(node *TreeNode, level int, res *[]int) {
	if node == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, node.Val)
	}
	// level表示树层数，先递归右子树，从右子树下来的res长度
	// 恰好等于level，如果再次进入其左子树时，res长度已经大于level
	// 不会被加到res中
	helper(node.Right, level+1, res)
	helper(node.Left, level+1, res)
}