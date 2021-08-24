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

// 思路:将奇数、偶数的结点直接指向下一个奇数、偶数结点
// 然后把奇数的末尾连接偶数的开头
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    odd := head
    even := head.Next
    evenHead := even
    // even偶数结点更靠后，先到达末尾
    for even != nil && even.Next != nil {
        odd.Next = odd.Next.Next
        even.Next = even.Next.Next
        odd = odd.Next
        even = even.Next
    }
    odd.Next = evenHead
    return head
}
