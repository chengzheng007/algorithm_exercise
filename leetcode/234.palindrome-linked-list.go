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
    
    // 如果链表为奇数个节点，slow刚好指向中间节点（1->2->3->2->1，fast指向最后的1结点，slow指向3）
    // 此时如果前后断开反转，比较的将是:前半部分1->2，后半部分1->2->3，遍历时需以前半部分链表结点是否为空为条件判断
    // 如果是偶数个，1->2->2->1，slow将指向后一个2，前后断开，前半部分1->2，后半部分2->1
    // 为了统一以slow指向结点为判断依据，这里后移一个，便于下面统一处理
    if fast != nil {
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
