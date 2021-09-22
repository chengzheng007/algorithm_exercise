/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
// 更易理解的、直观的做法
func plusOne(digits []int) []int {
    carry := 0
    n := len(digits)
    // carry 表示进位，无需进位时carry为0，类似于两数相加的思想
    for i := n-1; i >= 0; i-- {
        newNum := digits[i]+carry
        if i == n-1 {
            newNum += 1
        }
        digits[i] = newNum%10
        carry = newNum/10
    }
    
    if carry == 0 {
        return digits
    }
    
    newDigits := make([]int, n+1)
    newDigits[0] = carry
    for i := 0; i < n; i++ {
        newDigits[i+1] = digits[i]
    }
    
    return newDigits
}


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

