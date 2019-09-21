/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapA := make(map[*ListNode]struct{}, 0)
	p := headA
	for p != nil {
		mapA[p] = struct{}{}
		p = p.Next
	}
	p = headB
	for p != nil {
		if _, ok := mapA[p]; ok {
			return p
		}
		p = p.Next
	}
	return nil
}

