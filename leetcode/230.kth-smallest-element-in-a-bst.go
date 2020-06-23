import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	// 寻找一个二叉搜索树的第k小的元素
	// 思路1：利用BST中序遍历生成一个有序数组，返回第k-1个数
	// 思路2：思路1耗费了额外O(n)的空间，可以使用栈对BST做非递归中序遍历，遍历到第k个数直接返回
	array := make([]int, 0)
	genSortedArray(root, &array)
	fmt.Printf("array:%v", array)
	if k <= len(array) {
		return array[k-1]
	}

	return array[len(array)-1]
}

func genSortedArray(node *TreeNode, array *[]int) {
	if node == nil {
		return
	}

	genSortedArray(node.Left, array)

	*array = append(*array, node.Val)

	genSortedArray(node.Right, array)
}