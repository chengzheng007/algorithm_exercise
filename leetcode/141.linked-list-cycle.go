/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for {
		// use a slow and a fast two pointer, every loop slow pointer move one step,
		// fast pointer move two step, if fast pointer is nil, then there's no cycle,
		// if finally fast pointer encouter with slow pointer, then there is cycle.
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
	}
	return false
}

