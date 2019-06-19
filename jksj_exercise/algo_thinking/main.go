package main

import "fmt"

func main() {
	fmt.Println("reverse order counting test:")
	list := []int{5,3,6,1,2,4,9}
	count := reverseOrderCounting(list)
	fmt.Printf("reverse order count:%d\n", count)
	fmt.Printf("after sort:%v\n", list)

	list = []int{5,3,6,1,2,4,9}
	count = brouteReverseOrderCountint(list)
	fmt.Printf("validate count:%d\n", count)
}
