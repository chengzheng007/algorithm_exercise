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
	// 如果只是单纯的压入左右括号，是不好计算长度的
	// 尤其是(())这种连续的，比如每次碰到)就出栈+2，但
	// 这种情况无法确定上一次是不是刚处理一对对称的，即，
	// 无法确定在上一次基础上加2
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

// dp
func longestValidParentheses2(s string) int {
	size := len(s)
	if size < 2 {
		return 0
	}
	// dp[i]表示截止下标为i字符的最长格式化好的括号个数
	dp := make([]int, size)
	max := 0
	for i := 1; i < size; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// 如果前一个是(，直接与它配对
				// 自身长度等于(前一个字符长度(dp[i-2])+当前长度2
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}