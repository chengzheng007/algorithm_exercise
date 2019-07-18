/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	// 由于链表升序排列，由二叉搜索树的特性得知，其中序遍历即为一个升序序列，等于链表这个序列
	// 所以链表的中间节点一定是二叉树的根节点
	// 不必担心不平衡问题，因为节点是逐个处理
	middleNode := getMiddleNode(head)
	// 将中间节点建立为树节点
	treeNode := &TreeNode{Val: middleNode.Val}
	// 如果中间节点就等于head，表示当前链表中只有一个节点
	if head == middleNode {
		return treeNode
	}
	// dfs递归处理树的左右子节点
	treeNode.Left = sortedListToBST(head)
	treeNode.Right = sortedListToBST(middleNode.Next)
	return treeNode
}

func getMiddleNode(node *ListNode) *ListNode {
	var prePslow *ListNode
	pslow := node
	pfast := node
	for pfast != nil && pfast.Next != nil {
		prePslow = pslow
		pslow = pslow.Next
		pfast = pfast.Next.Next
	}
	// 解除链接，将链表分为左右两部分
	// 如不断开，dfs递归会出问题
	if prePslow != nil {
		prePslow.Next = nil
	}
	return pslow
}

