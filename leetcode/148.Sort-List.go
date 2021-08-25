/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 链表排序，归并算法，尾递归合并，先到底层再合并退回
func sortList(head *ListNode) *ListNode {
    // 递归结束条件
    if head == nil || head.Next == nil {
        return head
    }
    mid := splitList(head)
    // 先递归到底层，处理完再回退
    left := sortList(head)
    right := sortList(mid)
    return mergeList(left, right)
}

func mergeList(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    } else if l1.Val < l2.Val {
        l1.Next = mergeList(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeList(l1, l2.Next)
        return l2
    }
}

// get middle node and disconnect them
func splitList(node *ListNode) *ListNode {
    if node == nil || node.Next == nil {
        return node
    }
    var (
        p = node
        mid *ListNode
    )
    for p != nil && p.Next != nil {
        if mid == nil {
            mid = p
        } else {
            mid = mid.Next
        }
        p = p.Next.Next
    }
    // 循环结束后，mid指向中间节点的前驱
    if mid == nil {
        return mid
    }
    
    next := mid.Next
    mid.Next = nil // 断开后半部分链表
    mid = next
    return mid
}
