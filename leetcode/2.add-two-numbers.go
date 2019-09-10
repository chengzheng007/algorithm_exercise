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
	p := l1
	q := l2
	carry := 0
	newList := new(ListNode)
	curr := newList
	for p != nil || q != nil {
		num1 := 0
		if p != nil {
			num1 = p.Val
		}
		num2 := 0
		if q != nil {
			num2 = q.Val
		}
		sum := num1 + num2 + carry

		curr.Next = &ListNode{Val:sum%10}
		curr = curr.Next

		carry = sum/10
		// 如果两数的最后一位如果两数之和大于10，在链表最后需要再加一位
		// 如果不是最后一位的和也>=10，则该数会被覆盖，可以忽略
		if carry > 0 {
			curr.Next = &ListNode{Val:carry}
		}

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	return newList.Next
}

