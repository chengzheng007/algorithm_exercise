/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 利用二叉搜索树的特性，p、q要么分布在祖先的两边，要么在一边
    // 假设在当前结点两边，那么当前结点即为祖先，直接返回
    // 如果均分布在左边，p、q值均小于当前结点，分布右边则均大于当
    // 前结点
    // 当然，题目没考虑结点为空、结点不存在的情况
    
    for root != nil && (root.Val-p.Val)*(root.Val-q.Val) > 0 {
        if p.Val < root.Val {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return root
}

/*
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
*/
