/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 前序遍历的第一个节点肯定是整棵树根节点
	// 在中序遍历的数组中，这个节点的左边的数就是左子树的节点
	// 这个节点右边的数就是右子树的节点
	// 依据这样的特性，每次确定根节点，递归处理左右子树
	// 加个左右子树及下一个根节点计算好，就可以确定整棵树
	return genTree(preorder, inorder, 0, 0, len(inorder)-1)
}

// rootIdxAtPreArr：根节点在前序数组中的下标
// lIdxAtIn：在中序数组中左边界，用于定位目前是哪一段正在计算当前子树的根节点
// rIdxAtIn：在中序数组中右边界
func genTree(preorder, inorder []int, rootIdxAtPreArr, lIdxAtIn, rIdxAtIn int) *TreeNode {
	if rootIdxAtPreArr >= len(preorder) || lIdxAtIn > rIdxAtIn {
		return nil
	}

	// 在中序数组中找根节点的位置
	var rootIdxAtInArr int
	for i := lIdxAtIn; i <= rIdxAtIn; i++ {
		if inorder[i] == preorder[rootIdxAtPreArr] {
			rootIdxAtInArr = i
			break
		}
	}

	node := &TreeNode{Val: inorder[rootIdxAtInArr]}
	// 左右边界相等，只有这一个节点，直接返回无需递归
	if lIdxAtIn == rIdxAtIn {
		return node
	}

	// 计算左子树的根节点下标
	node.Left = genTree(preorder, inorder, rootIdxAtPreArr+1, lIdxAtIn, rootIdxAtInArr-1)
	// 计算右子树的根节点下标
	// 注意右半部分下标计算
	node.Right = genTree(preorder, inorder, rootIdxAtPreArr+rootIdxAtInArr-lIdxAtIn+1, rootIdxAtInArr+1, rIdxAtIn)
	return node
}

