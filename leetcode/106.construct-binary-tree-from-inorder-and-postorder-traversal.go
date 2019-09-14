/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	return genTree(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
}

func genTree(inorder, postorder []int, rootIdxAtPost, lIdxAtIn, rIdxAtIn int) *TreeNode {
	if rootIdxAtPost < 0 || lIdxAtIn > rIdxAtIn {
		return nil
	}

	node := &TreeNode{Val: postorder[rootIdxAtPost]}
	// 该段只有一个节点，无需向下递归
	if lIdxAtIn == rIdxAtIn {
		return node
	}

	// 寻找在中序数组中根节点的位置
	var rootIdxAtIn int
	for i := lIdxAtIn; i <= rIdxAtIn; i++ {
		if postorder[rootIdxAtPost] == inorder[i] {
			rootIdxAtIn = i
			break
		}
	}

	// rootIdxAtIn左边为左子树，右边为右子树
	// 重点是索引边界如何计算，当前范围是
	node.Right = genTree(inorder, postorder, rootIdxAtPost-1, rootIdxAtIn+1, rIdxAtIn)
	node.Left = genTree(inorder, postorder, rootIdxAtPost-rIdxAtIn+rootIdxAtIn-1, lIdxAtIn, rootIdxAtIn-1)
	return node
}

