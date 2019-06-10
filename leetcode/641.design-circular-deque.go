/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */
type MyCircularDeque struct {
	size     int
	list     []int
	currSize int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	mydeque := MyCircularDeque{
		size: k,
		list: make([]int, k),
	}
	for i := 0; i < k; i++ {
		mydeque.list[i] = -1
	}
	return mydeque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.currSize+1 > this.size {
		return false
	}
	for i := this.currSize - 1; i >= 0; i-- {
		this.list[i+1] = this.list[i]
	}
	this.list[0] = value
	this.currSize += 1
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.currSize+1 > this.size {
		return false
	}
	this.list[this.currSize] = value
	this.currSize += 1
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.currSize == 0 {
		return false
	}
	for i := 0; i < this.currSize-1; i++ {
		this.list[i] = this.list[i+1]
	}
	this.currSize -= 1
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.currSize == 0 {
		return false
	}
	this.list[this.currSize-1] = 0
	this.currSize -= 1
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.currSize == 0 {
		return -1
	}
	return this.list[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.currSize == 0 {
		return -1
	}
	return this.list[this.currSize-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.currSize == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.currSize == this.size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

