/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || p == root || q == root {
        return root
    }
    
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    // 查询子树中p、q结点，若能查到，则在后续遍历回来时
    // 利用递归越深入越早回来的特点，如果分别在结点的左右子树中
    // 最近祖先就是当前节点
    if left != nil && right != nil {
        return root
    }
    
    if left != nil {
        return left
    }
    return right  
}
