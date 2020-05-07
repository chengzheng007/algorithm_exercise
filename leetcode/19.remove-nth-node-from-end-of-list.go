/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = head

	tail := newHead
	for i := 0; i < n; i++ {
		tail = tail.Next
		// 节点数不够n
		if tail == nil {
			return head
		}
	}

	// 设定好，要删除倒数第3个节点，那么tail与preLastN之间的节点数为n-1
	// 例如总共5个节点，for循环后tail应指向第5个节点，preLastN则指向头结点
	// 删除的时候preLastN则指向头结点指向他的下一个即可
	// 也就是tail要指向链表中的第n个节点（不含头结点）

	// 倒数第n个节点的前驱节点
	preLastN := newHead
	for tail.Next != nil {
		preLastN = preLastN.Next
		tail = tail.Next
	}
	// 删除节点
	preLastN.Next = preLastN.Next.Next
	return newHead.Next
}