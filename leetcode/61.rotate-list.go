/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	// 利用和翻转数组一样的思想
	// 先确定真正翻转的个数，如果k比链表节点数还大，就是k对length取模

	// 如果为空或只有一个节点，再怎么翻转也不会边
	if head == nil || head.Next == nil {
		return head
	}

	// 统计链表节点个数
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}

	if n == 0 {
		return head
	}

	// 确定真实翻转个数
	k = k % n
	if k == 0 { // 无需翻转
		return head
	}

	// 先将整个链表翻转
	newHead := reverseList(head)

	// 寻找第k个节点，分别翻转该节点前后的链表，再连接到一起
	kNode := newHead
	for i := 1; i < k && kNode != nil; i++ {
		kNode = kNode.Next
	}

	if kNode == nil { // 不够凑成一组
		return newHead
	}

	// 后半部分
	secondHalf := kNode.Next
	// 将前半部分从主链断开
	kNode.Next = nil

	p = newHead // 记录第一次翻转的第一个结点

	newFirst := reverseList(newHead)
	newSecond := reverseList(secondHalf)
	// 将前后两部分连接起来
	p.Next = newSecond
	return newFirst
}

func reverseList(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}
	newh := reverseList(h.Next)
	h.Next.Next = h
	h.Next = nil
	return newh
}

