/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// h1存储小于x的节点，h2存储大于等于x的节点
	h1 := new(ListNode)
	h2 := new(ListNode)
	p := h1
	q := h2
	r := head
	for r != nil {
		if r.Val < x {
			p.Next = r
			p = p.Next
		} else {
			q.Next = r
			q = q.Next
		}
		r = r.Next
	}

	q.Next = nil     // 将大于x的链表末尾截断
	p.Next = h2.Next // p1接上p2
	return h1.Next
}