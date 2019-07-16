/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 利用p1，p2两个节点
	p1 := head
	p2 := head.Next
	newh := new(ListNode)
	q := newh
	for p2 != nil {
		p1.Next = p2.Next
		p2.Next = p1
		q.Next = p2
		// 需要记录当前这一段交换后的末尾的节点p1，换到第二次时，将前后两段衔接起来
		q = p1

		p1 = p1.Next
		if p1 != nil {
			p2 = p1.Next
		} else {
			p2 = nil
		}
	}
	return newh.Next
}

