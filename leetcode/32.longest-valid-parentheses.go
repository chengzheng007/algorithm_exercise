import "container/list"

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := 0
	stack := list.New()
	stack.PushFront(-1)
	for i, char := range s {
		str := string(char)
		if str == "(" {
			stack.PushFront(i)
		} else if str == ")" {
			stack.Remove(stack.Front())

			if stack.Len() > 0 {
				idx := stack.Front().Value.(int)
				tmp := i - idx
				if tmp > length {
					length = tmp
				}
			} else {
				stack.PushFront(i)
			}
		}
	}
	return length
}

