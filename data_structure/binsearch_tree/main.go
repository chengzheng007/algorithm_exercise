package main

import (
	"fmt"
)

// 二叉搜索树

type Node struct {
	Data int
	LChild *Node
	RChild *Node
}

type BinSearchTree struct {
	Size int
	Root *Node
}

func main() {

	bst := &BinSearchTree{}

	list := []int{33,16,50,13,18,34,58,15,17,25,51,66,19,27,55}
	
	for _, num := range list {
		insert(bst, num)
	}

	x := 6
	retNode := find(bst, x)
	fmt.Printf("retNodenode:%+v\n", retNode)

	fmt.Printf("inOrderTraversal:")
	inOrderTraversal(bst.Root)
	fmt.Printf("\n")

	//x = 55
	//delete(bst, x)
	//fmt.Printf("delete %d, inOrderTraversal:", x)
	//inOrderTraversal(bst.Root)
	//fmt.Printf("\n")
	//
	//
	//x = 13
	//delete(bst, x)
	//fmt.Printf("delete %d, inOrderTraversal:", x)
	//inOrderTraversal(bst.Root)
	//fmt.Printf("\n")
	//
	//x = 66
	//delete(bst, x)
	//fmt.Printf("delete %d, inOrderTraversal:", x)
	//inOrderTraversal(bst.Root)
	//fmt.Printf("\n")

	x = 18
	delete(bst, x)
	fmt.Printf("delete %d, inOrderTraversal:", x)
	inOrderTraversal(bst.Root)
	fmt.Printf("\n")
}

func insert(bst *BinSearchTree, data int) {
	if bst == nil {
		panic("BinSearchTree is nil")
		return
	}
	if bst.Root == nil {
		bst.Root = &Node{Data:data}
		bst.Size = 1
		return
	}
	
	p := bst.Root
	for p != nil {
		if data < p.Data {
			if p.LChild == nil {
				p.LChild = &Node{Data:data}
				bst.Size++
				return
			}
			p = p.LChild
		} else {
			if p.RChild == nil {
				p.RChild = &Node{Data:data}
				bst.Size++
				return
			}
			p = p.RChild
		}
	}		
}

func find(bst *BinSearchTree, x int) *Node {
	if bst == nil || bst.Root == nil {
		return nil
	}
	p := bst.Root
	for p != nil {
		if x == p.Data {
			return p
		} else if x < p.Data {
			p = p.LChild
		} else {
			p = p.RChild
		}
	}
	return nil
}

func inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.LChild)
	fmt.Printf("%v ", node.Data)
	inOrderTraversal(node.RChild)
}

// true: 删除成功
func delete(bst *BinSearchTree, x int) bool {
	if bst == nil || bst.Root == nil {
		return false
	}

	var pre *Node
	p := bst.Root
	for p != nil && p.Data != x {
		pre = p
		if x < p.Data {
			p = p.LChild
		} else {
			p = p.RChild
		}
	}

	//fmt.Printf("pre:%+v, p:%+v\n", pre, p)
	
	// 未找到该节点
	if p == nil {
		return false
	}

	//待删除节点有两个子节点
	if p.LChild != nil && p.RChild != nil {
		minP := p.RChild
		minPP := p // 表示minP的父节点
		
		for minP.LChild != nil {
			minPP = minP
			minP = minP.LChild
		}
		// 只替换数据部分，变成了删除minP节点
		p.Data = minP.Data
		p = minP
		pre = minPP
		//fmt.Printf("after transfom, pre:%+v, p:%+v\n", pre, p)
	}
	
	// 删除的节点是叶子节点或只有一个子节点
	var child *Node
	if p.LChild != nil {
		child = p.LChild
	} else if p.RChild != nil {
		child = p.RChild
	}

	if pre == nil {  // 删除节点为根节点
		bst.Root = nil
	} else if pre.LChild == p {
		pre.LChild = child
	} else {
		pre.RChild = child
	}
	p = nil
	return true
}