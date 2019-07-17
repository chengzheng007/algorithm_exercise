/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	newh := new(ListNode)
	newh.Next = head

	// 寻找m-n之间的头尾节点
	// preRev为待翻转节点区间的前驱，tailRev为待翻转的最后一个节点
	var preRev, tailRev *ListNode
	p := newh
	i := 1
	for i <= n {
		if i <= m {
			preRev = p
		}
		// 链表长度不够
		if p == nil {
			break
		}
		p = p.Next
		if i == n {
			tailRev = p
		}
		i++
	} // end:待翻转节点区间为(preDel、tailDel]

	var afterRev *ListNode
	if tailRev != nil {
		afterRev = tailRev.Next
	}

	p = preRev.Next
	var q, r *ListNode
	if p != nil {
		q = p.Next
	}
	// 翻转(preDel、tailDel]之间的节点
	// 注意for循环的终止条件
	for q != nil && q != afterRev {
		r = q.Next
		q.Next = p
		p = q
		q = r
	}
	// 将翻转后区间重新链到链表中
	preRev.Next.Next = afterRev
	preRev.Next = tailRev
	return newh.Next
}

