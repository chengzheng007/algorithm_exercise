/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	middleNode := head
	endNode := head
	for endNode != nil && endNode.Next != nil {
		middleNode = middleNode.Next
		endNode = endNode.Next.Next
	} // end: middleNode为中间节点

	// 翻转middleNode至尾节点间的节点
	p := middleNode
	var q *ListNode
	if p != nil {
		q = p.Next
	}
	// 注意节点顺序
	for q != nil {
		// endNode为真正的尾节点，第一个for循环只能保证middleNode正确
		endNode = q

		r := q.Next
		q.Next = p
		p = q
		q = r
	}

	p = head
	q = endNode
	// 从两头开始遍历，将后半部分节点逐个插入前半部分
	for p != nil && q != nil {
		if p == middleNode || q == middleNode {
			break
		}
		tempP := p.Next
		tempQ := q.Next
		q.Next = tempP
		p.Next = q
		p = tempP
		q = tempQ
	}
	// 不要忘了将中间节点的Next置空，否则形成循环链表
	if middleNode != nil {
		middleNode.Next = nil
	}
}

