/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	var encouter *ListNode
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		// 两个指针相遇表示有环
		if p1 == p2 {
			encouter = p1
			break
		}
	}
	// 不存在环
	if p2 == nil || p2.Next == nil {
		return nil
	}

	// 寻找换入口
	// 分析可参考：https://blog.csdn.net/wuzhekai1985/article/details/6725263
	p1 = head
	p2 = encouter
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

