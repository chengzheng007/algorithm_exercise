/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func levelOrderBottom(root *TreeNode) [][]int {
//     if root == nil {
// 		return nil
// 	}
//     // 借助队列的BFS搜索模式
// 	queue := list.New()
// 	queue.PushBack(root)
// 	res := make([][]int, 0)
// 	for queue.Len() > 0 {
// 		size := queue.Len()
// 		var subList []int

// 		// 处理改一整层节点，以该层的长度作为记录数
//         // 这里是精华：怎么知道一层是多少？每次将
// 		for i := 0; i < size; i++ {
// 			frontElem := queue.Front()
// 			node, ok := frontElem.Value.(*TreeNode)
// 			if !ok || node == nil {
// 				continue
// 			}

// 			if node.Left != nil {
// 				queue.PushBack(node.Left)
// 			}
// 			if node.Right != nil {
// 				queue.PushBack(node.Right)
// 			}
// 			// 将数据记录到数组
// 			subList = append(subList, node.Val)
// 			// 删除处理节点，接着处理这一层的下一个节点
// 			queue.Remove(frontElem)
// 		}
// 		res = append(res, subList)
// 	}
// 	// reverse
// 	i := 0
// 	j := len(res) - 1
// 	for i < j {
// 		res[i], res[j] = res[j], res[i]
// 		i++
// 		j--
// 	}
// 	return res
// }

// DFS模式
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	levelMaker(root, &res, 0)
	i := 0
	j := len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func levelMaker(node *TreeNode, res *[][]int, level int) {
	if node == nil {
		return
	}
	// level表示遍历的层级或者说深度
	// 最开始时res是没有数据的，长度为0，但是递归深入到下一层时
	// 需要先在这一层占个坑，下面才能把数据正确放在res数组的
	// 层级对应的数组后面（二维数组）
	if len(*res) <= level {
		*res = append(*res, make([]int, 0))
	}
	levelMaker(node.Left, res, level+1)
	levelMaker(node.Right, res, level+1)
	// 将数据append到对应的层级的数组
	(*res)[level] = append((*res)[level], node.Val)
}