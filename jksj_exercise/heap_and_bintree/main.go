package main

import "fmt"

func main() {
	list := []int{4,6,1,7,2,9,5,8,3}
	heapSort(list, len(list))
	fmt.Printf("list:%v\n", list)

	list = []int{4,6,1,7,2,9,5,8,3}
	topKElems := topK(list, 4)
	fmt.Printf("topKElems:%v\n", topKElems)
}