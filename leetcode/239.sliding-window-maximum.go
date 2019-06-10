import "container/list"

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	numArr := make([]int, 0)
	if size == 0 || size < k || k == 0 {
		return numArr
	}
	// double ended queue
	deque := list.New()
	for i := 0; i < size; i++ {
		// only store index that it point to num in array great or equal than current num
		for deque.Len() > 0 {
			elem := deque.Back()
			idx := elem.Value.(int)
			if nums[idx] < nums[i] {
				deque.Remove(elem)
			} else {
				break
			}
		}
		// add index to last
		deque.PushBack(i)
		// remove last sliding window's num by comparing index
		elem := deque.Front()
		if i-elem.Value.(int) == k {
			deque.Remove(elem)
		}
		// process done one window, take out biggest num
		if i >= k-1 {
			idx := deque.Front().Value.(int)
			numArr = append(numArr, nums[idx])
		}
	}
	return numArr
}

