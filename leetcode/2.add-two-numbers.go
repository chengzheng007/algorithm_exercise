/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	var num1, num2, carry int
	p := l1
	q := l2
	for p != nil || q != nil {
		num1 = 0
		if p != nil {
			num1 = p.Val
			p = p.Next
		}
		num2 = 0
		if q != nil {
			num2 = q.Val
			q = q.Next
		}
		sum := num1 + num2 + carry
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if sum >= 10 {
			carry = 1
			// 当最后两个数和>=10时，需要在最后再加一个节点
			// 如果是中间节点>=10，这里的Next会被覆盖掉，是没有影响的
			curr.Next = &ListNode{Val: 1}
		} else {
			carry = 0
		}
	}
	return head.Next
}

