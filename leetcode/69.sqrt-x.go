/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
func mySqrt(x int) int {
	var (
		precision float64 = 0.01
		low       float64 = 0.0
		high              = float64(x)
		mid       float64
	)
	for {
		mid = (low + high) / 2.0
		mid2 := mid * mid
		diff := mid2 - float64(x)
		if diff <= precision && diff > -1*precision {
			break
		} else if mid2 < float64(x) {
			low = mid
		} else {
			high = mid
		}
	}
	if mid > 0.0 && mid < 1.0 {
		return 1
	}
	return int(mid)
}

