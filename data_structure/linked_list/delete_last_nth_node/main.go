package main

import (
	"errors"
	"fmt"
)

// 删除单链表倒数第n个节点
// 时间复杂度：O(n)，需遍历

type Node struct {
	Data int
	Next *Node
}

func main() {
	list := []int{1,2,3,4,5}
	head, err := initNodeList(list)
	if err != nil {
		fmt.Printf("initNodeList(list) error(%v)\n", err)
		return
	}
	print(head)
	fmt.Println("===============")

	delRet := delLastNthNode(head, 4)
	fmt.Println("after delete:")
	print(head)
	fmt.Printf("delete success? %t\n", delRet)
}


func initNodeList(list []int) (*Node, error) {
	if len(list) == 0 {
		return nil, nil
	}
	head := new(Node)
	p := head
	for _, val := range list {
		node := new(Node)
		if node == nil {
			return nil, errors.New("no spare space")
		}
		node.Data = val
		p.Next = node
		p = node

	}
	return head, nil
}

func print(head *Node) {
	i := 1
	for head.Next != nil {
		fmt.Printf("node id: %d, value:%d, next:%p, addr:%p\n", i, head.Next.Data, head.Next.Next, head.Next)
		i++
		head = head.Next
	}
}

// return true:删除成功  false:失败，没有该节点
func delLastNthNode(head *Node, n int) bool {
	if head == nil {
		return false
	}
	
	// p先走n步，达到第n个节点
	p := head
	for i := 0; i < n; i++ {
		p = p.Next
		if p == nil {          // 节点数不够
			return false
		}
	}

	fmt.Printf("p value:%d, next:%p, addr:%p\n", p.Data, p.Next, p)

	prevDel := head
	for p.Next != nil {
		p = p.Next
		prevDel = prevDel.Next
	}
	// p走到尾节点，prevDel为待删除节点的前驱节点
	fmt.Printf("prevDel value:%d, next:%p, addr:%p\n", prevDel.Data, prevDel.Next, prevDel)
	
	// 将待删除节点摘除
	prevDel.Next = prevDel.Next.Next
	
	return true
}
