/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	newh := new(ListNode)
	newh.Next = head
	middleNode := newh
	endNode := newh
	// 寻找中间节点和结束节点
	// 注意：循环结束，endNode不一定指向结尾，可能为nil，当节点个数为计数时
	for endNode != nil && endNode.Next != nil {
		middleNode = middleNode.Next
		endNode = endNode.Next.Next
	}

	// 翻转链表后半部分
	p := middleNode
	var q *ListNode
	if p != nil {
		q = p.Next
	}
	for q != nil {
		// 翻转遍历时查找真正的尾节点
		if q != nil {
			endNode = q
		}
		r := q.Next
		q.Next = p
		p = q
		q = r
	}

	// 从两边向中间遍历比较，注意break退出条件
	p = head
	q = endNode
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if p == middleNode || q == middleNode {
			break
		}
		p = p.Next
		q = q.Next
	}
	return true
}

