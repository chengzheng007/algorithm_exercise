/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
func subsets(nums []int) [][]int {
	size := len(nums)
	subSets := make([][]int, 0)
	currSet := make([]int, 0)
	// backtrack(0, size, nums, currSet, &subSets)
	backtrack2(nums, size, 0, currSet, &subSets)
	return subSets
}

// 回溯算法版本1-双递归实现
// func backtrack(i, size int, nums, currSet []int, sets *[][]int) {
// 	if i == size {
// 		*sets = append(*sets, currSet)
// 		return
// 	}
// 	// 不选第i个元素，先递归到最后一个索引处，然后每次向前一步，接着又往后递归
//  // 因此在i==size时才将子集加到结果集中
// 	backtrack(i+1, size, nums, currSet, sets)

// 	// 选第i个元素
// 	// 不能共用一个currSet变量，下面的append可能对上面的有影响，因go语言的实现机制导致
// 	currSet2 := make([]int, len(currSet)+1)
// 	copy(currSet2, currSet)
// 	currSet2[len(currSet)] = nums[i]
// 	backtrack(i+1, size, nums, currSet2, sets)
// }

// 回溯算法版本2-加一个循环实现
// i		k		curSet		curSet2
// 0		0		[]			[1]
// 1		1		[1]			[1,2]
// 2		2		[1,2]		[1,2,3]
// 2		1		[1]			[1,3]   (由i=2/k=2递归回退，由i=1/k=1下一次for循环得来，i++后i=2)
// 1		0		[]			[2]		(由i=2/k=1递归回退，由i=0/k=0下一次for循环得来，i++后i=1)
// 2		2		[2]			[2,3]
// 2		0		[]			[3]		(y由i=2/k=2递归回退，由i=1/k=0下一次for循环得来，i++后i=2)
func backtrack2(nums []int, size, start int, currSet []int, sets *[][]int) {
	*sets = append(*sets, currSet)
	for i := start; i < size; i++ {
		currSet2 := make([]int, len(currSet)+1)
		copy(currSet2, currSet)
		currSet2[len(currSet)] = nums[i]
		backtrack2(nums, size, i+1, currSet2, sets)
	}
}

// 另一种方法：每一个数可选可不选，假设n为总个数
// 则集合中肯定有2^n个元素（包含空集合）
func subsets(nums []int) [][]int {
	// 位运算
	totalSetNum := 1 << len(nums)
	totalSet := make([][]int, totalSetNum)
	for i := 0; i < totalSetNum; i++ {
		set := make([]int, 0)
		// 100、010、001，跟三个数与运算，判断某个位置
		// 的数是否在需要放在集合中
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) != 0 {
				set = append(set, nums[j])
			}
		}
		totalSet[i] = set
	}
	return totalSet
}

