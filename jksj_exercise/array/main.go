package main

import "fmt"

func main() {
	arr1 := []int{1,3,5,7}
	arr2 := []int{2,4,6,8,10}
	newArr := mergeTwoSequentialArray(arr1, arr2)
	fmt.Printf("newArr:%v\n", newArr)

	arr1 = []int{1,7,20}
	arr2 = []int{2,4,8,10,13,15,19}
	newArr = mergeTwoSequentialArray(arr1, arr2)
	fmt.Printf("newArr:%v\n", newArr)

	list := make([][]int, 3)
	list[0] = []int{1,2,3,9}
	list[1] = []int{4,5}
	list[2] = []int{6,7,8}
	fmt.Printf("roundTripMerge:%v\n", roundTripMerge(list))
}
