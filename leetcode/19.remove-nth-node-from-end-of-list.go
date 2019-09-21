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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 两个指针，一个在第一个节点，一个指向后面与之相距n-1个节点
	// 然后两个指针同时移到到末尾，前一个指针指向的是倒数第n个节点
	// 的前驱，如果存在的话
	if head == nil {
		return nil
	}
	p := head
	q := head
	i := 0
	// p指向第一个节点，通过计数方式将q指向与之相距n-1的节点
	for i < n {
		if q != nil {
			q = q.Next
		}
		i++
		if q == nil {
			break
		}
	}
	// 倒数第n个不存在，链表不够长
	if i < n {
		return head
	}
	// p、q依次遍历到末尾
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next
	} // end:p指向倒数第n个节点的前驱，q指向倒数第1个节点

	// 删除的是第1个节点
	// 特例：只有1个节点并删除它
	if q == nil && p == head {
		head = head.Next
		return head
	}
	p.Next = p.Next.Next
	return head
}

