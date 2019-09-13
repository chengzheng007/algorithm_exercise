/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 先让一个节点q指向第n个，然后将p、q顺序后移，到q是最后一个节点时
// p指向倒数第n个节点的前驱，特殊情况：链表本身不足n个节点，删除的是第一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	// p 作为待删除节点的前驱，q指向与p距离n-1个节点的节点
	p := head
	q := head
	i := 0
	for i < n {
		if q != nil {
			q = q.Next
		}
		i++
		// 此处包含链表长度为1，或者删除节点为第一个节点时
		if q == nil {
			break
		}
	}
	// 倒数第n个节点不存在，因链表本身不够长
	if i != n {
		return nil
	}
	// 循环链表找到倒数第n个节点前驱
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next
	}
	// 删除节点为第一个节点
	if p == head && q == nil {
		head = head.Next
		return head
	}
	// 删除节点不是第一个节点
	temp := p.Next
	p.Next = temp.Next
	temp = nil
	return head
}

