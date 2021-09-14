/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    max := math.MinInt64
    // max记录寻找过程中最大的路径和
    findMaxPathSum(root, &max)
    return max
}

// 后续遍历
func findMaxPathSum(root *TreeNode, max *int) int {
    if root == nil {
        return 0
    }
    // 左右子树的值只有大于0，加上当前的值才能更大，否则只有变小的趋势
    // 故左右子树的值为负的需舍弃
    left := findMaxPathSum(root.Left, max)
    if left < 0 {
        left = 0
    }
    
    right := findMaxPathSum(root.Right, max)
    if right < 0 {
        right = 0
    }
    
    temp := left + right + root.Val
    if temp > *max {
        *max = temp
    }
    
    if left > right {
        return left + root.Val
    }
    return right + root.Val
}
