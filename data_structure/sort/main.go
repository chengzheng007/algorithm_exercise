package main

import "fmt"

func main() {
	list := []int{5,4,3,3,2,1,9,0}
	bubbleSort(list)
	//fmt.Println("bubbleSort:", list)

	list = []int{2,6,1,6,-5}
	insertSort(list)
	//fmt.Println("insertSort:", list)

	list = []int{5,2,6,1,6,-5}
	binInsertSort(list)
	fmt.Println("binInsertSort:", list)

	list = []int{2,6,1,6,-5}
	selectSort(list)
	//fmt.Println("selectSort:", list)

	list = []int{2,35,1429,64,5,10,12,30,-5}
	shellSort(list)
	//fmt.Println("shellSort:", list)

	list = []int{2,-5,10,3,10,12}
	quickSort(list, 0, len(list)-1)
	//fmt.Println("quickSort:", list)

	list = []int{2,-5,6,1,3}
	fmt.Println("list:", list)
	mergeSort(list, 0, len(list))
	fmt.Println("mergeSort:", list)
}

// 冒泡
func bubbleSort(list []int) {
	size := len(list)
	if size <= 1 {
		return
	}
	for i := 0; i < size - 1; i++ { // 最后一趟只剩一个元素不用交换
		flag := false
		for j := 0; j < size - i - 1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
			}
		}
		// 无交换，已有序
		if !flag {
			break
		}
	}
}

// 直接插入法
func insertSort(list []int) {
	size := len(list)
	if size <= 1 {
		return
	}
	for i := 1; i < size; i++ {
		temp := list[i]
		// 寻找插入位置
		j := i-1
		for ; j >= 0; j-- {
			if list[j] > temp {
				list[j+1] = list[j]
			} else {
				break
			}
		}
		list[j+1] = temp
	}
}

// 二分插入排序
func binInsertSort(list []int) {
	size := len(list)
	if size <= 1 {
		return
	}
	for i := 1; i < size; i++ {
		temp := list[i]
		l := 0
		r := i - 1
		for l <= r {
			mid := (l+r)/2
			if temp < list[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		// l为最终数据插入位置, l及右边数据右移
		for j := i-1; j >= l; j-- {
			list[j+1] = list[j]
		}
		list[l] = temp
	}
}

// 选择法
func selectSort(list []int) {
	size := len(list)
	if size <= 1 {
		return
	}
	for i := 0; i < size; i++ {
		minIndex := i
		for j := i+1; j < size; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}
}

// 希尔排序
func shellSort(list []int) {
	size := len(list)
	if size <= 1 {
		return
	}
	for N := size/2; N > 0; N-- {
		for i := N; i < size; i++ {
			temp := list[i]
			j := i - N
			for ; j >= 0; j -= N {
				if list[j] > temp {
					list[j+N] = list[j]
				} else {
					break
				}
			}
			list[j+N] = temp
		}
	}
}

// 快排 - 分治思想
// l:0 r: length of list - 1
func quickSort(list []int, l, r int) {
	if l >= r {
		return
	}

	// 选取一个分区点，左边比它小，右边比它大，返回这个分区点最终下标
	// 排序实际在分区函数中完成
	pIndex := partion(list, l, r)   
	// 递归处理分区点左边部分
	quickSort(list, l, pIndex-1)
	// 递归处理分区点右边部分
	quickSort(list, pIndex+1, r)
}

func partion(list []int, l, r int) int {
	// 选最后一个元素作为分界点，比它小的移到list左边，大的移到右边
	temp := list[r]
	i := l
	for j := l; j < r; j++ {
		if list[j] < temp {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	// 循环结束，i为分界点元素插入位置
	list[i], list[r] = list[r], list[i]
	return i
}

// 归并排序
func mergeSort(list []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l+r)/2
	
	// 分治递归
	mergeSort(list, l, mid)
	mergeSort(list, mid+1, r)
	// 将两个部分合并到一起
	// 排序实际在该合并函数中完成
	//fmt.Printf("befeor merge == l:%d, mid:%d, r:%d, list[%d:%d]:%v, list[%d:%d]:%v\n", l, mid, r, l, mid, list[l:mid], mid, r, list[mid:r])
	merge(list, l, r, list[l:mid], list[mid:r])
}

func merge(list []int, l, r int, preList, postList []int,) {
	i := 0
	j := 0
	p := 0
	tmp := make([]int, r-l)
	//fmt.Printf("r(%d)-l(%d)=>%d\n", r, l, r-l)
	for i < len(preList) && j < len(postList) {
		if preList[i] <= postList[j] {
			tmp[p] = preList[i]
			i++
		} else {
			tmp[p] = postList[j]
			j++
		}
		p++
	}

	//fmt.Println("tmp merge two sametime:", tmp)

	//fmt.Println("p=>", p, "i=>", i, "j=>", j)
	if i < len(preList) {
		copy(tmp[p:], preList[i:])
		//fmt.Printf("preList tmp:%v, preList[%d:]:%v\n", tmp, i, preList[i:])
	} else if j < len(postList) {
		copy(tmp[p:], postList[j:])
		//fmt.Printf("postList tmp:%v, postList[%d:]:%v\n", tmp, j, postList[j:])
	}

	//fmt.Println("end tmp:", tmp)
	copy(list[l:r], tmp)
	//fmt.Printf("list[%d:%d]:%v ,list:%v, preList:%v, postList:%v\n", l, r, list[l:r], list, preList, postList)
}