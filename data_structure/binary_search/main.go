package main

import "fmt"

func main() {
	list := []int{1,2,3,4,5,6,7,8,9}
	x := 7
	fmt.Printf("search %d index:%d\n", x, binSearch(list, x))
	fmt.Printf("recursive search %d index:%d\n", x, binSearchRecursive(list, 0, len(list)-1, x))

	x = 9
	fmt.Printf("search %d index:%d\n", x, binSearch(list, x))
	fmt.Printf("recursive search %d index:%d\n", x, binSearchRecursive(list, 0, len(list)-1, x))
	
	x = 1
	fmt.Printf("search %d index:%d\n", x, binSearch(list, x))
	fmt.Printf("recursive search %d index:%d\n", x, binSearchRecursive(list, 0, len(list)-1, x))
	
	x = 55
	list = []int{1,2,3}
	fmt.Printf("search %d index:%d\n", x, binSearch(list, x))
	fmt.Printf("recursive search %d index:%d\n", x, binSearchRecursive(list, 0, len(list)-1, x))
	
	fmt.Println("==============================")
	
	list = []int{5}
	x = 1
	fmt.Printf("search %d index:%d\n", x, binSearch(list, x))
	fmt.Printf("recursive search %d index:%d\n", x, binSearchRecursive(list, 0, len(list)-1, x))
	
	x = 5
	fmt.Printf("search %d index:%d\n", x, binSearch(list, x))
	fmt.Printf("recursive search %d index:%d\n", x, binSearchRecursive(list, 0, len(list)-1, x))
	
}

func binSearch(list []int, x int) int {
	start := 0
	end := len(list) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if list[mid] == x {
			return mid
		} else if list[mid] < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func binSearchRecursive(list []int, start, end, x int) int {
	if start > end {
		return -1
	}

	mid := start + (end - start) >> 1
	if list[mid] == x {
		return mid
	} else if list[mid] < x {
		return binSearchRecursive(list, mid+1, end, x)
	} else {
		return binSearchRecursive(list, start, mid-1, x)
	}
}