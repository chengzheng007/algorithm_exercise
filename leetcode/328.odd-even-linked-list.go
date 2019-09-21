/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddH := new(ListNode)
	p1 := oddH
	evenH := new(ListNode)
	p2 := evenH
	q := head
	i := 1
	for q != nil {
		// 将q摘出来，并断开
		temp := q
		q = q.Next
		temp.Next = nil
		if i%2 == 1 {
			p1.Next = temp
			p1 = temp
		} else {
			p2.Next = temp
			p2 = temp
		}
		i++
	}
	// 将偶数序号的节点接在奇数序号最后一个节点的后边
	p1.Next = evenH.Next
	return oddH.Next
}

