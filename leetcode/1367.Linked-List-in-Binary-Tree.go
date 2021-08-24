/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
    // check from root node if it's satisfied
    match := check(root, head) 
    if match {
        return true
    }
    // if root not satisfied, checkout root's child node to head
    if root != nil && !match {
        return isSubPath(head, root.Left) || isSubPath(head, root.Right)
    }
    
    return false
}

func check(treeNode *TreeNode, listNode *ListNode) bool {
    // linked list have validate to the end
    if listNode == nil {
        return true
    }
    if treeNode != nil && treeNode.Val == listNode.Val {
        return check(treeNode.Left, listNode.Next) || check(treeNode.Right, listNode.Next)
    }
    return false
}
