/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return genTree(nums, 0, len(nums)-1)
}

func genTree(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}

	maxNum := nums[l]
	maxIdx := l

	// 寻找最大值
	for i := l + 1; i <= r; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIdx = i
		}
	}

	node := &TreeNode{Val: maxNum}
	node.Left = genTree(nums, l, maxIdx-1)
	node.Right = genTree(nums, maxIdx+1, r)
	return node
}