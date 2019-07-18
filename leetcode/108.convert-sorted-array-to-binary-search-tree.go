/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}
func dfs(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	node := &TreeNode{Val: nums[mid]}
	// 不需要再递归下去，可以提前终止
	if l == r {
		return node
	}

	node.Left = dfs(nums, l, mid-1)
	node.Right = dfs(nums, mid+1, r)
	return node
}

