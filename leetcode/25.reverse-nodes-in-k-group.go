/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// p指向k个分组节点的最后一个节点
	p := head
	for i := 1; i < k && p != nil; i++ {
		p = p.Next
	}

	if p == nil { // 不足k个节点
		return head
	}
	q := p.Next
	p.Next = nil // 将当前k个节点分组从主链断开

	// 问题分解为剩下的节点的进行k个分组倒置的子问题
	newq := reverseKGroup(q, k)
	// 倒置当前的分组
	newp := reverseList(head)
	// 以上两个递归函数的先后调用顺序没有要求

	head.Next = newq
	return newp
}

// 尾递归翻转不带头结点的链表
func reverseList(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	// 先递归到尾部
	newh := reverseList(h.Next)
	// 返回时修改节点信息
	h.Next.Next = h
	h.Next = nil
	return newh
}

