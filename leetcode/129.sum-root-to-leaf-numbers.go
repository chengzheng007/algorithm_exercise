/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return helper(root, 0)
}

func helper(root *TreeNode, pathSum int) int {
    if root == nil {
        return 0
    }
    
    pathSum = pathSum*10 + root.Val
    
    if root.Left == nil && root.Right == nil {
        return pathSum
    }
    
    return helper(root.Left, pathSum) + helper(root.Right, pathSum)
}
