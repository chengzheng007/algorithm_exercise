/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	lc := root.Left
	rc := root.Right
	root.Left = rc
	root.Right = lc
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

