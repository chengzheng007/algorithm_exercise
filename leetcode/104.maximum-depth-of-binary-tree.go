/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := 1 + maxDepth(root.Left)
	rigthDepth := 1 + maxDepth(root.Right)
	if leftDepth > rigthDepth {
		return leftDepth
	}
	return rigthDepth
}

