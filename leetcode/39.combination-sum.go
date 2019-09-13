/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	cand := make([]int, 0)
	bt(candidates, target, 0, cand, 0, &result)
	return result
}

// 回溯
func bt(candidates []int, target, start int, currCand []int, currSum int, result *[][]int) {
	if currSum == target {
		*result = append(*result, currCand)
		return
	} else if currSum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		candNew := make([]int, len(currCand)+1)
		copy(candNew[0:len(currCand)], currCand)
		candNew[len(currCand)] = candidates[i]
		// 继续递归用当前值计算，直到比target大才换成下一个值
		bt(candidates, target, i, candNew, currSum+candidates[i], result)
	}
}

