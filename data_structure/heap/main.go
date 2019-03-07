package main

import (
	"fmt"
)

// 堆
type Heap struct {
	count int       // 实际个数
	capacity int    // 最大容量
	list []int
}

func main() {
	list := []int{5,7,2,1,9,10,6}
	heap := NewHeap(len(list))
	for _, num := range list {
		heap.Push(num)
	}
	heap.Print()
	fmt.Println()

	//size := heap.Len()
	//for i := 0; i < size; i++ {
	//	num, ret := heap.Pop()
	//	fmt.Printf("heap pop num:%v, ret:%v\n", num, ret)
	//	heap.Print()
	//}
	//
	//num, ret := heap.Pop()
	//fmt.Printf("heap pop num:%v, ret:%v\n", num, ret)
	//heap.Print()
	
	heap.HeapSort()
	heap.Print()


	list = []int{0,5,6,3,9,1,2,8}
	buildHeap(list, len(list)-1)
	fmt.Printf("list transfor heap:%v\n", list[1:])

}

func NewHeap(size int) *Heap {
	return &Heap{
		capacity:size,
		list:make([]int, size+1),
	}
}

// 以小顶堆为例

// 入堆
func (h *Heap) Push(num int) {
	if h.count >= h.capacity {
		return
	}

	h.count++
	// 第一个位置空出，根节点存储在数组1开始的地方
	h.list[h.count] = num
	// 堆调整，从下往上交换建堆
	pos := h.count
	for pos/2 > 0 && h.list[pos/2] > h.list[pos] {
		h.list[pos/2], h.list[pos] = h.list[pos], h.list[pos/2]
		pos /= 2
	}
}

func (h *Heap) Print() {
	if h.count > 0 {
		fmt.Printf("heap count:%d, elems:%v\n", h.count, h.list[1:h.count+1])
	} else {
		fmt.Println("heap count:0, elems:no elem\n")
	}
}

func (h *Heap) Pop() (int, bool) {
	var (
		num = -1
		ret = false
	)
	if h.count <= 0 {
		return num, ret
	}

	num = h.list[1]
	ret = true
	  

	h.list[1] = h.list[h.count]
	h.count--

	// 从上往下进行堆调整
	i := 1
	for {
		pos := i
		if 2*i <= h.count && h.list[2*i] < h.list[pos] {
			pos = 2*i
		}
		if 2*i+1 <= h.count && h.list[2*i+1] < h.list[pos] {
			pos = 2*i+1
		}
		if pos == i { // 调整完毕
			break
		}
		h.list[pos], h.list[i] =  h.list[i], h.list[pos]
		i = pos
	}
	return num, ret
}

func (h *Heap) Len() int {
	return h.count
}

// 堆排序，将最小值放在数组末尾
func (h *Heap) HeapSort() {
	i := h.count
	for i > 1 {
		h.list[1], h.list[i] = h.list[i], h.list[1]
		i--
		heapfy(h.list, i, 1)
	}
}

// 从上到下堆化
func heapfy(list []int, n, i int) {
	for {
		pos := i
		if 2*i <= n && list[2*i] < list[pos] {
			pos = 2*i
		}
		if 2*i+1 <= n && list[2*i+1] < list[pos] {
			pos = 2*i+1
		}
		if pos == i {
			break
		}
		list[pos], list[i] = list[i], list[pos]
		i = pos
	}
}

// 从1到n/2开始建堆，从上到下调整堆，n/2+1开始的都是叶子节点，不用动
// 第一个位置空着
func buildHeap(list []int, n int) {
	for i := n/2; i >= 1; i-- {
		heapfy(list, n, i)
	}
}