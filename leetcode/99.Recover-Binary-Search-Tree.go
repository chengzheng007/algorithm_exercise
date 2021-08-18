
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    // pointers[0]:prev node 
    // pointers[1]:first node 
    // pointers[2]:second node 
    pointers := make([]*TreeNode, 3)
    pointers[0] = &TreeNode{Val:math.MinInt64}
    // 小技巧：将指针放在数组里面，利用go数组特性修改
    traverse(root, pointers)
    
    if pointers[1] != nil && pointers[2] != nil {
        temp := pointers[1].Val
        pointers[1].Val = pointers[2].Val
        pointers[2].Val = temp     
    }
}

func traverse(node *TreeNode, pointers []*TreeNode) {
    if node == nil {
        return
    }
   
    traverse(node.Left, pointers)
    
    if pointers[1] == nil && pointers[0].Val >= node.Val {
        pointers[1] = pointers[0]
    }
    // 前一个if将pointers[1]赋值为非空，紧接着会运行这里
    // 相当于first = prev, second = curr_node
    if pointers[1] != nil && pointers[0].Val >= node.Val {
        pointers[2] = node
    }
        
    pointers[0] = node
    
    traverse(node.Right, pointers)
}
