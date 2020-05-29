import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		// res stores the max product
		res = nums[0]
		// imax/imin stores the max/min product of subarray that ends with current number nums[i]
		imax = res
		imin = res
	)

	for i := 1; i < len(nums); i++ {
		// multiplied by a negative makes big number smaller, small number bigger, so we redefine the extremum by swapping them
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		// max/min product for the current number is either the current number itself, or the max/min by the previous times(乘以) the current one
		imax = int(math.Max(float64(nums[i]), float64(nums[i]*imax)))
		imin = int(math.Min(float64(nums[i]), float64(nums[i]*imin)))

		if imax > res {
			res = imax
		}
	}
	return res
}