/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    depth := -1 // 注意初始化为-1，单个结点的情况
    val := 0
    findNode(root, 0, &depth, &val)
    return val
}

// 记录最深的节点和深度值
func findNode(root *TreeNode, currDepth int, depth, val *int) {
    if root == nil {
        return 
    }
    if currDepth > *depth {
        *depth = currDepth
        *val = root.Val
    }
    findNode(root.Left, currDepth+1, depth, val)
    findNode(root.Right, currDepth+1, depth, val)
}
