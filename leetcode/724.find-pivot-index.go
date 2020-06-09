func pivotIndex(nums []int) int {
	n := len(nums)
	leftSum := make([]int, n)
	rightSum := make([]int, n)

	sum := 0
	for i := 1; i < n; i++ {
		sum += nums[i-1]
		leftSum[i] = sum
	}

	sum = 0
	for i := n - 2; i >= 0; i-- {
		sum += nums[i+1]
		rightSum[i] = sum
	}

	for i := 0; i < n; i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}