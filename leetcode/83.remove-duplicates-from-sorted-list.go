/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		q := p.Next
		// 寻找下一个不等于当前节点值的节点q
		for q != nil && q.Val == p.Val {
			q = q.Next
		}
		// 如果不等值节点q不是p的下一个节点，删掉p与q之间节点
		if p.Next != q {
			p.Next = q
		}
		// p指向q开始下一轮寻找
		p = q
	}
	return head
}

