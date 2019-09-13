import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	currCand := make([]int, 0)
	dfs(candidates, target, 0, 0, currCand, &result)
	return result
}

func dfs(candidates []int, target, start, currSum int, currCand []int, result *[][]int) {
	if currSum == target {
		// 在结果中寻找相同的结果集，如果有则丢弃
		// 这种检测方式很蠢
		// var owned bool
		// for _, haveCand := range *result {
		// 	if len(haveCand) != len(currCand) {
		// 		continue
		// 	}
		// 	owned = true
		// 	for i := 0; i < len(haveCand); i++ {
		// 		if haveCand[i] != currCand[i] {
		// 			owned = false
		// 			break
		// 		}
		// 	}
		// 	// 如果其中某一个匹配当前的可选数字，退出
		// 	// 注意这里必须检查，否则后面的结果会覆盖前面的
		// 	if owned {
		// 		break
		// 	}
		// }
		// // 已拥有该选项
		// if len(*result) > 0 && owned {
		// 	return
		// }
		*result = append(*result, currCand)
	} else if currSum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		// 当前面一个可选值已经在上一层递归中使用过了，下一层直接过滤
		// 举例当target=8,candidates=[10,1,2,7,6,1,5]=[1,1,2,5,6,7,10]
		// 回溯算法第一次遇到1,1,6时会加到结果中，因为此时i==start
		// 当到下一层递归时，i>start，这时从第二个1开始，这是新一轮递归
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		newCand := make([]int, len(currCand)+1)
		copy(newCand, currCand)
		newCand[len(currCand)] = candidates[i]
		dfs(candidates, target, i+1, currSum+candidates[i], newCand, result)
	}
}

