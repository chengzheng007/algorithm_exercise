/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return genTreeList(1, n)
}

// 1..n(视为递增数组)构建所有唯一的二叉搜索树(BST)，数目等于卡特兰数
func genTreeList(start, end int) []*TreeNode {
	list := make([]*TreeNode, 0)
	if start > end {
		list = append(list, nil)
	}
	// every num as root to build a tree
	for i := start; i <= end; i++ {
		left := genTreeList(start, i-1)
		right := genTreeList(i+1, end)
		// left/right are several tree node which has been root node of each child tree
		for _, lnode := range left {
			for _, rnode := range right {
				// 笛卡尔积，每个左子树分别与右子树对应于root组合，作为一种可选结构返回
				root := &TreeNode{Val: i}
				root.Left = lnode
				root.Right = rnode
				list = append(list, root)
			}
		}
	}
	return list
}