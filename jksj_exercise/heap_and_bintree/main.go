package main

import "fmt"

func main() {
	list := []int{4,6,1,7,2,9,5,8,3}
	heapSort(list, len(list))
	fmt.Printf("list:%v\n", list)

	list = []int{4,6,1,7,2,9,5,8,3}
	topKElems := topK(list, 3)
	fmt.Printf("topKElems:%v\n", topKElems)

	dlist := [][]int{
		[]int{1,4,5},
		[]int{1,3,4},
		[]int{2,6},
	}
	ret := mergeKList(dlist)
	fmt.Printf("mergeKList ret:%v\n", ret)
}