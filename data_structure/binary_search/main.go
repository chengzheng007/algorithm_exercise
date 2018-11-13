package main

import (
	"fmt"
)

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


	fmt.Println("====二分法求平方根====")
	var (
		number float32 = 2
		threshold float32 = 0.000001
	)
	
	ret, err := sqrtByBinSearch(number, threshold)
	if err != nil {
		fmt.Printf("sqrtByBinSearch(%f, %f) error(%v)\n", number, threshold, err)
	} else {
		fmt.Printf("sqrtByBinSearch(%f, %f) ret(%f)\n", number, threshold, ret)
	}

	number = 0.008
	ret, err = sqrtByBinSearch(number, threshold)
	if err != nil {
		fmt.Printf("sqrtByBinSearch(%f, %f) error(%v)\n", number, threshold, err)
	} else {
		fmt.Printf("sqrtByBinSearch(%f, %f) ret(%f)\n", number, threshold, ret)
	}
}

func binSearch(list []int, x int) int {
	start := 0
	end := len(list) - 1
	for start <= end {
		mid := start + ((end - start) >> 1)
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

	mid := start + ((end - start) >> 1)
	if list[mid] == x {
		return mid
	} else if list[mid] < x {
		return binSearchRecursive(list, mid+1, end, x)
	} else {
		return binSearchRecursive(list, start, mid-1, x)
	}
}

// 二分法求平方根
// precision表示保留小数后多少位
func sqrtByBinSearch(num, threshold float32) (float32, error) {
	if num < 0 {
		return 0.0, fmt.Errorf("invalid num:%v", num)
	}
	
	var (
		low float32 = 0.0
		high = num
	)
	if num < 1 {
		low = num
		high = 1.0
	}

	mid := (low + high) / 2.0
	for {
		// 与原数据误差控制在threshold以内
		dval := (mid * mid) - num
		if dval > -1*threshold && dval < threshold {
			break
		}

		if mid * mid > num {
			high = mid
		} else if mid * mid < num {
			low = mid
		}
		mid = (low + high) / 2.0
	}
	return mid, nil
}