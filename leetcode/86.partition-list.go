/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	p := head
	newh := new(ListNode)
	newh.Next = p
	q := newh
	for p != nil {
		// 找到大于等于x的开始节点p
		for p != nil && p.Val < x {
			// q为>=x节点中第一个节点的前驱
			q.Next = p
			q = p
			p = p.Next
		}

		// 寻找从p开始的大于等于x的连续节点
		r := p
		// 如果r为空，则当前列表说明不存在值大于x的节点
		for r != nil && r.Next != nil && r.Next.Val >= x {
			r = r.Next
		} // end: 大于等于x的节点区间为[p, r]

		var nextr *ListNode
		if r != nil {
			nextr = r.Next
		}

		// 将[p,r]间的节点移动到nextr节点后面
		// nextr不为空，说明还有后续节点
		if nextr != nil {
			r.Next = nextr.Next
			nextr.Next = p
			q.Next = nextr
		} else {
			// nextr为空说明已到链表末尾，且末尾一段都是>=x的节点，直接将q连接上
			q.Next = p
		}

		p = nextr
	}
	return newh.Next
}

