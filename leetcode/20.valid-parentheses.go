/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
import "container/list"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := list.New()
	for _, byt := range s {
		s := string(byt)
		if s == "(" || s == "[" || s == "{" {
			stack.PushFront(s)
		}
		if s == ")" || s == "]" || s == "}" {
			if stack.Len() == 0 {
				return false
			}
			elem := stack.Front()
			val := elem.Value.(string)
			switch s {
			case ")":
				if val != "(" {
					return false
				}
			case "]":
				if val != "[" {
					return false
				}
			case "}":
				if val != "{" {
					return false
				}
			default:
				return false
			}
			stack.Remove(elem)
		}
	}
	if stack.Len() > 0 {
		return false
	}
	return true
}

