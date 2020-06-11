/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h1 := new(ListNode)
	h2 := new(ListNode)
	pOdd := h1
	pEven := h2
	q := head
	num := 1
	for q != nil {
		if num%2 == 1 {
			pOdd.Next = q
			pOdd = pOdd.Next
		} else {
			pEven.Next = q
			pEven = pEven.Next
		}
		q = q.Next
		num++
	}

	pEven.Next = nil
	pOdd.Next = h2.Next
	return h1.Next
}