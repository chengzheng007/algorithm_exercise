import "math"

type MedianFinder struct {
	// big heap store first half of stream
	bigHeap []int
	// small heap store second half of stream
	smallHeap []int
	// then all data big heap is smaller than small heap
	// top of big heap is the median of array if the number
	// of array is odd, median is (top of big heap +
	// top of small heap) / 2.0 if the number is even
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		bigHeap:   make([]int, 0),
		smallHeap: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	bigHeapTopNum := math.MaxInt64
	bLen := len(this.bigHeap)
	sLen := len(this.smallHeap)

	if bLen != 0 {
		bigHeapTopNum = this.bigHeap[0]
	}

	if num <= bigHeapTopNum {
		// insert bigheap heapfy, from botton to up
		this.bigHeap = append(this.bigHeap, num)
		bLen++
		bottomToUpHeapfy(this.bigHeap, bLen, true)
		// 两个堆数据不均衡，多出超过2个时要调整
		if bLen > sLen+1 {
			bigHeapTopNum = this.bigHeap[0]

			// 删除大顶堆堆顶元素，插入小顶堆
			// 将大顶堆末尾元素放置堆顶并调整
			this.bigHeap[0] = this.bigHeap[bLen-1]
			this.bigHeap = this.bigHeap[:bLen-1]
			bLen--
			upToBottomHeapfy(this.bigHeap, bLen, 0, true)

			// 插入并调整小顶堆
			this.smallHeap = append(this.smallHeap, bigHeapTopNum)
			sLen++
			bottomToUpHeapfy(this.smallHeap, sLen, false)
		}
	} else {
		this.smallHeap = append(this.smallHeap, num)
		sLen++
		bottomToUpHeapfy(this.smallHeap, sLen, false)

		if sLen > bLen+1 {
			smallHeapTopNum := this.smallHeap[0]

			// 删除小顶堆堆顶元素，插入大顶堆
			this.smallHeap[0] = this.smallHeap[sLen-1]
			this.smallHeap = this.smallHeap[:sLen-1]
			sLen--
			upToBottomHeapfy(this.smallHeap, sLen, 0, false)
			// 插入并调整大顶堆
			this.bigHeap = append(this.bigHeap, smallHeapTopNum)
			bLen++
			bottomToUpHeapfy(this.bigHeap, bLen, true)
		}
	}
}

func bottomToUpHeapfy(nums []int, n int, isBigHeap bool) {
	i := n - 1
	if isBigHeap {
		for i > 0 && (i-1)/2 >= 0 && nums[(i-1)/2] < nums[i] {
			nums[(i-1)/2], nums[i] = nums[i], nums[(i-1)/2]
			i = (i - 1) / 2
		}
	} else {
		for i > 0 && (i-1)/2 >= 0 && nums[(i-1)/2] > nums[i] {
			nums[(i-1)/2], nums[i] = nums[i], nums[(i-1)/2]
			i = (i - 1) / 2
		}
	}
}

func upToBottomHeapfy(nums []int, n, i int, isBigHeap bool) {
	for {
		pos := i
		// find the most bigger num index
		if 2*i+1 < n && ((isBigHeap && nums[2*i+1] > nums[pos]) || (!isBigHeap && nums[2*i+1] < nums[pos])) {
			pos = 2*i + 1
		}
		if 2*i+2 < n && ((isBigHeap && nums[2*i+2] > nums[pos]) || (!isBigHeap && nums[2*i+2] < nums[pos])) {
			pos = 2*i + 2
		}
		if i == pos {
			break
		}
		// swap parent and son node data
		nums[i], nums[pos] = nums[pos], nums[i]
		i = pos
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var (
		median float64
		sLen   = len(this.smallHeap)
		bLen   = len(this.bigHeap)
	)

	// 当两个堆数据相等，总数为偶数，中位数取中间两数的均值
	if sLen == bLen {
		if sLen == 0 {
			return median
		}
		median = float64(this.smallHeap[0]+this.bigHeap[0]) / 2.0
	} else if bLen > sLen {
		// 当两个堆的元素不等时，说明总数据个数是奇数，比较两个堆数据个数，选取一个多的堆堆顶即可
		median = float64(this.bigHeap[0])
	} else {
		median = float64(this.smallHeap[0])
	}
	return median
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */