/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	// p待删节点，pp待删节点前驱
	var pp *TreeNode
	p := root
	for p != nil && p.Val != key {
		pp = p
		if p.Val < key {
			p = p.Right
		} else {
			p = p.Left
		}
	}

	// 节点不存在
	if p == nil {
		return root
	}

	// 节点拥有左右子节点
	if p.Left != nil && p.Right != nil {
		// pMin指向右子树中最左子节点(右子树中最小值节点)，ppMin指向其前驱
		ppMin := p
		pMin := p.Right
		for pMin.Left != nil {
			ppMin = pMin
			pMin = pMin.Left
		}

		// 将右子树中最小子节点值替换掉待删节点
		// 接下来删除pMin，pMin可能为叶子节点或只有一个子节点
		p.Val = pMin.Val
		p = pMin
		pp = ppMin
	}

	var child *TreeNode
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	}

	if pp == nil { // 删除的是根节点
		return child
	} else if pp.Left == p { // 删除其他节点，接上节点的子节点到前驱
		pp.Left = child
	} else {
		pp.Right = child
	}

	return root
}