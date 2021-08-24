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

// 另一种巧妙解法，空间复杂度为O(1)
// 假设链表A的头节点到相交点的距离是a，链表B的头节点到相交点的距离是b，相交点
// 到链表终点的距离为c。我们使用两个指针，分别指向两个链表的头节点，并以相同的速度前进，
// 若到达链表结尾，则移动到另一条链表的头节点继续前进。按照这种前进方法，两个指针会在
// a+b+c次前进后同时到达相交节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1 := headA
    p2 := headB
    // 若不存在交点，相当于最后共同访问到的c1都是null
    for p1 != p2 {
        if p1 != nil {
            p1 = p1.Next
        } else {
            p1 = headB
        }
        if p2 != nil {
            p2 = p2.Next
        } else {
            p2 = headA
        }
    }
    return p1
}
