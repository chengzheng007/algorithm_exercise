package main

import "fmt"

func main() {
	list := []int{5,4,3,3,2,1,9,0}
	bubbleSort(list)
	fmt.Println("bubbleSort:", list)

	list = []int{2,6,1,6,-5}
	insertSort(list)
	fmt.Println("insertSort:", list)

	list = []int{5,2,6,1,6,-5}
	binInsertSort(list)
	fmt.Println("binInsertSort:", list)

	list = []int{2,6,1,6,-5}
	selectSort(list)
	fmt.Println("selectSort:", list)

	list = []int{2,35,1429,64,5,10,12,30,-5}
	shellSort(list)
	fmt.Println("shellSort:", list)

	list = []int{2,-5,10,3,10,12}
	quickSort(list, 0, len(list)-1)
	fmt.Println("quickSort:", list)

	list = []int{2,-5,6,1,3}
	mergeSort(list, 0, len(list)-1)
	fmt.Println("mergeSort:", list)

	list = []int{2,5,3,0,2,3,0,3,9,1}
	if err := countingSort(list); err != nil {
		fmt.Printf("countingSort() error(%v)\n", err)
	} else {
		fmt.Println("countingSort:", list)
	}

	list = []int{36,5,16,98,95,47,32,36,48,10}
	if err := radixSort(list); err != nil {
		fmt.Printf("radixSort() error(%v)\n", err)
	} else {
		fmt.Println("radixSort:", list)
	}

	
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
	merge(list, l, mid, r)
}

func merge(list []int, l, mid, r int) {
	tmpArr := make([]int, r-l+1)

	i := l
	j := mid+1
	k := 0
	// slice需要访问到最左或最右的元素
	for ; i <= mid && j <= r; k++ {
		if list[i] < list[j] {
			tmpArr[k] = list[i]
			i++
		} else {
			tmpArr[k] = list[j]
			j++
		}
	}
	
	for ; i <= mid; i++ {
		tmpArr[k] = list[i]
		k++
	}

	for ; j <= r; j++ {
		tmpArr[k] = list[j]
		k++
	}
	copy(list[l:r+1], tmpArr)
}

// 计数排序 - 【这里数组中只允许非负数】
// 适用于数据大小范围足够小的情况
func countingSort(list []int) error {
	size := len(list)
	if size <= 1 {
		return nil
	}

	// 遍历找到最大值
	max := -1
	for _, val := range list {
		if val < 0 {
			return fmt.Errorf("data can't be less than zero")
		}
		if max < val {
			max = val
		}
	}

	// 申请max+1个容量的数组，计算数据个数
	c := make([]int, max+1)
	for _, val := range list {
		c[val]++
	}

	// 将c中的数据个数累加
	for i := 1; i < max+1; i++ {
		c[i] = c[i] + c[i-1]
	}

	// 反向遍历，将排序数据放入排序列表
	rankList := make([]int, size)
	for i := size-1; i >= 0; i-- {
		rankList[c[list[i]]-1] = list[i]
		c[list[i]]--
	}

	copy(list, rankList)
	return nil
}

// 基数排序 - 【以最大数字为两位数为例】
func radixSort(list []int) error {
	size := len(list)
	if size <= 1 {
		if list[0] < 0 || list[0] >= 100 {
			return fmt.Errorf("num in list should between 0 and 99")
		}
		return nil
	}

	// 分配10个桶，0...9，每个桶可动态扩容
	bucket := make([][]int, 10)  // 个位数
		
	// 第一次按低位遍历分配
	for _, num := range list {
		if num < 0 || num >= 100 {
			return fmt.Errorf("num in list should between 0 and 99")
		}
		bucket[num%10] = append(bucket[num%10], num)
	}

	i := 0
	for _, elems := range bucket {
		for _, num := range elems {
			list[i] = num
			i++
		}
	}

	// 第二次按高位遍历分配到桶(这里桶元素类型和长度相等)
	bucket = make([][]int, 10)
	for _, num := range list {
		bucket[num/10] = append(bucket[num/10], num)
	}

	i = 0
	for _, elems := range bucket {
		for _, num := range elems {
			list[i] = num
			i++
		}
	}
	return nil
}