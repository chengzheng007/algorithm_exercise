func mySqrt(x int) int {
	var (
		precision float64 = 0.01
		low       float64 = 0
		high              = float64(x)
		mid       float64
	)
	for {
		mid = (low + high) / 2.0
		diff := mid*mid - float64(x)
		if diff <= precision && diff > -1*precision {
			break
		} else if diff < 0 {
			low = mid
		} else {
			high = mid
		}
	}

	// 0-1的数统一返回1
	if mid > 0.0 && mid < 1.0 {
		return 1
	}

	return int(mid)
}

