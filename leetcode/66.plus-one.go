/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	num := 0
	unit := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num += digits[i] * unit
		unit *= 10
	}
	num += 1
	newDigits := make([]int, 0)
	for num > 0 {
		newDigits = append(newDigits, num%10)
		num /= 10
	}
	// 翻转
	for i := 0; i < len(newDigits)/2; i++ {
		newDigits[i], newDigits[len(newDigits)-i-1] = newDigits[len(newDigits)-i-1], newDigits[i]
	}
	return newDigits
}

