/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	newh := new(ListNode)
	newh.Next = head
	reverseRecur(head, newh)
	return newh.Next
}

// 尾递归处理
func reverseRecur(p *ListNode, head *ListNode) {
	if p == nil || head == nil {
		return
	}
	// 尾节点特殊处理
	if p.Next == nil {
		head.Next = p
		return
	}
	reverseRecur(p.Next, head)
	// 递归返回时，将当前节点的next节点的next指针指向自己
	p.Next.Next = p
	// 该行主要功能是将链表中第一个节点的next指针置为空，可能会陷入死循环
	p.Next = nil
}

// 更简单的尾递归
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    node := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return node
}

// 非递归-从头遍历形式
func reverseList(head *ListNode) *ListNode {
   var prev, next *ListNode
    for head != nil {
        next = head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}
