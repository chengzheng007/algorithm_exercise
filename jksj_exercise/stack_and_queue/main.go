package main

import "fmt"

func main() {
	fmt.Println("stack test:")
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Print()
	snode := stack.Pop()
	if snode != nil {
		fmt.Printf("stack pop elem:%v\n", snode.Val)
	}
	stack.Print()
	stack.Push(3)
	stack.Print()

	fmt.Println("queue test:")
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Print()
	qnode := queue.Pop()
	if qnode != nil {
		fmt.Printf("stack pop elem:%v\n", qnode.Val)
	}
	queue.Print()
	queue.Push(5)
	queue.Print()
	qnode = queue.Pop()
	if qnode != nil {
		fmt.Printf("stack pop elem:%v\n", qnode.Val)
	}
	queue.Print()

	list := []int{1,2,3}
	fullPermutation(list, 0, len(list))
}