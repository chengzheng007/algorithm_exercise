package main

import (
	"strconv"
	"fmt"
)

// 小顶堆

type Heap struct {
	// 存储数据的下标从0开始，左右子节点分别为2i+1、2i+2
	list []int
	length int
	capacity int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		list:make([]int, cap),
		capacity:cap,
	}
}

func (h *Heap) Insert(n int) bool {
	if h.length >= h.capacity {
		return false
	}
	// 将插入数据放在最后一位
	h.list[h.length] = n
	h.length++
	// 从下往上堆化
	i := h.length - 1
	// 注意下标从0开始，如果是(2i+2)/2=i+1，得到的可能不是父节点，所以必须用(i-1)/2得到父节点下标
	// 下标0的位置存储元素，所以堆化时，i必须大于0，因为i是跟自己的父节点比较，i如果等于0表示他自己已经是堆顶，不用动
	// 自下往上堆化
	for i > 0 && (i-1)/2 >= 0 && h.list[i] < h.list[(i-1)/2] {
		h.list[i], h.list[(i-1)/2] = h.list[(i-1)/2], h.list[i]
		i = (i-1)/2
	}
	return true
}

// 删除，得到堆顶元素
func (h *Heap) Remove() (int, bool) {
	if h.length <= 0 {
		return 0, false
	}

	ret := h.list[0]
	// 将最后一个元素覆盖到堆顶元素
	h.list[0] = h.list[h.length-1]
	h.length--

	// 从上往下堆化
	var (
		i = 0
		minPos = i
	)
	for {
		minPos = i
		if 2*i+1 < h.length && h.list[2*i+1] < h.list[i] {
			minPos = 2*i+1
		}
		if 2*i+2 < h.length && h.list[2*i+2] < h.list[minPos] {
			minPos = 2*i+2
		}
		// 已调整完毕
		if minPos == i {
			break
		}
		// swap
		h.list[i], h.list[minPos] = h.list[minPos], h.list[i]
		i = minPos
	}

	return ret, true
}

func (h Heap) String() string {
	outputFormat := "heap len:%d, cap:%d, data:[%s]"
	datastr := ""
	for i := 0; i < h.length; i++ {
		if datastr != "" {
			datastr += ","
		}
		datastr += strconv.Itoa(h.list[i])
	}
	return fmt.Sprintf(outputFormat, h.length, h.capacity, datastr)
}
