/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	newh := new(ListNode)
	q := newh
	for p != nil {
		r := p.Next
		// 寻找下一个与p不同值的节点r
		for r != nil && r.Val == p.Val {
			r = r.Next
		}
		// 如果r不是p的下一个节点，说明p之后r之前，均为重复值节点，直接删除
		// q.Next指向r，表示删除p到r之前的重复节点
		if r != p.Next {
			q.Next = r
		} else {
			// r与p是不同的节点，p节点值不重复，将p衔接进新链表
			q.Next = p
			q = p
		}
		p = r
	}
	return newh.Next
}

