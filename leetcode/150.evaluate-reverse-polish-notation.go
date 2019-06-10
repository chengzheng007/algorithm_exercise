import (
	"container/list"
	"strconv"
)

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
func evalRPN(tokens []string) int {
	stack := list.New()
	for _, str := range tokens {
		var num int
		if str == "+" || str == "-" || str == "*" || str == "/" {
			elem := stack.Front()
			num1 := elem.Value.(int)
			stack.Remove(elem)
			elem = stack.Front()
			num2 := elem.Value.(int)
			stack.Remove(elem)

			switch str {
			case "+":
				num = num2 + num1
			case "-":
				num = num2 - num1
			case "*":
				num = num2 * num1
			case "/":
				if num1 != 0 {
					num = num2 / num1
				} else {
					return 0
				}
			}
		} else {
			num, _ = strconv.Atoi(str)
		}
		stack.PushFront(num)
	}
	ret := stack.Front().Value.(int)
	return ret
}

