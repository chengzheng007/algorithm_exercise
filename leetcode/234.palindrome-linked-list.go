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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 寻找中间节点slow，若是偶数，slow指向中间偏后一个节点
    if fast != nil { // 表示链表为奇数个节点，slow刚好指向中间节点（1->2->3->2->1，fast指向最后的1，slow指向3）
        slow = slow.Next // 将slow移动至下一个节点 
    }
    
    // slow及之后进行反转，反转算法自动断开连接
    slow = reverse(slow)
    fast = head
        
    // 两部分遍历比较，奇数结点数的话，slow指向的链表先到末尾
    for slow != nil {
        if slow.Val != fast.Val {
            return false
        }
        slow = slow.Next
        fast = fast.Next
    }
    
    return true
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    node := reverse(head.Next)
    head.Next.Next = head
    head.Next = nil
    return node
}
