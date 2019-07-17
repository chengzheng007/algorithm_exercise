/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	p := head
	newh := new(ListNode)
	q := newh
	// 需要一个头节点，方便操作
	// 注意删除第一个元素的情况，这时带头结点更方便
	for p != nil {
		if p.Val == val {
			q.Next = p.Next
		} else {
			q.Next = p
			q = p
		}
		p = p.Next
	}
	return newh.Next
}

