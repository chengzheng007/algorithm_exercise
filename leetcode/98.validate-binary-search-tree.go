/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}

func validate(node, leftMax, rightMin *TreeNode) bool {
    if node == nil {
        return true
    }
    // leftMax:最大值，左子树不能超过该值
    // rightMin:最小值，右子树不能小于该值
    if leftMax != nil && node.Val >= leftMax.Val {
        return false
    }
    if rightMin != nil && node.Val <= rightMin.Val {
        return false
    }
    // 递归定义，左子树不只是小于父节点，还得小于祖父等节点，右子树同理
    // 右子树中的节点不仅大于父节点，还得大于祖父等节点(坑)
    return validate(node.Left, node, rightMin) && validate(node.Right, leftMax, node)
}
