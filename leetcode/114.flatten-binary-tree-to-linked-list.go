/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {

	var handler func(node *TreeNode)
	var next *TreeNode
	// 如果函数需要使用公共变量，考虑【闭包】
	handler = func(node *TreeNode) {
		// 尾递归，递归回来时将节点的右指针指向子节点
		if node == nil {
			return
		}

		handler(node.Right)
		handler(node.Left)

		// 递归回来时，next指向已经遍历过的上一个节点，将当前节点(root)
		// 的右指针指向它，将当前节点的做指针指空，就像【尾递归倒置链表】
		// 那样因为次数左指针已经没用需要
		// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/discuss/36977/My-short-post-order-traversal-Java-solution-for-share

		node.Right = next
		node.Left = nil
		next = node
	}

	handler(root)
}