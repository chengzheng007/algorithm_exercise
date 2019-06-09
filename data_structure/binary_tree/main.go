package main

import (
	"fmt"
)

// 链式存储
type Node struct {
	Data string
	LChild *Node
	RChild *Node
}

func main() {
	root := initBTree()
	fmt.Printf("pre order traversal(recursive):")
	preOrderTraversalRec(root, 0)

	fmt.Printf("in order traversal(recursive):")
	inOrderTraversalRec(root, 0)

	fmt.Printf("post order traversal(recursive):")
	postOrderTraversalRec(root, 0)

	fmt.Printf("pre order traversal(non-recursive):")
	preOrderTraversal(root)

	fmt.Printf("in order traversal(non-recursive):")
	inOrderTraversal(root)

	fmt.Printf("post order traversal(non-recursive):")
	postOrderTraversal(root)

	fmt.Printf("level order traversal:")
	levelOrder(root)

	fmt.Printf("tree width:%d\n", treeWidth(root))

	fmt.Printf("tree height:%d\n", treeHeightRec(root))
	fmt.Printf("tree height treeHeightByQueue:%d\n", treeHeightByQueue(root))
}

func initBTree() *Node {
	root := &Node{Data:"A"}
	root.LChild = &Node{Data:"B"}
	root.RChild = &Node{Data:"C"}
	root.LChild.LChild = &Node{Data:"D"}
	//root.LChild.RChild = &Node{Data:"E"}
	//root.RChild.LChild = &Node{Data:"F"}
	root.RChild.RChild = &Node{Data:"G"}
	return root
}

// 遍历，递归实现
func preOrderTraversalRec(root *Node, dep int) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.Data)
	dep++
	preOrderTraversalRec(root.LChild, dep)
	preOrderTraversalRec(root.RChild, dep)
	if dep == 1 {
		fmt.Printf("\n")
	}
}

func inOrderTraversalRec(root *Node, dep int) {
	if root == nil {
		return
	}
	dep++
	inOrderTraversalRec(root.LChild, dep)
	fmt.Printf("%v ", root.Data)
	inOrderTraversalRec(root.RChild, dep)
	if dep == 1 {
		fmt.Printf("\n")
	}
}

func postOrderTraversalRec(root *Node, dep int) {
	if root == nil {
		return
	}
	dep++
	postOrderTraversalRec(root.LChild, dep)
	postOrderTraversalRec(root.RChild, dep)
	fmt.Printf("%v ", root.Data)
	if dep == 1 {
		fmt.Printf("\n")
	}
}

// 遍历：非递归
// 递归的地方，就是入栈的地方。由于先进后出，所以重点要看哪个先入栈
func preOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	stack := InitArrayStack()
	stack.Push(root)

	for !stack.Empty() {
		node, _ := stack.Pop()
		fmt.Printf("%v ", node.Data)

		if node.RChild != nil {
			stack.Push(node.RChild)
		}
		if node.LChild != nil {
			stack.Push(node.LChild)
		}
	}
	fmt.Printf("\n")
}

func inOrderTraversal(root *Node) {
	stack := InitArrayStack()
	node := root

	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.LChild
		}

		node, _ = stack.Pop()
		
		fmt.Printf("%v ", node.Data)
		
		node = node.RChild
	}
	fmt.Printf("\n")
}

func postOrderTraversal(root *Node) {
	stack := InitArrayStack()
	var nodeRC *Node
	node := root
	
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			nodeRC = node.RChild
			node = node.LChild
			if node == nil {
				node = nodeRC
			}
		}

		node, _ = stack.Pop()
		fmt.Printf("%v ", node.Data)

		topNode := stack.Top()
		if topNode != nil && topNode.LChild == node {
			node = topNode.RChild
		} else {
			node = nil
		}
	}
	fmt.Printf("\n")
}

func levelOrder(root *Node) {
	if root == nil {
		return
	}
	queue := NewQueue(9)
	queue.Enter(root)
	for !queue.Empty() {
		node, _ := queue.Pop()
		if node.LChild != nil {
			queue.Enter(node.LChild)
		}
		if node.RChild != nil {
			queue.Enter(node.RChild)
		}
	}
	fmt.Printf("\n")
}

// 求树宽度
func treeWidth(root *Node) int {
	if root == nil {
		return 0
	}
	
	var (
		width = 1
		tempWidth int
		temp = 1
		tempNum int
	)

	queue := NewQueue(10)
	queue.Enter(root)

	for !queue.Empty() {
		if queue.Head() != nil {
			temp--
			tempWidth++
		}      

		if queue.Head().LChild != nil {
			tempNum++
			queue.Enter(queue.Head().LChild)
		}
		if queue.Head().RChild != nil {
			tempNum++
			queue.Enter(queue.Head().RChild)
		}

		if temp == 0 {
			if tempWidth > width {
				width = tempWidth
			}
			temp = tempNum
			tempNum = 0
			tempWidth = 0
		}
		queue.Pop()
	}
	return width
}

// 树高度-递归
func treeHeightRec(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := treeHeightRec(root.LChild)+1
	rightHeight := treeHeightRec(root.RChild)+1
	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}

// 借助队列
func treeHeightByQueue(root *Node) int {
	if root == nil {
		return 0
	}
	queue := NewQueue(12)
	queue.Enter(root)

	height := 0
	levelNodeNum := 0
	temp := 1
	
	for !queue.Empty() {
		if queue.Head() != nil {
			temp--
		}

		if queue.Head().LChild != nil {
			queue.Enter(queue.Head().LChild)
			levelNodeNum++
		}
		if queue.Head().RChild != nil {
			queue.Enter(queue.Head().RChild)
			levelNodeNum++
		}

		if temp == 0 {
			temp = levelNodeNum
			height++
			levelNodeNum = 0
		}

		queue.Pop()
	}
	
	return height
}