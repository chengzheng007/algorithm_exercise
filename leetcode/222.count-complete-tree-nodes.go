/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	//     简单方法：递归所有节点统计
	//     if root == nil {
	//         return 0
	//     }

	//     return 1+countNodes(root.Left)+countNodes(root.Right)
	// 因为是完全二叉树，左子树深度>=右子树深度
	// 先沿左子树查到一直找到左子节点为空(leftDepth逻辑)
	// 如果左右子树深度相等，能找到左子树+当前节点的总数
	// 即1<<dl，题目描述的例子里面就是节点1+节点2(包括其
	// 子节点)的总数，然后加上countNodes(root.Right)
	// 否则反之

	if root == nil {
		return 0
	}
	dl := leftDepth(root.Left)
	dr := leftDepth(root.Right)
	if dl == dr {
		return 1<<dl + countNodes(root.Right)
	} else {
		return 1<<dr + countNodes(root.Left)
	}
}

func leftDepth(node *TreeNode) int {
	var d int
	for ; node != nil; d, node = d+1, node.Left {
	}
	return d
}