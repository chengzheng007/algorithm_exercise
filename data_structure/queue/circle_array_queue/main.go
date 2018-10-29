package main

import (
	"fmt"
)

// 循环队列，队尾进队头出

type arrayQueue struct {
	head, tail int
	capacity int // 最大容量
	list []int
}


func main() {
	q := NewArrayQueue(8)

	for i := 1; i < 10; i++ {
		ok := q.Enqueue(i)
		fmt.Printf("enqueue ret:%v, queue len:%v q.head:%v q.tail:%v\n", ok, q.Len(), q.head, q.tail)
	}

	fmt.Println("===================")

	val, ok := q.Dequeue()
	fmt.Printf("dequeue ret:%v, val:%v, queue len:%v q.head:%v q.tail:%v\n", ok, val, q.Len(), q.head, q.tail)

	val, ok = q.Dequeue()
	fmt.Printf("dequeue ret:%v, val:%v, queue len:%v q.head:%v q.tail:%v\n", ok, val, q.Len(), q.head, q.tail)

	val, ok = q.Dequeue()
	fmt.Printf("dequeue ret:%v, val:%v, queue len:%v q.head:%v q.tail:%v\n", ok, val, q.Len(), q.head, q.tail)

	val, ok = q.Dequeue()
	fmt.Printf("dequeue ret:%v, val:%v, queue len:%v q.head:%v q.tail:%v\n", ok, val, q.Len(), q.head, q.tail)

	val, ok = q.Dequeue()
	fmt.Printf("dequeue ret:%v, val:%v, queue len:%v q.head:%v q.tail:%v\n", ok, val, q.Len(), q.head, q.tail)

	fmt.Println("===================")

	ok = q.Enqueue(10)
	fmt.Printf("enqueue ret:%v, queue len:%v q.head:%v q.tail:%v\n", ok, q.Len(), q.head, q.tail)
	ok = q.Enqueue(11)
	fmt.Printf("enqueue ret:%v, queue len:%v q.head:%v q.tail:%v\n", ok, q.Len(), q.head, q.tail)


	fmt.Println("===================")
	
	for i := 20; i < 26; i++ {
		ok := q.Enqueue(i)
		fmt.Printf("enqueue ret:%v, queue len:%v q.head:%v q.tail:%v\n", ok, q.Len(), q.head, q.tail)
	}

	fmt.Println("===================")
	
	
}

func NewArrayQueue(cap int) *arrayQueue {
	return &arrayQueue{
		capacity:cap,
		list:make([]int, cap),
	}
}

func (q *arrayQueue) Enqueue(x int) bool {
	// 队满
	if (q.tail + 1) % q.capacity == q.head {
		return false
	}
	q.list[q.tail] = x
	// 实际q.tail所在位置无元素，浪费一个存储空间
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *arrayQueue) Dequeue() (int, bool) {
	// 队空
	if q.head == q.tail {
		return -1, false
	}
	x := q.list[q.head]
	q.head = (q.head + 1) % q.capacity
	return x, true
}

func (q *arrayQueue) Len() int {
	return (q.tail + q.capacity - q.head) % q.capacity
}