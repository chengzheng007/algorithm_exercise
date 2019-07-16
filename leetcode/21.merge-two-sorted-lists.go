/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	head := new(ListNode)
	r := head
	for p != nil && q != nil {
		if p.Val < q.Val {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}
	if p != nil {
		r.Next = p
	}
	if q != nil {
		r.Next = q
	}
	return head.Next
}

