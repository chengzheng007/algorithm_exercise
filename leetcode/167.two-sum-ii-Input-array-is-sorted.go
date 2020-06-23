func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n < 2 {
		return nil
	}

	var res []int

	for i := 0; i < n; i++ {
		num := target - numbers[i]
		l := 0
		r := n - 1
		for l <= r {
			mid := (l + r) / 2
			if numbers[mid] == num {
				// 数组内数字不能复用
				if mid != i {
					res = []int{i + 1, mid + 1}
					return res
				}
				l = mid + 1
			} else if num > numbers[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return res
}