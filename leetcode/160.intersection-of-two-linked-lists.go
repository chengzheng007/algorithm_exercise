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
	p1 := headA
	var sameNode *ListNode
	nodeMapOfB := make(map[*ListNode]bool, 0)
	traversed := false
	for p1 != nil {
		// 空间换时间
		if traversed {
			if _, ok := nodeMapOfB[p1]; ok {
				sameNode = p1
				break
			}
		} else {
			p2 := headB
			for p2 != nil {
				if p1 == p2 {
					sameNode = p2
					break
				}
				nodeMapOfB[p2] = true
				p2 = p2.Next
			}
			traversed = true
		}

		p1 = p1.Next
	}
	return sameNode
}

