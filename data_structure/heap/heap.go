package heap

type HeapType int

const (
	SmallHeap HeapType = 1
	BigHeap   HeapType = 2
)

type Heap struct {
	heapType HeapType //堆类型
	cap      int
	tail     int
	list     []int
}

func NewHeap(cap int, heapType HeapType) *Heap {
	if cap <= 0 {
		panic("invalid heap cap")
	}
	if heapType != SmallHeap && heapType != BigHeap {
		panic("invalid heap type")
	}
	return &Heap{
		heapType: heapType,
		cap:      cap,
		tail:     -1,
		list:     make([]int, cap), // 下标从0开始，i的子节点下标为2i+1、2i+2
	}
}

func (h *Heap) Insert(num int) bool {
	if h.tail+1 >= h.cap {
		return false
	}
	h.tail++
	h.list[h.tail] = num
	// 堆调整
	h.tailHeapfy()
	return true
}

func (h *Heap) Top() (int, bool) {
	if h.tail < 0 {
		return -1, false
	}
	return h.list[0], true
}

func (h *Heap) List() []int {
	if h.tail < 0 {
		return nil
	}
	return h.list[0 : h.tail+1]
}

// 从末尾调整
func (h *Heap) tailHeapfy() {
	if h.tail <= 0 {
		return
	}
	i := h.tail
	for {
		if i <= 0 {
			break
		}
		parentIdx := (i - 1) / 2
		// 小顶堆&父节点已经比子节点小
		// 大顶堆&父节点已经比子节点大
		if (h.heapType == SmallHeap && h.list[parentIdx] < h.list[i]) ||
			(h.heapType == BigHeap && h.list[parentIdx] > h.list[i]) {
			break
		}
		h.list[parentIdx], h.list[i] = h.list[i], h.list[parentIdx]
		i = parentIdx
	}
}

// Delete 永远删除堆顶元素
func (h *Heap) Delete() (int, bool) {
	if h.tail < 0 {
		return -1, false
	}
	num := h.list[0]
	h.list[0] = h.list[h.tail]
	h.tail--
	h.headHeapfy(0)
	return num, true
}

// 从上往下调整
func (h *Heap) headHeapfy(i int) {
	for {
		pos := i
		if 2*i+1 <= h.tail {
			if (h.heapType == SmallHeap && h.list[i] > h.list[2*i+1]) ||
				(h.heapType == BigHeap && h.list[i] < h.list[2*i+1]) {
				pos = 2*i + 1
			}
		}
		if 2*i+2 <= h.tail {
			if (h.heapType == SmallHeap && h.list[pos] > h.list[2*i+2]) ||
				(h.heapType == BigHeap && h.list[pos] < h.list[2*i+2]) {
				pos = 2*i + 2
			}
		}
		if i == pos {
			break
		}
		h.list[i], h.list[pos] = h.list[pos], h.list[i]
		i = pos
	}
}
