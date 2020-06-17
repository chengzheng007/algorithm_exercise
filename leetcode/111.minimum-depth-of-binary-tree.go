import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	// 树为空
	if root == nil {
		return 0
	}

	minDep := math.MaxInt64
	findMinDep(root, 0, &minDep)
	return minDep
}

func findMinDep(node *TreeNode, curr int, min *int) {
	if node == nil {
		return
	}
	curr++

	// 叶子节点时检测深度
	if node.Left == nil && node.Right == nil {
		if curr < *min {
			*min = curr
		}
		return
	}
	findMinDep(node.Left, curr, min)
	findMinDep(node.Right, curr, min)
	// 递归返回时撤销之前加上的深度
	curr--
}