/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    traverse(root, &res, 0)
    // 偶数层元素反转
    for i := 1; i < len(res); i += 2 {
        start := 0
        end := len(res[i])-1
        for start < end {
            res[i][start], res[i][end] = res[i][end], res[i][start]
            start++
            end--
        }
    }
    return res
}

func traverse(root *TreeNode, res *[][]int, depth int) {
    if root == nil {
        return
    }
    if len(*res) <= depth { 
        *res = append(*res, make([]int, 0))
    }
    // 每一层追加到对应的切片
    (*res)[depth] = append((*res)[depth], root.Val)
    traverse(root.Left, res, depth+1)
    traverse(root.Right, res, depth+1)
}
